/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"crypto/tls"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/google/uuid"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/tests/e2e/driver"
	"github.com/kubernetes-sigs/aws-ebs-csi-driver/tests/e2e/testsuites"
	"k8s.io/kubernetes/test/e2e/framework"
	"net/http"
	"os"
	"strconv"

	volumesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	. "github.com/onsi/ginkgo/v2"
	v1 "k8s.io/api/core/v1"
	clientset "k8s.io/client-go/kubernetes"
	restclientset "k8s.io/client-go/rest"
	admissionapi "k8s.io/pod-security-admission/api"

	awscloud "github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/cloud"
	ebscsidriver "github.com/kubernetes-sigs/aws-ebs-csi-driver/pkg/driver"
)

const testTagNamePrefix = "testTag"
const testTagValue = "3.1415926"

// generateTagName appends a random uuid to tag name to prevent clashes on parallel e2e test runs on shared cluster
func generateTagName() string {
	return testTagNamePrefix + uuid.NewString()[:8]
}

func validateEc2Snapshot(ec2Client ec2iface.EC2API, input *ec2.DescribeSnapshotsInput) *ec2.DescribeSnapshotsOutput {
	describeResult, err := ec2Client.DescribeSnapshots(input)
	if err != nil {
		Fail(fmt.Sprintf("failed to describe snapshot: %v", err))
	}

	if len(describeResult.Snapshots) != 1 {
		Fail(fmt.Sprintf("expected 1 snapshot, got %d", len(describeResult.Snapshots)))
	}

	return describeResult
}

var _ = Describe("[ebs-csi-e2e] [single-az] [requires-aws-api] Dynamic Provisioning", func() {
	f := framework.NewDefaultFramework("ebs")
	f.NamespacePodSecurityEnforceLevel = admissionapi.LevelPrivileged

	var (
		cs          clientset.Interface
		snapshotrcs restclientset.Interface
		ns          *v1.Namespace
		ebsDriver   driver.PVTestDriver
	)

	BeforeEach(func() {
		cs = f.ClientSet
		var err error
		snapshotrcs, err = restClient(testsuites.SnapshotAPIGroup, testsuites.APIVersionv1)
		if err != nil {
			Fail(fmt.Sprintf("could not get rest clientset: %v", err))
		}
		ns = f.Namespace
		ebsDriver = driver.InitEbsCSIDriver()
	})

	var awsConfig *aws.Config
	region := os.Getenv("AWS_REGION")
	envEndpointInsecure := os.Getenv("AWS_EC2_ENDPOINT_UNSECURE")
	isEndpointInsecure := false
	if envEndpointInsecure != "" {
		isEndpointInsecure, _ = strconv.ParseBool(envEndpointInsecure)
	}

	if isEndpointInsecure {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		awsConfig = &aws.Config{
			Region:                        aws.String(region),
			CredentialsChainVerboseErrors: aws.Bool(true),
			HTTPClient:                    client,
		}
	} else {
		awsConfig = &aws.Config{
			Region:                        aws.String(region),
			CredentialsChainVerboseErrors: aws.Bool(true),
			// Set MaxRetries to a high value. It will be "overwritten" if context deadline comes sooner.
			MaxRetries: aws.Int(8),
		}
	}

	endpoint := os.Getenv("AWS_EC2_ENDPOINT")
	if endpoint != "" {
		customResolver := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
			if service == ec2.EndpointsID {
				return endpoints.ResolvedEndpoint{
					URL:           endpoint,
					SigningRegion: region,
				}, nil
			}
			return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
		}
		awsConfig.EndpointResolver = endpoints.ResolverFunc(customResolver)
	}

	// Tests that require that the e2e runner has access to the AWS API
	ec2Client := ec2.New(session.Must(session.NewSession(awsConfig)))

	It("should create a volume with additional tags", func() {
		testTag := generateTagName()
		pods := []testsuites.PodDetails{
			{
				Cmd: testsuites.PodCmdWriteToVolume("/mnt/test-1"),
				Volumes: []testsuites.VolumeDetails{
					{
						CreateVolumeParameters: map[string]string{
							ebscsidriver.VolumeTypeKey: awscloud.VolumeTypeGP2,
							ebscsidriver.FSTypeKey:     ebscsidriver.FSTypeExt4,
							ebscsidriver.TagKeyPrefix:  fmt.Sprintf("%s=%s", testTag, testTagValue),
						},
						ClaimSize:   driver.MinimumSizeForVolumeType(awscloud.VolumeTypeGP2),
						VolumeMount: testsuites.DefaultGeneratedVolumeMount,
					},
				},
			},
		}
		test := testsuites.DynamicallyProvisionedCmdVolumeTest{
			CSIDriver: ebsDriver,
			Pods:      pods,
			ValidateFunc: func() {
				result, err := ec2Client.DescribeVolumes(&ec2.DescribeVolumesInput{
					Filters: []*ec2.Filter{
						{
							Name:   aws.String("tag:" + testTag),
							Values: []*string{aws.String(testTagValue)},
						},
					},
				})
				if err != nil {
					Fail(fmt.Sprintf("failed to describe volume: %v", err))
				}

				if len(result.Volumes) != 1 {
					Fail(fmt.Sprintf("expected 1 volume, got %d", len(result.Volumes)))
				}
			},
		}
		test.Run(cs, ns)
	})

	It("should create a snapshot with additional tags", func() {
		testTag := generateTagName()
		pod := testsuites.PodDetails{
			Cmd: testsuites.PodCmdWriteToVolume("/mnt/test-1"),
			Volumes: []testsuites.VolumeDetails{
				{
					CreateVolumeParameters: map[string]string{
						ebscsidriver.VolumeTypeKey: awscloud.VolumeTypeGP2,
						ebscsidriver.FSTypeKey:     ebscsidriver.FSTypeExt4,
					},
					ClaimSize:   driver.MinimumSizeForVolumeType(awscloud.VolumeTypeGP2),
					VolumeMount: testsuites.DefaultGeneratedVolumeMount,
				},
			},
		}
		restoredPod := testsuites.PodDetails{
			Cmd: testsuites.PodCmdGrepVolumeData("/mnt/test-1"),
			Volumes: []testsuites.VolumeDetails{
				{
					CreateVolumeParameters: map[string]string{
						ebscsidriver.VolumeTypeKey: awscloud.VolumeTypeGP2,
						ebscsidriver.FSTypeKey:     ebscsidriver.FSTypeExt4,
					},
					ClaimSize:   driver.MinimumSizeForVolumeType(awscloud.VolumeTypeGP2),
					VolumeMount: testsuites.DefaultGeneratedVolumeMount,
				},
			},
		}
		test := testsuites.DynamicallyProvisionedVolumeSnapshotTest{
			CSIDriver:   ebsDriver,
			Pod:         pod,
			RestoredPod: restoredPod,
			Parameters: map[string]string{
				ebscsidriver.TagKeyPrefix: fmt.Sprintf("%s=%s", testTag, testTagValue),
			},
			ValidateFunc: func(_ *volumesnapshotv1.VolumeSnapshot) {
				validateEc2Snapshot(ec2Client, &ec2.DescribeSnapshotsInput{
					Filters: []*ec2.Filter{
						{
							Name:   aws.String("tag:" + testTag),
							Values: []*string{aws.String(testTagValue)},
						},
					},
				})
			},
		}
		test.Run(cs, snapshotrcs, ns)
	})

	It("should create a snapshot with FSR enabled", func() {

		Skip("Fast Snapshot restore is not supported.")

		azList, err := ec2Client.DescribeAvailabilityZones(&ec2.DescribeAvailabilityZonesInput{})
		if err != nil {
			Fail(fmt.Sprintf("failed to list AZs: %v", err))
		}
		fsrAvailabilityZone := *azList.AvailabilityZones[0].ZoneName

		pod := testsuites.PodDetails{
			Cmd: testsuites.PodCmdWriteToVolume("/mnt/test-1"),
			Volumes: []testsuites.VolumeDetails{
				{
					CreateVolumeParameters: map[string]string{
						ebscsidriver.VolumeTypeKey: awscloud.VolumeTypeGP2,
						ebscsidriver.FSTypeKey:     ebscsidriver.FSTypeExt4,
					},
					ClaimSize:   driver.MinimumSizeForVolumeType(awscloud.VolumeTypeGP2),
					VolumeMount: testsuites.DefaultGeneratedVolumeMount,
				},
			},
		}
		restoredPod := testsuites.PodDetails{
			Cmd: testsuites.PodCmdGrepVolumeData("/mnt/test-1"),
			Volumes: []testsuites.VolumeDetails{
				{
					CreateVolumeParameters: map[string]string{
						ebscsidriver.VolumeTypeKey: awscloud.VolumeTypeGP2,
						ebscsidriver.FSTypeKey:     ebscsidriver.FSTypeExt4,
					},
					ClaimSize:   driver.MinimumSizeForVolumeType(awscloud.VolumeTypeGP2),
					VolumeMount: testsuites.DefaultGeneratedVolumeMount,
				},
			},
		}
		test := testsuites.DynamicallyProvisionedVolumeSnapshotTest{
			CSIDriver:   ebsDriver,
			Pod:         pod,
			RestoredPod: restoredPod,
			Parameters: map[string]string{
				ebscsidriver.FastSnapshotRestoreAvailabilityZones: fsrAvailabilityZone,
			},
			ValidateFunc: func(snapshot *volumesnapshotv1.VolumeSnapshot) {
				describeResult := validateEc2Snapshot(ec2Client, &ec2.DescribeSnapshotsInput{
					Filters: []*ec2.Filter{
						{
							Name:   aws.String("tag:" + awscloud.SnapshotNameTagKey),
							Values: []*string{aws.String("snapshot-" + string(snapshot.UID))},
						},
					},
				})

				result, err := ec2Client.DescribeFastSnapshotRestores(&ec2.DescribeFastSnapshotRestoresInput{
					Filters: []*ec2.Filter{
						{
							Name:   aws.String("snapshot-id"),
							Values: []*string{describeResult.Snapshots[0].SnapshotId},
						},
					},
				})
				if err != nil {
					Fail(fmt.Sprintf("failed to list AZs: %v", err))
				}

				if len(result.FastSnapshotRestores) != 1 {
					Fail(fmt.Sprintf("expected 1 FSR, got %d", len(result.FastSnapshotRestores)))
				}

				if *result.FastSnapshotRestores[0].AvailabilityZone != fsrAvailabilityZone {
					Fail(fmt.Sprintf("expected FSR to be enabled for %s, got %s", fsrAvailabilityZone, *result.FastSnapshotRestores[0].AvailabilityZone))
				}
			},
		}
		test.Run(cs, snapshotrcs, ns)
	})
})

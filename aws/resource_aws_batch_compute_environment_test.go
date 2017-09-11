package aws

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAWSBatchComputeEnvironment_createEc2(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSBatchComputeEnvironmentConfigEC2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
				),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_createEc2WithTags(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSBatchComputeEnvironmentConfigEC2WithTags,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
					resource.TestCheckResourceAttr("aws_batch_compute_environment.ec2", "compute_resources.0.tags.%", "1"),
					resource.TestCheckResourceAttr("aws_batch_compute_environment.ec2", "compute_resources.0.tags.Key1", "Value1"),
				),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_createSpot(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSBatchComputeEnvironmentConfigSpot,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
				),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_createUnmanaged(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSBatchComputeEnvironmentConfigUnmanaged,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
				),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_updateMaxvCpus(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSBatchComputeEnvironmentConfigEC2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
					resource.TestCheckResourceAttr("aws_batch_compute_environment.ec2", "compute_resources.0.max_vcpus", "16"),
				),
			},
			{
				Config: testAccAWSBatchComputeEnvironmentConfigEC2UpdateMaxvCpus,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
					resource.TestCheckResourceAttr("aws_batch_compute_environment.ec2", "compute_resources.0.max_vcpus", "32"),
				),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_updateInstanceType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSBatchComputeEnvironmentConfigEC2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
					resource.TestCheckResourceAttr("aws_batch_compute_environment.ec2", "compute_resources.0.instance_type.#", "1"),
				),
			},
			{
				Config: testAccAWSBatchComputeEnvironmentConfigEC2UpdateInstanceType,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
					resource.TestCheckResourceAttr("aws_batch_compute_environment.ec2", "compute_resources.0.instance_type.#", "2"),
				),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_updateComputeEnvironmentName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSBatchComputeEnvironmentConfigEC2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
					resource.TestCheckResourceAttr("aws_batch_compute_environment.ec2", "compute_environment_name", "sample"),
				),
			},
			{
				Config: testAccAWSBatchComputeEnvironmentConfigEC2UpdateComputeEnvironmentName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
					resource.TestCheckResourceAttr("aws_batch_compute_environment.ec2", "compute_environment_name", "sample_updated"),
				),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_createEc2WithoutComputeResources(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccAWSBatchComputeEnvironmentConfigEC2WithoutComputeResources,
				ExpectError: regexp.MustCompile(`One compute environment is expected, but no compute environments are set`),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_createUnmanagedWithComputeResources(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSBatchComputeEnvironmentConfigUnmanagedWithComputeResources,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
					resource.TestCheckResourceAttr("aws_batch_compute_environment.unmanaged", "type", "UNMANAGED"),
				),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_createSpotWithoutBidPercentage(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccAWSBatchComputeEnvironmentConfigSpotWithoutBidPercentage,
				ExpectError: regexp.MustCompile(`ComputeResources.spotIamFleetRole cannot not be null or empty`),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_createEc2WithGoodComputeEnvironmentName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAWSBatchComputeEnvironmentConfigEC2WithGoodComputeEnvironmentName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsBatchComputeEnvironmentExists(),
				),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_createEc2WithBadComputeEnvironmentName1(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccAWSBatchComputeEnvironmentConfigEC2WithBadComputeEnvironmentName1,
				ExpectError: regexp.MustCompile(`computeEnvironmentName must be up to 128 letters \(uppercase and lowercase\), numbers, and underscores.`),
			},
		},
	})
}

func TestAccAWSBatchComputeEnvironment_createEc2WithBadComputeEnvironmentName2(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBatchComputeEnvironmentDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccAWSBatchComputeEnvironmentConfigEC2WithBadComputeEnvironmentName2,
				ExpectError: regexp.MustCompile(`computeEnvironmentName must be up to 128 letters \(uppercase and lowercase\), numbers, and underscores.`),
			},
		},
	})
}

func testAccCheckBatchComputeEnvironmentDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*AWSClient).batchconn

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_batch_compute_environment" {
			continue
		}

		result, err := conn.DescribeComputeEnvironments(&batch.DescribeComputeEnvironmentsInput{
			ComputeEnvironments: []*string{
				aws.String(rs.Primary.ID),
			},
		})

		if err != nil {
			return fmt.Errorf("Error occured when get compute environment information.")
		}
		if len(result.ComputeEnvironments) == 1 {
			return fmt.Errorf("Compute environment still exists.")
		}

	}

	return nil
}

func testAccCheckAwsBatchComputeEnvironmentExists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*AWSClient).batchconn

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_batch_compute_environment" {
				continue
			}

			result, err := conn.DescribeComputeEnvironments(&batch.DescribeComputeEnvironmentsInput{
				ComputeEnvironments: []*string{
					aws.String(rs.Primary.ID),
				},
			})

			if err != nil {
				return fmt.Errorf("Error occured when get compute environment information.")
			}
			if len(result.ComputeEnvironments) == 0 {
				return fmt.Errorf("Compute environment doesn't exists.")
			} else if len(result.ComputeEnvironments) >= 2 {
				return fmt.Errorf("Too many compute environments exist.")
			}
		}

		return nil
	}
}

const testAccAWSBatchComputeEnvironmentConfigBase = `

########## ecs_instance_role ##########

resource "aws_iam_role" "ecs_instance_role" {
  name = "ecs_instance_role"
  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
	{
	    "Action": "sts:AssumeRole",
	    "Effect": "Allow",
	    "Principal": {
		"Service": "ec2.amazonaws.com"
	    }
	}
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "ecs_instance_role" {
  role       = "${aws_iam_role.ecs_instance_role.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
}

resource "aws_iam_instance_profile" "ecs_instance_role" {
  name  = "ecs_instance_role"
  role = "${aws_iam_role.ecs_instance_role.name}"
}

########## aws_batch_service_role ##########

resource "aws_iam_role" "aws_batch_service_role" {
  name = "aws_batch_service_role"
  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
	{
	    "Action": "sts:AssumeRole",
	    "Effect": "Allow",
	    "Principal": {
		"Service": "batch.amazonaws.com"
	    }
	}
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "aws_batch_service_role" {
  role       = "${aws_iam_role.aws_batch_service_role.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole"
}

########## aws_ec2_spot_fleet_role ##########

resource "aws_iam_role" "aws_ec2_spot_fleet_role" {
  name = "aws_ec2_spot_fleet_role"
  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
	{
	    "Action": "sts:AssumeRole",
	    "Effect": "Allow",
	    "Principal": {
		"Service": "spotfleet.amazonaws.com"
	    }
	}
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "aws_ec2_spot_fleet_role" {
  role       = "${aws_iam_role.aws_ec2_spot_fleet_role.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2SpotFleetRole"
}

########## security group ##########

resource "aws_security_group" "test_acc" {
  name = "aws_batch_compute_environment_security_group"
}

########## subnets ##########

resource "aws_vpc" "test_acc" {
  cidr_block = "10.1.0.0/16"
}

resource "aws_subnet" "test_acc" {
  vpc_id = "${aws_vpc.test_acc.id}"
  cidr_block = "10.1.1.0/24"
}
`

const testAccAWSBatchComputeEnvironmentConfigEC2 = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "ec2" {
  compute_environment_name = "sample"
  compute_resources {
    instance_role = "${aws_iam_instance_profile.ecs_instance_role.arn}"
    instance_type = [
      "c4.large",
    ]
    max_vcpus = 16
    min_vcpus = 0
    security_group_ids = [
      "${aws_security_group.test_acc.id}"
    ]
    subnets = [
      "${aws_subnet.test_acc.id}"
    ]
    type = "EC2"
  }
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "MANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigEC2WithTags = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "ec2" {
  compute_environment_name = "sample"
  compute_resources {
    instance_role = "${aws_iam_instance_profile.ecs_instance_role.arn}"
    instance_type = [
      "c4.large",
    ]
    max_vcpus = 16
    min_vcpus = 0
    security_group_ids = [
      "${aws_security_group.test_acc.id}"
    ]
    subnets = [
      "${aws_subnet.test_acc.id}"
    ]
    type = "EC2"
    tags {
      Key1 = "Value1"
    }
  }
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "MANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigSpot = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "spot" {
  compute_environment_name = "sample"
  compute_resources {
    bid_percentage = 100
    instance_role = "${aws_iam_instance_profile.ecs_instance_role.arn}"
    instance_type = [
      "c4.large",
    ]
    max_vcpus = 16
    min_vcpus = 0
    security_group_ids = [
      "${aws_security_group.test_acc.id}"
    ]
    spot_iam_fleet_role = "${aws_iam_role.aws_ec2_spot_fleet_role.arn}"
    subnets = [
      "${aws_subnet.test_acc.id}"
    ]
    type = "SPOT"
  }
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "MANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigUnmanaged = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "unmanaged" {
  compute_environment_name = "sample"
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "UNMANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigEC2UpdateMaxvCpus = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "ec2" {
  compute_environment_name = "sample"
  compute_resources {
    instance_role = "${aws_iam_instance_profile.ecs_instance_role.arn}"
    instance_type = [
      "c4.large",
    ]
    max_vcpus = 32
    min_vcpus = 0
    security_group_ids = [
      "${aws_security_group.test_acc.id}"
    ]
    subnets = [
      "${aws_subnet.test_acc.id}"
    ]
    type = "EC2"
  }
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "MANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigEC2UpdateInstanceType = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "ec2" {
  compute_environment_name = "sample"
  compute_resources {
    instance_role = "${aws_iam_instance_profile.ecs_instance_role.arn}"
    instance_type = [
      "c4.large",
      "c4.xlarge",
    ]
    max_vcpus = 16
    min_vcpus = 0
    security_group_ids = [
      "${aws_security_group.test_acc.id}"
    ]
    subnets = [
      "${aws_subnet.test_acc.id}"
    ]
    type = "EC2"
  }
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "MANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigEC2UpdateComputeEnvironmentName = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "ec2" {
  compute_environment_name = "sample_updated"
  compute_resources {
    instance_role = "${aws_iam_instance_profile.ecs_instance_role.arn}"
    instance_type = [
      "c4.large",
    ]
    max_vcpus = 16
    min_vcpus = 0
    security_group_ids = [
      "${aws_security_group.test_acc.id}"
    ]
    subnets = [
      "${aws_subnet.test_acc.id}"
    ]
    type = "EC2"
  }
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "MANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigEC2WithoutComputeResources = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "ec2" {
  compute_environment_name = "sample"
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "MANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigUnmanagedWithComputeResources = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "unmanaged" {
  compute_environment_name = "sample"
  compute_resources {
    instance_role = "${aws_iam_instance_profile.ecs_instance_role.arn}"
    instance_type = [
      "c4.large",
    ]
    max_vcpus = 16
    min_vcpus = 0
    security_group_ids = [
      "${aws_security_group.test_acc.id}"
    ]
    subnets = [
      "${aws_subnet.test_acc.id}"
    ]
    type = "EC2"
  }
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "UNMANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigSpotWithoutBidPercentage = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "ec2" {
  compute_environment_name = "sample"
  compute_resources {
    instance_role = "${aws_iam_instance_profile.ecs_instance_role.arn}"
    instance_type = [
      "c4.large",
    ]
    max_vcpus = 16
    min_vcpus = 0
    security_group_ids = [
      "${aws_security_group.test_acc.id}"
    ]
    subnets = [
      "${aws_subnet.test_acc.id}"
    ]
    type = "SPOT"
  }
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "MANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigEC2WithGoodComputeEnvironmentName = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "ec2" {
  compute_environment_name = "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678"
  compute_resources {
    instance_role = "${aws_iam_instance_profile.ecs_instance_role.arn}"
    instance_type = [
      "c4.large",
    ]
    max_vcpus = 16
    min_vcpus = 0
    security_group_ids = [
      "${aws_security_group.test_acc.id}"
    ]
    subnets = [
      "${aws_subnet.test_acc.id}"
    ]
    type = "EC2"
  }
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "MANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigEC2WithBadComputeEnvironmentName1 = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "ec2" {
  compute_environment_name = "sam@ple"
  compute_resources {
    instance_role = "${aws_iam_instance_profile.ecs_instance_role.arn}"
    instance_type = [
      "c4.large",
    ]
    max_vcpus = 16
    min_vcpus = 0
    security_group_ids = [
      "${aws_security_group.test_acc.id}"
    ]
    subnets = [
      "${aws_subnet.test_acc.id}"
    ]
    type = "EC2"
  }
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "MANAGED"
}
`

const testAccAWSBatchComputeEnvironmentConfigEC2WithBadComputeEnvironmentName2 = testAccAWSBatchComputeEnvironmentConfigBase + `
resource "aws_batch_compute_environment" "ec2" {
  compute_environment_name = "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"
  compute_resources {
    instance_role = "${aws_iam_instance_profile.ecs_instance_role.arn}"
    instance_type = [
      "c4.large",
    ]
    max_vcpus = 16
    min_vcpus = 0
    security_group_ids = [
      "${aws_security_group.test_acc.id}"
    ]
    subnets = [
      "${aws_subnet.test_acc.id}"
    ]
    type = "EC2"
  }
  service_role = "${aws_iam_role.aws_batch_service_role.arn}"
  type = "MANAGED"
}
`

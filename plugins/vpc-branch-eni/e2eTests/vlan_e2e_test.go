// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// +build e2e_test

package e2e

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/aws/amazon-vpc-cni-plugins/network/netns"
	"github.com/containernetworking/cni/pkg/invoke"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	containerID = "container-id"
	nsName      = "testNS"
	ifName      = "testIfName"
	netConfFormat                   = `
{
    "type":"vpc-branch-eni",
    "cniVersion":"0.3.0",
    "trunkMacAddress":"%s",
    "branchVlanID":"%s",
    "branchMACAddress":"%s",
    "branchIPAddress":"%s",
    "branchGatewayIPAddress":"%s"
}`
)

func TestAddDelVLAN(t *testing.T) {
	// Ensure that the cni plugin exists.
	pluginPath, err := invoke.FindInPath("vpc-branch-eni", []string{os.Getenv("CNI_PATH")})
	require.NoError(t, err, "Unable to find vpc-branch-eni plugin in path")

	// Create a directory for storing test logs.
	testLogDir, err := ioutil.TempDir("", "vpc-branch-eni-cni-e2eTests-test-")
	require.NoError(t, err, "Unable to create directory for storing test logs")

	trunkMACAddress := os.Getenv("TRUNK_MAC_ADDRESS")
	require.NotEmpty(t, trunkMACAddress, "TRUNK_MAC_ADDRESS needs to be specified")

	branchVlanID := os.Getenv("BRANCH_VLAN_ID")
	require.NotEmpty(t, branchVlanID, "BRANCH_VLAN_ID needs to be specified")

	branchMACAddress := os.Getenv("BRANCH_MAC_ADDRESS")
	require.NotEmpty(t, branchMACAddress, "BRANCH_MAC_ADDRESS needs to be specified")

	branchIPAddress := os.Getenv("BRANCH_IP_ADDRESS")
	require.NotEmpty(t, branchIPAddress, "BRANCH_IP_ADDRESS needs to be specified")

	branchGatewayIPAddress := os.Getenv("BRANCH_GATEWAY_IP_ADDRESS")
	require.NotEmpty(t, branchGatewayIPAddress, "BRANCH_GATEWAY_IP_ADDRESS needs to be specified")

	netConf := fmt.Sprintf(netConfFormat, trunkMACAddress, branchVlanID, branchMACAddress, branchIPAddress, branchGatewayIPAddress)

	// Configure the env var to use the test logs directory.
	os.Setenv("CNI_LOG_FILE", fmt.Sprintf("%s/vpc-branch-eni.log", testLogDir))
	t.Logf("Using %s for test logs", testLogDir)
	defer os.Unsetenv("CNI_LOG_FILE")

	// Handle deletion of test logs at the end of the test execution if specified.
	ok, err := strconv.ParseBool(getEnvOrDefault("ECS_PRESERVE_E2E_TEST_LOGS", "false"))
	assert.NoError(t, err, "Unable to parse ECS_PRESERVE_E2E_TEST_LOGS env var")
	defer func(preserve bool) {
		if !t.Failed() && !preserve {
			os.RemoveAll(testLogDir)
		}
	}(ok)

	// Create a network namespace to mimic the container's network namespace.
	targetNS, err := netns.NewNetNS(nsName)
	require.NoError(t, err,
		"Unable to create the network namespace that represents the network namespace of the container")
	defer targetNS.Close()

	// Construct args to invoke the CNI plugin with.
	//execInvokeArgs := &invoke.Args{
	//	ContainerID: containerID,
	//	NetNS:       targetNS.GetPath(),
	//	IfName:      ifName,
	//	Path:        os.Getenv("CNI_PATH"),
	//}

	t.Logf("Got config: %s", netConf)
	t.Logf("Plugin path: %s", pluginPath)
}

// getEnvOrDefault gets the value of an env var. It returns the fallback value
// if the env var is not set.
func getEnvOrDefault(name string, fallback string) string {
	val := os.Getenv(name)
	if val == "" {
		return fallback
	}

	return val
}

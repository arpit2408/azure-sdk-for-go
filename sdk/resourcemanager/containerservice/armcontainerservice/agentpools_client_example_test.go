//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsAbortOperation.json
func ExampleAgentPoolsClient_BeginAbortLatestOperation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginAbortLatestOperation(ctx, "rg1", "clustername1", "agentpool1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsList.json
func ExampleAgentPoolsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("rg1", "clustername1", nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsGet.json
func ExampleAgentPoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "rg1", "clustername1", "agentpool1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsAssociate_CRG.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_associateAgentPoolWithCapacityReservationGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			CapacityReservationGroupID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.Compute/CapacityReservationGroups/crg1"),
			Count:                      to.Ptr[int32](3),
			OrchestratorVersion:        to.Ptr(""),
			OSType:                     to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:                     to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_Snapshot.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolUsingAnAgentPoolSnapshot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count: to.Ptr[int32](3),
			CreationData: &armcontainerservice.CreationData{
				SourceResourceID: to.Ptr("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.ContainerService/snapshots/snapshot1"),
			},
			EnableFIPS:          to.Ptr(true),
			OrchestratorVersion: to.Ptr(""),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:              to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_EnableCustomCATrust.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithCustomCaTrustEnabled() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:               to.Ptr[int32](3),
			EnableCustomCATrust: to.Ptr(true),
			OrchestratorVersion: to.Ptr(""),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:              to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_DedicatedHostGroup.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithDedicatedHostGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:               to.Ptr[int32](3),
			HostGroupID:         to.Ptr("/subscriptions/subid1/resourcegroups/rg/providers/Microsoft.Compute/hostGroups/hostgroup1"),
			OrchestratorVersion: to.Ptr(""),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:              to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_EnableEncryptionAtHost.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithEncryptionAtHostEnabled() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:                  to.Ptr[int32](3),
			EnableEncryptionAtHost: to.Ptr(true),
			OrchestratorVersion:    to.Ptr(""),
			OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:                 to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_Ephemeral.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithEphemeralOsDisk() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:               to.Ptr[int32](3),
			OrchestratorVersion: to.Ptr(""),
			OSDiskSizeGB:        to.Ptr[int32](64),
			OSDiskType:          to.Ptr(armcontainerservice.OSDiskTypeEphemeral),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:              to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_EnableFIPS.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithFipsEnabledOs() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:               to.Ptr[int32](3),
			EnableFIPS:          to.Ptr(true),
			OrchestratorVersion: to.Ptr(""),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:              to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_GPUMIG.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithGpumig() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:              to.Ptr[int32](3),
			GpuInstanceProfile: to.Ptr(armcontainerservice.GPUInstanceProfileMIG2G),
			KubeletConfig: &armcontainerservice.KubeletConfig{
				AllowedUnsafeSysctls: []*string{
					to.Ptr("kernel.msg*"),
					to.Ptr("net.core.somaxconn")},
				CPUCfsQuota:           to.Ptr(true),
				CPUCfsQuotaPeriod:     to.Ptr("200ms"),
				CPUManagerPolicy:      to.Ptr("static"),
				FailSwapOn:            to.Ptr(false),
				ImageGcHighThreshold:  to.Ptr[int32](90),
				ImageGcLowThreshold:   to.Ptr[int32](70),
				TopologyManagerPolicy: to.Ptr("best-effort"),
			},
			LinuxOSConfig: &armcontainerservice.LinuxOSConfig{
				SwapFileSizeMB: to.Ptr[int32](1500),
				Sysctls: &armcontainerservice.SysctlConfig{
					KernelThreadsMax:        to.Ptr[int32](99999),
					NetCoreWmemDefault:      to.Ptr[int32](12345),
					NetIPv4IPLocalPortRange: to.Ptr("20000 60000"),
					NetIPv4TCPTwReuse:       to.Ptr(true),
				},
				TransparentHugePageDefrag:  to.Ptr("madvise"),
				TransparentHugePageEnabled: to.Ptr("always"),
			},
			OrchestratorVersion: to.Ptr(""),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:              to.Ptr("Standard_ND96asr_v4"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_WasmWasi.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithKrustletAndTheWasiRuntime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:               to.Ptr[int32](3),
			Mode:                to.Ptr(armcontainerservice.AgentPoolModeUser),
			OrchestratorVersion: to.Ptr(""),
			OSDiskSizeGB:        to.Ptr[int32](64),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:              to.Ptr("Standard_DS2_v2"),
			WorkloadRuntime:     to.Ptr(armcontainerservice.WorkloadRuntimeWasmWasi),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_CustomNodeConfig.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithKubeletConfigAndLinuxOsConfig() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count: to.Ptr[int32](3),
			KubeletConfig: &armcontainerservice.KubeletConfig{
				AllowedUnsafeSysctls: []*string{
					to.Ptr("kernel.msg*"),
					to.Ptr("net.core.somaxconn")},
				CPUCfsQuota:           to.Ptr(true),
				CPUCfsQuotaPeriod:     to.Ptr("200ms"),
				CPUManagerPolicy:      to.Ptr("static"),
				FailSwapOn:            to.Ptr(false),
				ImageGcHighThreshold:  to.Ptr[int32](90),
				ImageGcLowThreshold:   to.Ptr[int32](70),
				TopologyManagerPolicy: to.Ptr("best-effort"),
			},
			LinuxOSConfig: &armcontainerservice.LinuxOSConfig{
				SwapFileSizeMB: to.Ptr[int32](1500),
				Sysctls: &armcontainerservice.SysctlConfig{
					KernelThreadsMax:        to.Ptr[int32](99999),
					NetCoreWmemDefault:      to.Ptr[int32](12345),
					NetIPv4IPLocalPortRange: to.Ptr("20000 60000"),
					NetIPv4TCPTwReuse:       to.Ptr(true),
				},
				TransparentHugePageDefrag:  to.Ptr("madvise"),
				TransparentHugePageEnabled: to.Ptr("always"),
			},
			OrchestratorVersion: to.Ptr(""),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:              to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_MessageOfTheDay.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithMessageOfTheDay() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:               to.Ptr[int32](3),
			MessageOfTheDay:     to.Ptr("Zm9vCg=="),
			Mode:                to.Ptr(armcontainerservice.AgentPoolModeUser),
			OrchestratorVersion: to.Ptr(""),
			OSDiskSizeGB:        to.Ptr[int32](64),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:              to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_OSSKU.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithOssku() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count: to.Ptr[int32](3),
			KubeletConfig: &armcontainerservice.KubeletConfig{
				AllowedUnsafeSysctls: []*string{
					to.Ptr("kernel.msg*"),
					to.Ptr("net.core.somaxconn")},
				CPUCfsQuota:           to.Ptr(true),
				CPUCfsQuotaPeriod:     to.Ptr("200ms"),
				CPUManagerPolicy:      to.Ptr("static"),
				FailSwapOn:            to.Ptr(false),
				ImageGcHighThreshold:  to.Ptr[int32](90),
				ImageGcLowThreshold:   to.Ptr[int32](70),
				TopologyManagerPolicy: to.Ptr("best-effort"),
			},
			LinuxOSConfig: &armcontainerservice.LinuxOSConfig{
				SwapFileSizeMB: to.Ptr[int32](1500),
				Sysctls: &armcontainerservice.SysctlConfig{
					KernelThreadsMax:        to.Ptr[int32](99999),
					NetCoreWmemDefault:      to.Ptr[int32](12345),
					NetIPv4IPLocalPortRange: to.Ptr("20000 60000"),
					NetIPv4TCPTwReuse:       to.Ptr(true),
				},
				TransparentHugePageDefrag:  to.Ptr("madvise"),
				TransparentHugePageEnabled: to.Ptr("always"),
			},
			OrchestratorVersion: to.Ptr(""),
			OSSKU:               to.Ptr(armcontainerservice.OSSKUCBLMariner),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:              to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_PPG.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithPpg() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:                     to.Ptr[int32](3),
			OrchestratorVersion:       to.Ptr(""),
			OSType:                    to.Ptr(armcontainerservice.OSTypeLinux),
			ProximityPlacementGroupID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.Compute/proximityPlacementGroups/ppg1"),
			VMSize:                    to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_EnableUltraSSD.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithUltraSsdEnabled() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:               to.Ptr[int32](3),
			EnableUltraSSD:      to.Ptr(true),
			OrchestratorVersion: to.Ptr(""),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:              to.Ptr("Standard_DS2_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_WindowsOSSKU.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithWindowsOssku() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "wnp2", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:               to.Ptr[int32](3),
			OrchestratorVersion: to.Ptr("1.23.3"),
			OSSKU:               to.Ptr(armcontainerservice.OSSKUWindows2022),
			OSType:              to.Ptr(armcontainerservice.OSTypeWindows),
			VMSize:              to.Ptr("Standard_D4s_v3"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_Spot.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createSpotAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count: to.Ptr[int32](3),
			NodeLabels: map[string]*string{
				"key1": to.Ptr("val1"),
			},
			NodeTaints: []*string{
				to.Ptr("Key1=Value1:NoSchedule")},
			OrchestratorVersion:    to.Ptr(""),
			OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
			ScaleSetEvictionPolicy: to.Ptr(armcontainerservice.ScaleSetEvictionPolicyDelete),
			ScaleSetPriority:       to.Ptr(armcontainerservice.ScaleSetPrioritySpot),
			Tags: map[string]*string{
				"name1": to.Ptr("val1"),
			},
			VMSize: to.Ptr("Standard_DS1_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_WindowsDisableOutboundNAT.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createWindowsAgentPoolWithDisablingOutboundNat() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "wnp2", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:               to.Ptr[int32](3),
			OrchestratorVersion: to.Ptr("1.23.8"),
			OSSKU:               to.Ptr(armcontainerservice.OSSKUWindows2022),
			OSType:              to.Ptr(armcontainerservice.OSTypeWindows),
			VMSize:              to.Ptr("Standard_D4s_v3"),
			WindowsProfile: &armcontainerservice.AgentPoolWindowsProfile{
				DisableOutboundNat: to.Ptr(true),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsCreate_Update.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createUpdateAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count: to.Ptr[int32](3),
			Mode:  to.Ptr(armcontainerservice.AgentPoolModeUser),
			NodeLabels: map[string]*string{
				"key1": to.Ptr("val1"),
			},
			NodeTaints: []*string{
				to.Ptr("Key1=Value1:NoSchedule")},
			OrchestratorVersion:    to.Ptr(""),
			OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
			ScaleSetEvictionPolicy: to.Ptr(armcontainerservice.ScaleSetEvictionPolicyDelete),
			ScaleSetPriority:       to.Ptr(armcontainerservice.ScaleSetPrioritySpot),
			Tags: map[string]*string{
				"name1": to.Ptr("val1"),
			},
			VMSize: to.Ptr("Standard_DS1_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPools_Start.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_startAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			PowerState: &armcontainerservice.PowerState{
				Code: to.Ptr(armcontainerservice.CodeRunning),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPools_Stop.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_stopAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			PowerState: &armcontainerservice.PowerState{
				Code: to.Ptr(armcontainerservice.CodeStopped),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPools_Update.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_updateAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Count:             to.Ptr[int32](3),
			EnableAutoScaling: to.Ptr(true),
			MaxCount:          to.Ptr[int32](2),
			MinCount:          to.Ptr[int32](2),
			NodeTaints: []*string{
				to.Ptr("Key1=Value1:NoSchedule")},
			OrchestratorVersion:    to.Ptr(""),
			OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
			ScaleSetEvictionPolicy: to.Ptr(armcontainerservice.ScaleSetEvictionPolicyDelete),
			ScaleSetPriority:       to.Ptr(armcontainerservice.ScaleSetPrioritySpot),
			VMSize:                 to.Ptr("Standard_DS1_v2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsDelete.json
func ExampleAgentPoolsClient_BeginDelete_deleteAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "rg1", "clustername1", "agentpool1", &armcontainerservice.AgentPoolsClientBeginDeleteOptions{IgnorePodDisruptionBudget: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsDelete_IgnorePodDisruptionBudget.json
func ExampleAgentPoolsClient_BeginDelete_deleteAgentPoolByIgnoringPodDisruptionBudget() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx, "rg1", "clustername1", "agentpool1", &armcontainerservice.AgentPoolsClientBeginDeleteOptions{IgnorePodDisruptionBudget: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsGetUpgradeProfile.json
func ExampleAgentPoolsClient_GetUpgradeProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetUpgradeProfile(ctx, "rg1", "clustername1", "agentpool1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsGetAgentPoolAvailableVersions.json
func ExampleAgentPoolsClient_GetAvailableAgentPoolVersions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetAvailableAgentPoolVersions(ctx, "rg1", "clustername1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-11-02-preview/examples/AgentPoolsUpgradeNodeImageVersion.json
func ExampleAgentPoolsClient_BeginUpgradeNodeImageVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewAgentPoolsClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpgradeNodeImageVersion(ctx, "rg1", "clustername1", "agentpool1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

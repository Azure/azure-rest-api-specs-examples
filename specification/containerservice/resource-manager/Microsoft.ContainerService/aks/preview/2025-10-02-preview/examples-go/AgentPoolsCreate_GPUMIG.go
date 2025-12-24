package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/455d20a5e76d8184f7cff960501a57e1f88986b7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/AgentPoolsCreate_GPUMIG.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithGpumig() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentPoolsClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
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
	}, &armcontainerservice.AgentPoolsClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgentPool = armcontainerservice.AgentPool{
	// 	Name: to.Ptr("agentpool1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/agentPools"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/agentpool1"),
	// 	Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
	// 		Count: to.Ptr[int32](3),
	// 		CurrentOrchestratorVersion: to.Ptr("1.17.8"),
	// 		GpuInstanceProfile: to.Ptr(armcontainerservice.GPUInstanceProfileMIG2G),
	// 		KubeletConfig: &armcontainerservice.KubeletConfig{
	// 			AllowedUnsafeSysctls: []*string{
	// 				to.Ptr("kernel.msg*"),
	// 				to.Ptr("net.core.somaxconn")},
	// 				CPUCfsQuota: to.Ptr(true),
	// 				CPUCfsQuotaPeriod: to.Ptr("200ms"),
	// 				CPUManagerPolicy: to.Ptr("static"),
	// 				FailSwapOn: to.Ptr(false),
	// 				ImageGcHighThreshold: to.Ptr[int32](90),
	// 				ImageGcLowThreshold: to.Ptr[int32](70),
	// 				TopologyManagerPolicy: to.Ptr("best-effort"),
	// 			},
	// 			LinuxOSConfig: &armcontainerservice.LinuxOSConfig{
	// 				SwapFileSizeMB: to.Ptr[int32](1500),
	// 				Sysctls: &armcontainerservice.SysctlConfig{
	// 					KernelThreadsMax: to.Ptr[int32](99999),
	// 					NetCoreWmemDefault: to.Ptr[int32](12345),
	// 					NetIPv4IPLocalPortRange: to.Ptr("20000 60000"),
	// 					NetIPv4TCPTwReuse: to.Ptr(true),
	// 				},
	// 				TransparentHugePageDefrag: to.Ptr("madvise"),
	// 				TransparentHugePageEnabled: to.Ptr("always"),
	// 			},
	// 			MaxPods: to.Ptr[int32](110),
	// 			OrchestratorVersion: to.Ptr("1.17.8"),
	// 			OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			VMSize: to.Ptr("Standard_ND96asr_v4"),
	// 		},
	// 	}
}

package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-10-02-preview/examples/AgentPoolsCreate_CustomNodeConfig.json
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

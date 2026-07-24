package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v3"
)

// Generated from example definition: 2026-06-16-preview/extension/Extension_Add.json
func ExampleManagementClient_BeginSetupExtensions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcompute.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagementClient().BeginSetupExtensions(ctx, "myResourceGroup", "myMachine", armhybridcompute.SetupExtensionRequest{
		Extensions: []*armhybridcompute.MachineExtensionProperties{
			{
				Type:      to.Ptr("AzureMonitorAgentLinux"),
				Publisher: to.Ptr("Microsoft.Azure.Monitoring"),
			},
			{
				Type:      to.Ptr("<extension_type>"),
				Publisher: to.Ptr("<extension_publisher>"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armhybridcompute.ManagementClientSetupExtensionsResponse{
	// 	SetupExtensionRequest: armhybridcompute.SetupExtensionRequest{
	// 	},
	// }
}

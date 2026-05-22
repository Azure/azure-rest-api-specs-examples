package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/SessionPools_McpServer_CreateOrUpdate.json
func ExampleContainerAppsSessionPoolsClient_BeginCreateOrUpdate_createOrUpdateSessionPoolWithMcpServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContainerAppsSessionPoolsClient().BeginCreateOrUpdate(ctx, "rg", "testsessionpool", armappcontainers.SessionPool{
		Location: to.Ptr("East US"),
		Properties: &armappcontainers.SessionPoolProperties{
			ContainerType: to.Ptr(armappcontainers.ContainerTypeShell),
			DynamicPoolConfiguration: &armappcontainers.DynamicPoolConfiguration{
				LifecycleConfiguration: &armappcontainers.LifecycleConfiguration{
					CooldownPeriodInSeconds: to.Ptr[int32](600),
					LifecycleType:           to.Ptr(armappcontainers.LifecycleTypeTimed),
				},
			},
			McpServerSettings: &armappcontainers.McpServerSettings{
				IsMcpServerAPIKeyDisabled: to.Ptr(false),
				IsMcpServerEnabled:        to.Ptr(true),
			},
			PoolManagementType: to.Ptr(armappcontainers.PoolManagementTypeDynamic),
			ScaleConfiguration: &armappcontainers.ScaleConfiguration{
				MaxConcurrentSessions: to.Ptr[int32](50),
			},
			SessionNetworkConfiguration: &armappcontainers.SessionNetworkConfiguration{
				Status: to.Ptr(armappcontainers.SessionNetworkStatusEgressEnabled),
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
	// res = armappcontainers.ContainerAppsSessionPoolsClientCreateOrUpdateResponse{
	// 	SessionPool: armappcontainers.SessionPool{
	// 		Name: to.Ptr("testsessionpool"),
	// 		Type: to.Ptr("Microsoft.App/sessionPools"),
	// 		ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/sessionPools/testsessionpool"),
	// 		Location: to.Ptr("East US"),
	// 		Properties: &armappcontainers.SessionPoolProperties{
	// 			ContainerType: to.Ptr(armappcontainers.ContainerTypeShell),
	// 			DynamicPoolConfiguration: &armappcontainers.DynamicPoolConfiguration{
	// 				LifecycleConfiguration: &armappcontainers.LifecycleConfiguration{
	// 					CooldownPeriodInSeconds: to.Ptr[int32](600),
	// 					LifecycleType: to.Ptr(armappcontainers.LifecycleTypeTimed),
	// 				},
	// 			},
	// 			McpServerSettings: &armappcontainers.McpServerSettings{
	// 				IsMcpServerAPIKeyDisabled: to.Ptr(false),
	// 				IsMcpServerEnabled: to.Ptr(true),
	// 				McpServerEndpoint: to.Ptr("https://eastus.dynamicsessions.io/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/sessionPools/testsessionpool/mcp"),
	// 			},
	// 			PoolManagementEndpoint: to.Ptr("https://eastus.dynamicsessions.io/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/sessionPools/testsessionpool"),
	// 			PoolManagementType: to.Ptr(armappcontainers.PoolManagementTypeDynamic),
	// 			ProvisioningState: to.Ptr(armappcontainers.SessionPoolProvisioningStateSucceeded),
	// 			ScaleConfiguration: &armappcontainers.ScaleConfiguration{
	// 				MaxConcurrentSessions: to.Ptr[int32](50),
	// 			},
	// 			SessionNetworkConfiguration: &armappcontainers.SessionNetworkConfiguration{
	// 				Status: to.Ptr(armappcontainers.SessionNetworkStatusEgressEnabled),
	// 			},
	// 		},
	// 	},
	// }
}

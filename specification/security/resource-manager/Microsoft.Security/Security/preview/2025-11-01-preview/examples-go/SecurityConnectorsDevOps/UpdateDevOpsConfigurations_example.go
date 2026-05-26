package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: 2025-11-01-preview/SecurityConnectorsDevOps/UpdateDevOpsConfigurations_example.json
func ExampleDevOpsConfigurationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurity.NewClientFactory("0806e1cd-cfda-4ff8-b99c-2b0af42cffd3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDevOpsConfigurationsClient().BeginUpdate(ctx, "myRg", "mySecurityConnectorName", armsecurity.DevOpsConfiguration{
		Properties: &armsecurity.DevOpsConfigurationProperties{
			AgentlessConfiguration: &armsecurity.AgentlessConfiguration{
				AgentlessAutoDiscovery: to.Ptr(armsecurity.AutoDiscoveryDisabled),
				AgentlessEnabled:       to.Ptr(armsecurity.AgentlessEnablementEnabled),
				InventoryList: []*armsecurity.InventoryList{
					{
						InventoryKind: to.Ptr(armsecurity.InventoryKindAzureDevOpsOrganization),
						Value:         to.Ptr("org1"),
					},
				},
				InventoryListType: to.Ptr(armsecurity.InventoryListKindInclusion),
				Scanners: []*string{
					to.Ptr("scanner1"),
					to.Ptr("scanner2"),
				},
			},
			AutoDiscovery: to.Ptr(armsecurity.AutoDiscoveryEnabled),
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
	// res = armsecurity.DevOpsConfigurationsClientUpdateResponse{
	// 	DevOpsConfiguration: armsecurity.DevOpsConfiguration{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Security/securityConnectors/devops"),
	// 		ID: to.Ptr("/subscriptions/0806e1cd-cfda-4ff8-b99c-2b0af42cffd3/resourceGroups/myRg/providers/Microsoft.Security/securityConnectors/mySecurityConnectorName/devops/default"),
	// 		Properties: &armsecurity.DevOpsConfigurationProperties{
	// 			AgentlessConfiguration: &armsecurity.AgentlessConfiguration{
	// 				AgentlessAutoDiscovery: to.Ptr(armsecurity.AutoDiscoveryDisabled),
	// 				AgentlessEnabled: to.Ptr(armsecurity.AgentlessEnablementEnabled),
	// 				InventoryList: []*armsecurity.InventoryList{
	// 					{
	// 						InventoryKind: to.Ptr(armsecurity.InventoryKindAzureDevOpsOrganization),
	// 						Value: to.Ptr("org1"),
	// 					},
	// 				},
	// 				InventoryListType: to.Ptr(armsecurity.InventoryListKindInclusion),
	// 				Scanners: []*string{
	// 					to.Ptr("scanner1"),
	// 					to.Ptr("scanner2"),
	// 				},
	// 			},
	// 			AutoDiscovery: to.Ptr(armsecurity.AutoDiscoveryEnabled),
	// 			ProvisioningState: to.Ptr(armsecurity.DevOpsProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}

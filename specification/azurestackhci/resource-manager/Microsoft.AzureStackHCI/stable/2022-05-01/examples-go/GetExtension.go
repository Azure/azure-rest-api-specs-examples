package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/GetExtension.json
func ExampleExtensionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExtensionsClient().Get(ctx, "test-rg", "myCluster", "default", "MicrosoftMonitoringAgent", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Extension = armazurestackhci.Extension{
	// 	Name: to.Ptr("MicrosoftMonitoringAgent"),
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/clusters/arcSettings/extensions"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster/arcSettings/default/extensions/MicrosoftMonitoringAgent"),
	// 	Properties: &armazurestackhci.ExtensionProperties{
	// 		AggregateState: to.Ptr(armazurestackhci.ExtensionAggregateStatePartiallySucceeded),
	// 		ExtensionParameters: &armazurestackhci.ExtensionParameters{
	// 			Type: to.Ptr("string"),
	// 			AutoUpgradeMinorVersion: to.Ptr(false),
	// 			Publisher: to.Ptr("Microsoft.Compute"),
	// 			Settings: map[string]any{
	// 				"workspaceId": "xx",
	// 			},
	// 			TypeHandlerVersion: to.Ptr("1.10.3"),
	// 		},
	// 		PerNodeExtensionDetails: []*armazurestackhci.PerNodeExtensionState{
	// 			{
	// 				Name: to.Ptr("Node-1"),
	// 				Extension: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1/Extensions/MicrosoftMonitoringAgent"),
	// 				State: to.Ptr(armazurestackhci.NodeExtensionStateSucceeded),
	// 			},
	// 			{
	// 				Name: to.Ptr("Node-2"),
	// 				Extension: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2/Extensions/MicrosoftMonitoringAgent"),
	// 				State: to.Ptr(armazurestackhci.NodeExtensionStateFailed),
	// 		}},
	// 		ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 	},
	// 	SystemData: &armazurestackhci.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 	},
	// }
}

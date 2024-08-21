package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/GetExtension.json
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
	// 	SystemData: &armazurestackhci.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 	},
	// 	Properties: &armazurestackhci.ExtensionProperties{
	// 		AggregateState: to.Ptr(armazurestackhci.ExtensionAggregateStatePartiallySucceeded),
	// 		ExtensionParameters: &armazurestackhci.ExtensionParameters{
	// 			Type: to.Ptr("string"),
	// 			AutoUpgradeMinorVersion: to.Ptr(false),
	// 			EnableAutomaticUpgrade: to.Ptr(true),
	// 			Publisher: to.Ptr("Microsoft.Compute"),
	// 			Settings: map[string]any{
	// 				"workspaceId": "xx",
	// 			},
	// 			TypeHandlerVersion: to.Ptr("1.10.3"),
	// 		},
	// 		ManagedBy: to.Ptr(armazurestackhci.ExtensionManagedByAzure),
	// 		PerNodeExtensionDetails: []*armazurestackhci.PerNodeExtensionState{
	// 			{
	// 				Name: to.Ptr("Node-1"),
	// 				Extension: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1/Extensions/MicrosoftMonitoringAgent"),
	// 				InstanceView: &armazurestackhci.ExtensionInstanceView{
	// 					Name: to.Ptr("MicrosoftMonitoringAgent"),
	// 					Type: to.Ptr("MicrosoftMonitoringAgent"),
	// 					Status: &armazurestackhci.ExtensionInstanceViewStatus{
	// 						Code: to.Ptr("success"),
	// 						DisplayStatus: to.Ptr("Provisioning succeeded"),
	// 						Level: to.Ptr(armazurestackhci.StatusLevelTypes("Information")),
	// 						Message: to.Ptr("Finished executing command, StdOut: , StdErr:"),
	// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-08T20:42:10.999Z"); return t}()),
	// 					},
	// 					TypeHandlerVersion: to.Ptr("1.10.0"),
	// 				},
	// 				State: to.Ptr(armazurestackhci.NodeExtensionStateSucceeded),
	// 				TypeHandlerVersion: to.Ptr("1.10.0"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Node-2"),
	// 				Extension: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2/Extensions/MicrosoftMonitoringAgent"),
	// 				InstanceView: &armazurestackhci.ExtensionInstanceView{
	// 					Name: to.Ptr("MicrosoftMonitoringAgent"),
	// 					Type: to.Ptr("MicrosoftMonitoringAgent"),
	// 					Status: &armazurestackhci.ExtensionInstanceViewStatus{
	// 						Code: to.Ptr("error"),
	// 						DisplayStatus: to.Ptr("Provisioning failed"),
	// 						Level: to.Ptr(armazurestackhci.StatusLevelTypesError),
	// 						Message: to.Ptr("Finished executing command, StdOut: , StdErr:"),
	// 						Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-08-08T20:42:10.999Z"); return t}()),
	// 					},
	// 					TypeHandlerVersion: to.Ptr("1.10.3"),
	// 				},
	// 				State: to.Ptr(armazurestackhci.NodeExtensionStateFailed),
	// 				TypeHandlerVersion: to.Ptr("1.10.3"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 	},
	// }
}

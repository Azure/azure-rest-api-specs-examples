package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PatchArcSetting.json
func ExampleArcSettingsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewArcSettingsClient().Update(ctx, "test-rg", "myCluster", "default", armazurestackhci.ArcSettingsPatch{
		Properties: &armazurestackhci.ArcSettingsPatchProperties{
			ConnectivityProperties: map[string]any{
				"enabled": true,
				"serviceConfigurations": []any{
					map[string]any{
						"port":        float64(6516),
						"serviceName": "WAC",
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ArcSetting = armazurestackhci.ArcSetting{
	// 	Type: to.Ptr("Microsoft.AzureStackHCI/clusters/arcSettings"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/clusters/myCluster/arcSettings/default"),
	// 	SystemData: &armazurestackhci.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 	},
	// 	Properties: &armazurestackhci.ArcSettingProperties{
	// 		AggregateState: to.Ptr(armazurestackhci.ArcSettingAggregateStateCreating),
	// 		ArcInstanceResourceGroup: to.Ptr("ArcInstance-rg"),
	// 		ConnectivityProperties: map[string]any{
	// 			"enabled": true,
	// 			"serviceConfigurations":[]any{
	// 				map[string]any{
	// 					"port": float64(6516),
	// 					"serviceName": "WAC",
	// 				},
	// 			},
	// 		},
	// 		DefaultExtensions: []*armazurestackhci.DefaultExtensionDetails{
	// 			{
	// 				Category: to.Ptr("Telemetry"),
	// 				ConsentTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-01-01T17:18:19.123Z"); return t}()),
	// 			},
	// 			{
	// 				Category: to.Ptr("Supportability"),
	// 		}},
	// 		PerNodeDetails: []*armazurestackhci.PerNodeState{
	// 			{
	// 				Name: to.Ptr("Node-1"),
	// 				ArcInstance: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1"),
	// 				State: to.Ptr(armazurestackhci.NodeArcStateCreating),
	// 			},
	// 			{
	// 				Name: to.Ptr("Node-2"),
	// 				ArcInstance: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2"),
	// 				State: to.Ptr(armazurestackhci.NodeArcStateCreating),
	// 		}},
	// 		ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 	},
	// }
}

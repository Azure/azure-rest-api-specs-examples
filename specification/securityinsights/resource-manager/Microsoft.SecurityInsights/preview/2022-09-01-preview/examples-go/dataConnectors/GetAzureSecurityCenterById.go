package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/GetAzureSecurityCenterById.json
func ExampleDataConnectorsClient_Get_getAAscDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "763f9fa1-c2d3-4fa2-93e9-bccd4899aa12", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.ASCDataConnector{
	// 		Name: to.Ptr("763f9fa1-c2d3-4fa2-93e9-bccd4899aa12"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/763f9fa1-c2d3-4fa2-93e9-bccd4899aa12"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindAzureSecurityCenter),
	// 		Properties: &armsecurityinsights.ASCDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.AlertsDataTypeOfDataConnector{
	// 				Alerts: &armsecurityinsights.DataConnectorDataTypeCommon{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			SubscriptionID: to.Ptr("c0688291-89d7-4bed-87a2-a7b1bff43f4c"),
	// 		},
	// 	},
	// 	                        }
}

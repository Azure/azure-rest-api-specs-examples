package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/GetDynamics365DataConnectorById.json
func ExampleDataConnectorsClient_Get_getADynamics365DataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "c2541efb-c9a6-47fe-9501-87d1017d1512", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.Dynamics365DataConnector{
	// 		Name: to.Ptr("c2541efb-c9a6-47fe-9501-87d1017d1512"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/3d3e955e-33eb-401d-89a7-251c81ddd660"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindDynamics365),
	// 		Properties: &armsecurityinsights.Dynamics365DataConnectorProperties{
	// 			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
	// 			DataTypes: &armsecurityinsights.Dynamics365DataConnectorDataTypes{
	// 				Dynamics365CdsActivities: &armsecurityinsights.Dynamics365DataConnectorDataTypesDynamics365CdsActivities{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 		},
	// 	},
	// 	                        }
}

package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CreateOfficePowerBIDataConnector.json
func ExampleDataConnectorsClient_CreateOrUpdate_createsOrUpdatesAnOfficePowerBiDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", &armsecurityinsights.OfficePowerBIDataConnector{
		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindOfficePowerBI),
		Properties: &armsecurityinsights.OfficePowerBIDataConnectorProperties{
			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
			DataTypes: &armsecurityinsights.OfficePowerBIConnectorDataTypes{
				Logs: &armsecurityinsights.OfficePowerBIConnectorDataTypesLogs{
					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
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
	// res = armsecurityinsights.DataConnectorsClientCreateOrUpdateResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.OfficePowerBIDataConnector{
	// 		Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindOfficePowerBI),
	// 		Properties: &armsecurityinsights.OfficePowerBIDataConnectorProperties{
	// 			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
	// 			DataTypes: &armsecurityinsights.OfficePowerBIConnectorDataTypes{
	// 				Logs: &armsecurityinsights.OfficePowerBIConnectorDataTypesLogs{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 		},
	// 	},
	// 	                        }
}

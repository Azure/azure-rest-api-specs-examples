package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c4f3c695f0250dcb261598a62004f0aef10b9db/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CreateThreatIntelligenceDataConnector.json
func ExampleDataConnectorsClient_CreateOrUpdate_createsOrUpdatesAnThreatIntelligencePlatformDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", &armsecurityinsights.TIDataConnector{
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindThreatIntelligence),
		Properties: &armsecurityinsights.TIDataConnectorProperties{
			TenantID: to.Ptr("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
			DataTypes: &armsecurityinsights.TIDataConnectorDataTypes{
				Indicators: &armsecurityinsights.TIDataConnectorDataTypesIndicators{
					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
				},
			},
			TipLookbackPeriod: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T13:00:30.123Z"); return t }()),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientCreateOrUpdateResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.TIDataConnector{
	// 		Name: to.Ptr("73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/73e01a99-5cd7-4139-a149-9f2736ff2ab5"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindThreatIntelligence),
	// 		Properties: &armsecurityinsights.TIDataConnectorProperties{
	// 			TenantID: to.Ptr("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
	// 			DataTypes: &armsecurityinsights.TIDataConnectorDataTypes{
	// 				Indicators: &armsecurityinsights.TIDataConnectorDataTypesIndicators{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			TipLookbackPeriod: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T13:00:30.123Z"); return t}()),
	// 		},
	// 	},
	// 	                        }
}

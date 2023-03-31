package armsecurityinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e24bbf6a66cb0a19c072c6f15cee163acbd7acf7/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/GetMicrosoftThreatIntelligenceById.json
func ExampleDataConnectorsClient_Get_getAMicrosoftThreatIntelligenceDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().Get(ctx, "myRg", "myWorkspace", "c345bf40-8509-4ed2-b947-50cb773aaf04", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientGetResponse{
	// 	                            DataConnectorClassification: &armsecurityinsights.MSTIDataConnector{
	// 		Name: to.Ptr("c345bf40-8509-4ed2-b947-50cb773aaf04"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/c345bf40-8509-4ed2-b947-50cb773aaf04"),
	// 		Etag: to.Ptr("\"0300bf09-0000-0000-0000-5c37296e0000\""),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindMicrosoftThreatIntelligence),
	// 		Properties: &armsecurityinsights.MSTIDataConnectorProperties{
	// 			TenantID: to.Ptr("2070ecc9-b4d5-4ae4-adaa-936fa1954fa8"),
	// 			DataTypes: &armsecurityinsights.MSTIDataConnectorDataTypes{
	// 				BingSafetyPhishingURL: &armsecurityinsights.MSTIDataConnectorDataTypesBingSafetyPhishingURL{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 					LookbackPeriod: to.Ptr("example ??"),
	// 				},
	// 				MicrosoftEmergingThreatFeed: &armsecurityinsights.MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 					LookbackPeriod: to.Ptr("example"),
	// 				},
	// 			},
	// 		},
	// 	},
	// 	                        }
}

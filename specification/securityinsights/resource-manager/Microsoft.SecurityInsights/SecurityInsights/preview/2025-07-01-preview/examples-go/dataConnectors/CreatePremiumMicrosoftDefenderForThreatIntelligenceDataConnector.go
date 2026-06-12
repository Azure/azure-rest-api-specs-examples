package armsecurityinsights_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/dataConnectors/CreatePremiumMicrosoftDefenderForThreatIntelligenceDataConnector.json
func ExampleDataConnectorsClient_CreateOrUpdate_createsOrUpdatesAPremiumMicrosoftDefenderForThreatIntelligenceDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("b66e5c69-e2eb-422a-81c3-002de57059f3", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "8c569548-a86c-4fb4-8ae4-d1e35a6146f8", &armsecurityinsights.PremiumMicrosoftDefenderForThreatIntelligence{
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindPremiumMicrosoftDefenderForThreatIntelligence),
		Properties: &armsecurityinsights.PremiumMdtiDataConnectorProperties{
			DataTypes: &armsecurityinsights.PremiumMdtiDataConnectorDataTypes{
				Connector: &armsecurityinsights.PremiumMdtiDataConnectorDataTypesConnector{
					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
				},
			},
			LookbackPeriod: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000Z"); return t }()),
			TenantID:       to.Ptr("e4afb3c4-813b-4e68-b6de-e5360866e798"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientCreateOrUpdateResponse{
	// 	DataConnectorClassification: &armsecurityinsights.PremiumMicrosoftDefenderForThreatIntelligence{
	// 		Name: to.Ptr("3deede2e-c6d1-4ee6-afc8-e0190ac34200"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		Etag: to.Ptr("56003401-0000-0100-0000-67314b0b0000"),
	// 		ID: to.Ptr("/subscriptions/b66e5c69-e2eb-422a-81c3-002de57059f3/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/3deede2e-c6d1-4ee6-afc8-e0190ac34200"),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindPremiumMicrosoftDefenderForThreatIntelligence),
	// 		Properties: &armsecurityinsights.PremiumMdtiDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.PremiumMdtiDataConnectorDataTypes{
	// 				Connector: &armsecurityinsights.PremiumMdtiDataConnectorDataTypesConnector{
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			LookbackPeriod: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-01T00:00:00Z"); return t}()),
	// 			RequiredSKUsPresent: to.Ptr(true),
	// 			TenantID: to.Ptr("e4afb3c4-813b-4e68-b6de-e5360866e798"),
	// 		},
	// 	},
	// }
}

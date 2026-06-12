package armsecurityinsights_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: 2025-07-01-preview/dataConnectors/CreateMicrosoftThreatIntelligenceDataConnector.json
func ExampleDataConnectorsClient_CreateOrUpdate_createsOrUpdatesAMicrosoftThreatIntelligenceDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsecurityinsights.NewClientFactory("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectorsClient().CreateOrUpdate(ctx, "myRg", "myWorkspace", "c345bf40-8509-4ed2-b947-50cb773aaf04", &armsecurityinsights.MSTIDataConnector{
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindMicrosoftThreatIntelligence),
		Properties: &armsecurityinsights.MSTIDataConnectorProperties{
			DataTypes: &armsecurityinsights.MSTIDataConnectorDataTypes{
				MicrosoftEmergingThreatFeed: &armsecurityinsights.MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed{
					LookbackPeriod: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000Z"); return t }()),
					State:          to.Ptr(armsecurityinsights.DataTypeStateEnabled),
				},
			},
			TenantID: to.Ptr("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsecurityinsights.DataConnectorsClientCreateOrUpdateResponse{
	// 	DataConnectorClassification: &armsecurityinsights.MSTIDataConnector{
	// 		Name: to.Ptr("c345bf40-8509-4ed2-b947-50cb773aaf04"),
	// 		Type: to.Ptr("Microsoft.SecurityInsights/dataConnectors"),
	// 		Etag: to.Ptr("d12423f6-a60b-4ca5-88c0-feb1a182d0f0"),
	// 		ID: to.Ptr("/subscriptions/d0cfe6b2-9ac0-4464-9919-dccaee2e48c0/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/dataConnectors/c345bf40-8509-4ed2-b947-50cb773aaf04"),
	// 		Kind: to.Ptr(armsecurityinsights.DataConnectorKindMicrosoftThreatIntelligence),
	// 		Properties: &armsecurityinsights.MSTIDataConnectorProperties{
	// 			DataTypes: &armsecurityinsights.MSTIDataConnectorDataTypes{
	// 				MicrosoftEmergingThreatFeed: &armsecurityinsights.MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed{
	// 					LookbackPeriod: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1970-01-01T00:00:00.000Z"); return t}()),
	// 					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
	// 				},
	// 			},
	// 			TenantID: to.Ptr("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
	// 		},
	// 	},
	// }
}

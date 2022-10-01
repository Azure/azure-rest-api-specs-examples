package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CreateThreatIntelligenceDataConnector.json
func ExampleDataConnectorsClient_CreateOrUpdate_createsOrUpdatesAnThreatIntelligencePlatformDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewDataConnectorsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", &armsecurityinsights.TIDataConnector{
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
	// TODO: use response item
	_ = res
}

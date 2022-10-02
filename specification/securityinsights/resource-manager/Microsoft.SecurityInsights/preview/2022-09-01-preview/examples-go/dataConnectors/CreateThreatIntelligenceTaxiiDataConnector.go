package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/dataConnectors/CreateThreatIntelligenceTaxiiDataConnector.json
func ExampleDataConnectorsClient_CreateOrUpdate_createsOrUpdatesAThreatIntelligenceTaxiiDataConnector() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewDataConnectorsClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", &armsecurityinsights.TiTaxiiDataConnector{
		Etag: to.Ptr("d12423f6-a60b-4ca5-88c0-feb1a182d0f0"),
		Kind: to.Ptr(armsecurityinsights.DataConnectorKindThreatIntelligenceTaxii),
		Properties: &armsecurityinsights.TiTaxiiDataConnectorProperties{
			TenantID:     to.Ptr("06b3ccb8-1384-4bcc-aec7-852f6d57161b"),
			CollectionID: to.Ptr("135"),
			DataTypes: &armsecurityinsights.TiTaxiiDataConnectorDataTypes{
				TaxiiClient: &armsecurityinsights.TiTaxiiDataConnectorDataTypesTaxiiClient{
					State: to.Ptr(armsecurityinsights.DataTypeStateEnabled),
				},
			},
			FriendlyName:        to.Ptr("testTaxii"),
			Password:            to.Ptr("--"),
			PollingFrequency:    to.Ptr(armsecurityinsights.PollingFrequencyOnceADay),
			TaxiiLookbackPeriod: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T13:00:30.123Z"); return t }()),
			TaxiiServer:         to.Ptr("https://limo.anomali.com/api/v1/taxii2/feeds"),
			UserName:            to.Ptr("--"),
			WorkspaceID:         to.Ptr("dd124572-4962-4495-9bd2-9dade12314b4"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

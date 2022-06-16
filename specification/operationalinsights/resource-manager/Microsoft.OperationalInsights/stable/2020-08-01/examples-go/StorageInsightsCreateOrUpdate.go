package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/StorageInsightsCreateOrUpdate.json
func ExampleStorageInsightConfigsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armoperationalinsights.NewStorageInsightConfigsClient("00000000-0000-0000-0000-00000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"OIAutoRest5123",
		"aztest5048",
		"AzTestSI1110",
		armoperationalinsights.StorageInsight{
			Properties: &armoperationalinsights.StorageInsightProperties{
				Containers: []*string{
					to.Ptr("wad-iis-logfiles")},
				StorageAccount: &armoperationalinsights.StorageAccount{
					ID:  to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/OIAutoRest6987/providers/microsoft.storage/storageaccounts/AzTestFakeSA9945"),
					Key: to.Ptr("1234"),
				},
				Tables: []*string{
					to.Ptr("WADWindowsEventLogsTable"),
					to.Ptr("LinuxSyslogVer2v0")},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

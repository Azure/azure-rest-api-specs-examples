package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f0a5127d4e8b1ea6007b0bf9570904d9df860b97/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/StorageInsightsCreateOrUpdate.json
func ExampleStorageInsightConfigsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStorageInsightConfigsClient().CreateOrUpdate(ctx, "OIAutoRest5123", "aztest5048", "AzTestSI1110", armoperationalinsights.StorageInsight{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StorageInsight = armoperationalinsights.StorageInsight{
	// 	Name: to.Ptr("AzTestSI1110"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/workspaces/storageinsightconfigs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/oiautorest6987/providers/microsoft.operationalinsights/workspaces/aztest5048/storageinsightconfigs/AzTestSI1110"),
	// 	Properties: &armoperationalinsights.StorageInsightProperties{
	// 		Containers: []*string{
	// 			to.Ptr("wad-iis-logfiles")},
	// 			Status: &armoperationalinsights.StorageInsightStatus{
	// 				State: to.Ptr(armoperationalinsights.StorageInsightStateOK),
	// 			},
	// 			StorageAccount: &armoperationalinsights.StorageAccount{
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/OIAutoRest6987/providers/microsoft.storage/storageaccounts/AzTestFakeSA9945"),
	// 				Key: to.Ptr("Storage Key"),
	// 			},
	// 			Tables: []*string{
	// 				to.Ptr("WADWindowsEventLogsTable"),
	// 				to.Ptr("LinuxSyslogVer2v0")},
	// 			},
	// 		}
}

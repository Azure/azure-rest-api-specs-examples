package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v3"
)

// Generated from example definition: 2025-07-01/StorageInsightsListByWorkspace.json
func ExampleStorageInsightConfigsClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageInsightConfigsClient().NewListByWorkspacePager("OIAutoRest5123", "aztest5048", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armoperationalinsights.StorageInsightConfigsClientListByWorkspaceResponse{
		// 	StorageInsightListResult: armoperationalinsights.StorageInsightListResult{
		// 		Value: []*armoperationalinsights.StorageInsight{
		// 			{
		// 				Name: to.Ptr("AzTestSI1110"),
		// 				Type: to.Ptr("Microsoft.OperationalInsights/workspaces/storageinsightconfigs"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/oiautorest6987/providers/microsoft.operationalinsights/workspaces/aztest5048/storageinsightconfigs/AzTestSI1110"),
		// 				Properties: &armoperationalinsights.StorageInsightProperties{
		// 					Containers: []*string{
		// 						to.Ptr("wad-iis-logfiles"),
		// 					},
		// 					Status: &armoperationalinsights.StorageInsightStatus{
		// 						State: to.Ptr(armoperationalinsights.StorageInsightStateOK),
		// 					},
		// 					StorageAccount: &armoperationalinsights.StorageAccount{
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/OIAutoRest6987/providers/microsoft.storage/storageaccounts/AzTestFakeSA9945"),
		// 						Key: to.Ptr("1234"),
		// 					},
		// 					Tables: []*string{
		// 						to.Ptr("WADWindowsEventLogsTable"),
		// 						to.Ptr("LinuxSyslogVer2v0"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

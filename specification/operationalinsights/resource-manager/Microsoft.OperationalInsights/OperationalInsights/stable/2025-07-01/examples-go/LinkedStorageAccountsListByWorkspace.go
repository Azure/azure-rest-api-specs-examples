package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f0a5127d4e8b1ea6007b0bf9570904d9df860b97/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-07-01/examples/LinkedStorageAccountsListByWorkspace.json
func ExampleLinkedStorageAccountsClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLinkedStorageAccountsClient().NewListByWorkspacePager("mms-eus", "testLinkStorageAccountsWS", nil)
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
		// page.LinkedStorageAccountsListResult = armoperationalinsights.LinkedStorageAccountsListResult{
		// 	Value: []*armoperationalinsights.LinkedStorageAccountsResource{
		// 		{
		// 			Name: to.Ptr("CustomLogs"),
		// 			Type: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedStorageAccounts"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/mms-eus/providers/microsoft.operationalinsights/workspaces/testLinkStorageAccountsWS/linkedStorageAccounts/CustomLogs"),
		// 			Properties: &armoperationalinsights.LinkedStorageAccountsProperties{
		// 				DataSourceType: to.Ptr(armoperationalinsights.DataSourceTypeCustomLogs),
		// 				StorageAccountIDs: []*string{
		// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageA"),
		// 					to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageB")},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("AzureWatson"),
		// 				Type: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedStorageAccounts"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/mms-eus/providers/microsoft.operationalinsights/workspaces/testLinkStorageAccountsWS/linkedStorageAccounts/AzureWatson"),
		// 				Properties: &armoperationalinsights.LinkedStorageAccountsProperties{
		// 					DataSourceType: to.Ptr(armoperationalinsights.DataSourceTypeAzureWatson),
		// 					StorageAccountIDs: []*string{
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageA"),
		// 						to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageC")},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("Query"),
		// 					Type: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedStorageAccounts"),
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/mms-eus/providers/microsoft.operationalinsights/workspaces/testLinkStorageAccountsWS/linkedStorageAccounts/Query"),
		// 					Properties: &armoperationalinsights.LinkedStorageAccountsProperties{
		// 						DataSourceType: to.Ptr(armoperationalinsights.DataSourceTypeQuery),
		// 						StorageAccountIDs: []*string{
		// 							to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageA"),
		// 							to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageC")},
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("Alerts"),
		// 						Type: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedStorageAccounts"),
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/mms-eus/providers/microsoft.operationalinsights/workspaces/testLinkStorageAccountsWS/linkedStorageAccounts/Alerts"),
		// 						Properties: &armoperationalinsights.LinkedStorageAccountsProperties{
		// 							DataSourceType: to.Ptr(armoperationalinsights.DataSourceTypeAlerts),
		// 							StorageAccountIDs: []*string{
		// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageA"),
		// 								to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageC")},
		// 							},
		// 					}},
		// 				}
	}
}

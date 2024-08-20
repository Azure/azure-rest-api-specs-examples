package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4ce13e8353a25125a41bc01705c0a7794dac32a7/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/LinkedStorageAccountsGet.json
func ExampleLinkedStorageAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoperationalinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLinkedStorageAccountsClient().Get(ctx, "mms-eus", "testLinkStorageAccountsWS", armoperationalinsights.DataSourceTypeCustomLogs, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LinkedStorageAccountsResource = armoperationalinsights.LinkedStorageAccountsResource{
	// 	Name: to.Ptr("CustomLogs"),
	// 	Type: to.Ptr("Microsoft.OperationalInsights/workspaces/linkedStorageAccounts"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000005/resourcegroups/mms-eus/providers/microsoft.operationalinsights/workspaces/testLinkStorageAccountsWS/linkedStorageAccounts/CustomLogs"),
	// 	Properties: &armoperationalinsights.LinkedStorageAccountsProperties{
	// 		DataSourceType: to.Ptr(armoperationalinsights.DataSourceTypeCustomLogs),
	// 		StorageAccountIDs: []*string{
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageA"),
	// 			to.Ptr("/subscriptions/00000000-0000-0000-0000-00000000000/resourceGroups/mms-eus/providers/Microsoft.Storage/storageAccounts/testStorageB")},
	// 		},
	// 	}
}

package armappcomplianceautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d88c94b22a8efdd47c0cadfe6d8d489107db2b23/specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/ListInUseStorageAccountsWithoutSubscriptions.json
func ExampleProviderActionsClient_ListInUseStorageAccounts_listInUseStorageAccountsWithoutSubscriptions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcomplianceautomation.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProviderActionsClient().ListInUseStorageAccounts(ctx, armappcomplianceautomation.ListInUseStorageAccountsRequest{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListInUseStorageAccountsResponse = armappcomplianceautomation.ListInUseStorageAccountsResponse{
	// 	StorageAccountList: []*armappcomplianceautomation.StorageInfo{
	// 		{
	// 			AccountName: to.Ptr("SA_name1"),
	// 			Location: to.Ptr("WEST US"),
	// 			ResourceGroup: to.Ptr("tetsRG"),
	// 			SubscriptionID: to.Ptr("0000000-0000-0000-0000-000000000001"),
	// 		},
	// 		{
	// 			AccountName: to.Ptr("SA_name2"),
	// 			Location: to.Ptr("WEST US"),
	// 			ResourceGroup: to.Ptr("tetsRG"),
	// 			SubscriptionID: to.Ptr("0000000-0000-0000-0000-000000000001"),
	// 	}},
	// }
}

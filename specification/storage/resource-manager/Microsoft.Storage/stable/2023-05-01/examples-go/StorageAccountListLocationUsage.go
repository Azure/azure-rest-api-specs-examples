package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountListLocationUsage.json
func ExampleUsagesClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewUsagesClient().NewListByLocationPager("eastus2(stage)", nil)
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
		// page.UsageListResult = armstorage.UsageListResult{
		// 	Value: []*armstorage.Usage{
		// 		{
		// 			Name: &armstorage.UsageName{
		// 				LocalizedValue: to.Ptr("Storage Accounts"),
		// 				Value: to.Ptr("StorageAccounts"),
		// 			},
		// 			CurrentValue: to.Ptr[int32](55),
		// 			Limit: to.Ptr[int32](250),
		// 			Unit: to.Ptr(armstorage.UsageUnitCount),
		// 	}},
		// }
	}
}

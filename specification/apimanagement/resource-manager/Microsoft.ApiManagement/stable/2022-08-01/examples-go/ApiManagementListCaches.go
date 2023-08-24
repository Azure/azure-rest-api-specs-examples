package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListCaches.json
func ExampleCacheClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCacheClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.CacheClientListByServiceOptions{Top: nil,
		Skip: nil,
	})
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
		// page.CacheCollection = armapimanagement.CacheCollection{
		// 	Count: to.Ptr[int64](1),
		// 	Value: []*armapimanagement.CacheContract{
		// 		{
		// 			Name: to.Ptr("c1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/caches"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/caches/c1"),
		// 			Properties: &armapimanagement.CacheContractProperties{
		// 				Description: to.Ptr("Redis cache instances in West India"),
		// 				ConnectionString: to.Ptr("{{5f7fbca77a891a2200f3db38}}"),
		// 				ResourceID: to.Ptr("https://management.azure.com/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redis/apimservice1"),
		// 				UseFromLocation: to.Ptr("default"),
		// 			},
		// 	}},
		// }
	}
}

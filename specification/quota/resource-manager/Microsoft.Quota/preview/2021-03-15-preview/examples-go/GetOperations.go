package armquota_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/quota/resource-manager/Microsoft.Quota/preview/2021-03-15-preview/examples/GetOperations.json
func ExampleOperationClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armquota.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationClient().NewListPager(nil)
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
		// page.OperationList = armquota.OperationList{
		// 	Value: []*armquota.OperationResponse{
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/usages/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get the current usages of the specified resource and location"),
		// 				Operation: to.Ptr("Get resource usages"),
		// 				Provider: to.Ptr("Microsoft Quota"),
		// 				Resource: to.Ptr("Resource current usages"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/quotas/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get the current service limit or quota of the specified resource and location"),
		// 				Operation: to.Ptr("Get resource quota limit"),
		// 				Provider: to.Ptr("Microsoft Quota"),
		// 				Resource: to.Ptr("Resource quota limit"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/quotas/write"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Create service limit or quota for the specified resource and location"),
		// 				Operation: to.Ptr("Create resource quota limit"),
		// 				Provider: to.Ptr("Microsoft Quota"),
		// 				Resource: to.Ptr("Resource quota limit"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.Quota/quotaRequests/read"),
		// 			Display: &armquota.OperationDisplay{
		// 				Description: to.Ptr("Get any quota request for the specified resource and location"),
		// 				Operation: to.Ptr("Get quota request"),
		// 				Provider: to.Ptr("Microsoft Quota"),
		// 				Resource: to.Ptr("Resource quota request"),
		// 			},
		// 	}},
		// }
	}
}

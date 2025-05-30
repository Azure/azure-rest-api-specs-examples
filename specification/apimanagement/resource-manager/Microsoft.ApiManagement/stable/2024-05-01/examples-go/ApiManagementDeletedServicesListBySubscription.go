package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementDeletedServicesListBySubscription.json
func ExampleDeletedServicesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeletedServicesClient().NewListBySubscriptionPager(nil)
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
		// page.DeletedServicesCollection = armapimanagement.DeletedServicesCollection{
		// 	Value: []*armapimanagement.DeletedServiceContract{
		// 		{
		// 			Name: to.Ptr("apimService3"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/deletedservices"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.ApiManagement/locations/westus/deletedservices/apimService3"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armapimanagement.DeletedServiceContractProperties{
		// 				DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-27T15:33:55.542Z"); return t}()),
		// 				ScheduledPurgeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-27T15:33:55.542Z"); return t}()),
		// 				ServiceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService3"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("apimService"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/deletedservices"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.ApiManagement/locations/westus2/deletedservices/apimService"),
		// 			Location: to.Ptr("West US 2"),
		// 			Properties: &armapimanagement.DeletedServiceContractProperties{
		// 				DeletionDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-27T15:33:55.542Z"); return t}()),
		// 				ScheduledPurgeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-27T15:33:55.542Z"); return t}()),
		// 				ServiceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService"),
		// 			},
		// 	}},
		// }
	}
}

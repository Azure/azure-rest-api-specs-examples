package armlargeinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/largeinstance/armlargeinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/examples/AzureLargeStorageInstance_ListBySubscription.json
func ExampleAzureLargeStorageInstanceClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlargeinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAzureLargeStorageInstanceClient().NewListBySubscriptionPager(nil)
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
		// page.AzureLargeStorageInstanceListResult = armlargeinstance.AzureLargeStorageInstanceListResult{
		// 	Value: []*armlargeinstance.AzureLargeStorageInstance{
		// 		{
		// 			Name: to.Ptr("myAzureLargeInstance1"),
		// 			Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeStorageInstances"),
		// 			ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeStorageInstances/myAzureLargeInstance1"),
		// 			SystemData: &armlargeinstance.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-20T23:10:22.682Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@microsoft.com"),
		// 				CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("westus2"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armlargeinstance.AzureLargeStorageInstanceProperties{
		// 				AzureLargeStorageInstanceUniqueIdentifier: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24e9"),
		// 				StorageProperties: &armlargeinstance.StorageProperties{
		// 					Generation: to.Ptr("Gen4"),
		// 					HardwareType: to.Ptr(armlargeinstance.AzureLargeInstanceHardwareTypeNamesEnum("NetApp")),
		// 					OfferingType: to.Ptr("EPIC"),
		// 					ProvisioningState: to.Ptr(armlargeinstance.ProvisioningStateSucceeded),
		// 					StorageBillingProperties: &armlargeinstance.StorageBillingProperties{
		// 						BillingMode: to.Ptr("PAYG"),
		// 						SKU: to.Ptr(""),
		// 					},
		// 					StorageType: to.Ptr("FC"),
		// 					WorkloadType: to.Ptr("ODB"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myAzureLargeInstance2"),
		// 			Type: to.Ptr("Microsoft.AzureLargeInstance/AzureLargeStorageInstances"),
		// 			ID: to.Ptr("/subscriptions/f0f4887f-d13c-4943-a8ba-d7da28d2a3fd/resourceGroups/myResourceGroup/providers/Microsoft.AzureLargeInstance/AzureLargeStorageInstances/myAzureLargeInstance2"),
		// 			SystemData: &armlargeinstance.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-21T08:01:22.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@microsoft.com"),
		// 				CreatedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-13T08:01:22.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armlargeinstance.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("westus2"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Properties: &armlargeinstance.AzureLargeStorageInstanceProperties{
		// 				AzureLargeStorageInstanceUniqueIdentifier: to.Ptr("23415635-4d7e-41dc-9598-8194f22c24f7"),
		// 				StorageProperties: &armlargeinstance.StorageProperties{
		// 					Generation: to.Ptr("Gen4"),
		// 					HardwareType: to.Ptr(armlargeinstance.AzureLargeInstanceHardwareTypeNamesEnum("NetApp")),
		// 					OfferingType: to.Ptr("EPIC"),
		// 					ProvisioningState: to.Ptr(armlargeinstance.ProvisioningStateSucceeded),
		// 					StorageBillingProperties: &armlargeinstance.StorageBillingProperties{
		// 						BillingMode: to.Ptr("RI"),
		// 						SKU: to.Ptr(""),
		// 					},
		// 					StorageType: to.Ptr("NFS"),
		// 					WorkloadType: to.Ptr("Cogito"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

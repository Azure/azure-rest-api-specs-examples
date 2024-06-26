package armstoragepool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPoolZones_List.json
func ExampleDiskPoolZonesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragepool.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiskPoolZonesClient().NewListPager("eastus", nil)
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
		// page.DiskPoolZoneListResult = armstoragepool.DiskPoolZoneListResult{
		// 	Value: []*armstoragepool.DiskPoolZoneInfo{
		// 		{
		// 			AdditionalCapabilities: []*string{
		// 			},
		// 			AvailabilityZones: []*string{
		// 				to.Ptr("1"),
		// 				to.Ptr("2"),
		// 				to.Ptr("3")},
		// 				SKU: &armstoragepool.SKU{
		// 					Name: to.Ptr("Basic"),
		// 					Tier: to.Ptr("Basic_V1"),
		// 				},
		// 			},
		// 			{
		// 				AdditionalCapabilities: []*string{
		// 				},
		// 				AvailabilityZones: []*string{
		// 					to.Ptr("1"),
		// 					to.Ptr("2")},
		// 					SKU: &armstoragepool.SKU{
		// 						Name: to.Ptr("Standard"),
		// 						Tier: to.Ptr("Standard_V1"),
		// 					},
		// 			}},
		// 		}
	}
}

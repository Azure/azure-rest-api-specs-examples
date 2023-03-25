package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-09-01/examples/Subvolumes_List.json
func ExampleSubvolumesClient_NewListByVolumePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubvolumesClient().NewListByVolumePager("myRG", "account1", "pool1", "volume1", nil)
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
		// page.SubvolumesList = armnetapp.SubvolumesList{
		// 	Value: []*armnetapp.SubvolumeInfo{
		// 		{
		// 			Name: to.Ptr("account1/pool1/volume1/subvolume1"),
		// 			Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools/volumes/subvolumes"),
		// 			ID: to.Ptr("/subscriptions/D633CC2E-722B-4AE1-B636-BBD9E4C60ED9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1/subvolumes/subvolume1"),
		// 			Properties: &armnetapp.SubvolumeProperties{
		// 				Path: to.Ptr("/pathToSubvol"),
		// 			},
		// 	}},
		// }
	}
}

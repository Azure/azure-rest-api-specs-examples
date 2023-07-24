package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/VolumeGroups_List.json
func ExampleVolumeGroupsClient_NewListByNetAppAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVolumeGroupsClient().NewListByNetAppAccountPager("myRG", "account1", nil)
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
		// page.VolumeGroupList = armnetapp.VolumeGroupList{
		// 	Value: []*armnetapp.VolumeGroup{
		// 		{
		// 			Name: to.Ptr("group1"),
		// 			Type: to.Ptr("Microsoft.NetApp/netAppAccounts/volumeGroups"),
		// 			ID: to.Ptr("/subscriptions/d633cc2e-722b-4ae1-b636-bbd9e4c60ed9/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/volumeGroups/group1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetapp.VolumeGroupListProperties{
		// 				GroupMetaData: &armnetapp.VolumeGroupMetaData{
		// 					ApplicationIdentifier: to.Ptr("DEV"),
		// 					ApplicationType: to.Ptr(armnetapp.ApplicationTypeSAPHANA),
		// 					GroupDescription: to.Ptr("Volume group"),
		// 					VolumesCount: to.Ptr[int64](5),
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}

package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/639ecfad68419328658bd4cfe7094af4ce472be2/specification/netapp/resource-manager/Microsoft.NetApp/preview/2023-05-01-preview/examples/NetworkSiblingSet_Query.json
func ExampleResourceClient_QueryNetworkSiblingSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewResourceClient().QueryNetworkSiblingSet(ctx, "eastus", armnetapp.QueryNetworkSiblingSetRequest{
		NetworkSiblingSetID: to.Ptr("9760acf5-4638-11e7-9bdb-020073ca3333"),
		SubnetID:            to.Ptr("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testVnet/subnets/testSubnet"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSiblingSet = armnetapp.NetworkSiblingSet{
	// 	NetworkFeatures: to.Ptr(armnetapp.NetworkFeaturesStandard),
	// 	NetworkSiblingSetID: to.Ptr("9760acf5-4638-11e7-9bdb-020073ca3333"),
	// 	NetworkSiblingSetStateID: to.Ptr("12345_44420.8001578125"),
	// 	NicInfoList: []*armnetapp.NicInfo{
	// 		{
	// 			IPAddress: to.Ptr("1.2.3.4"),
	// 			VolumeResourceIDs: []*string{
	// 				to.Ptr("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume10"),
	// 				to.Ptr("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume11")},
	// 			},
	// 			{
	// 				IPAddress: to.Ptr("1.2.3.5"),
	// 				VolumeResourceIDs: []*string{
	// 					to.Ptr("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account2/capacityPools/pool2/volumes/volume20"),
	// 					to.Ptr("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account2/capacityPools/pool2/volumes/volume21")},
	// 				},
	// 				{
	// 					IPAddress: to.Ptr("1.2.3.9"),
	// 					VolumeResourceIDs: []*string{
	// 					},
	// 			}},
	// 			ProvisioningState: to.Ptr(armnetapp.NetworkSiblingSetProvisioningStateSucceeded),
	// 			SubnetID: to.Ptr("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testVnet/subnets/testSubnet"),
	// 		}
}

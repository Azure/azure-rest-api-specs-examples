package armnetapp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ba99beec358a40ee08dae7f12f6a989aad6ce6d1/specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-06-01/examples/Pools_List.json
func ExamplePoolsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetapp.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPoolsClient().NewListPager("myRG", "account1", nil)
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
		// page.CapacityPoolList = armnetapp.CapacityPoolList{
		// 	Value: []*armnetapp.CapacityPool{
		// 		{
		// 			Name: to.Ptr("account1/pool1"),
		// 			Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetapp.PoolProperties{
		// 				PoolID: to.Ptr("9760acf5-4638-11e7-9bdb-020073ca7778"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				QosType: to.Ptr(armnetapp.QosTypeAuto),
		// 				ServiceLevel: to.Ptr(armnetapp.ServiceLevelPremium),
		// 				Size: to.Ptr[int64](4398046511104),
		// 				TotalThroughputMibps: to.Ptr[float32](281.474),
		// 				UtilizedThroughputMibps: to.Ptr[float32](100.47),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("account1/customPool1"),
		// 			Type: to.Ptr("Microsoft.NetApp/netAppAccounts/capacityPools"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/customPool1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armnetapp.PoolProperties{
		// 				CustomThroughputMibps: to.Ptr[float32](128),
		// 				PoolID: to.Ptr("9760acf5-4638-11e7-9bdb-020073ca7778"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				QosType: to.Ptr(armnetapp.QosTypeManual),
		// 				ServiceLevel: to.Ptr(armnetapp.ServiceLevelFlexible),
		// 				Size: to.Ptr[int64](4398046511104),
		// 				TotalThroughputMibps: to.Ptr[float32](128),
		// 				UtilizedThroughputMibps: to.Ptr[float32](100.47),
		// 			},
		// 	}},
		// }
	}
}

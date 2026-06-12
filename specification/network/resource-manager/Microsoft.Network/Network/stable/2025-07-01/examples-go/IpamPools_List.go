package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/IpamPools_List.json
func ExampleIpamPoolsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("11111111-1111-1111-1111-111111111111", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIpamPoolsClient().NewListPager("rg1", "TestNetworkManager", nil)
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
		// page = armnetwork.IpamPoolsClientListResponse{
		// 	IpamPoolList: armnetwork.IpamPoolList{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/TestNetworkManager/ipamPools?api-version=2025-07-01&$skipToken=10"),
		// 		Value: []*armnetwork.IpamPool{
		// 			{
		// 				Name: to.Ptr("TestPool"),
		// 				Type: to.Ptr("Microsoft.Network/networkManagers/ipamPools"),
		// 				Etag: to.Ptr("\"00000000-0000-0000-0000-000000000000\""),
		// 				ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/TestNetworkManager/ipamPools/TestPool"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armnetwork.IpamPoolProperties{
		// 					Description: to.Ptr("Test description."),
		// 					AddressPrefixes: []*string{
		// 						to.Ptr("10.0.0.0/24"),
		// 					},
		// 					IPAddressType: []*armnetwork.IPType{
		// 						to.Ptr(armnetwork.IPTypeIPv4),
		// 					},
		// 					ParentPoolName: to.Ptr(""),
		// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				},
		// 				SystemData: &armnetwork.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
		// 					CreatedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 					CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-11T18:52:27Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("b69a9388-9488-4534-b470-7ec6d41beef5"),
		// 					LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

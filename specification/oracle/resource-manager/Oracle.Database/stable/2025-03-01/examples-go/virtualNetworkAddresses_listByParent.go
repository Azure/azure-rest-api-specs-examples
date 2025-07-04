package armoracledatabase_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
)

// Generated from example definition: 2025-03-01/virtualNetworkAddresses_listByParent.json
func ExampleVirtualNetworkAddressesClient_NewListByCloudVMClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armoracledatabase.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkAddressesClient().NewListByCloudVMClusterPager("rg000", "cluster1", nil)
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
		// page = armoracledatabase.VirtualNetworkAddressesClientListByCloudVMClusterResponse{
		// 	VirtualNetworkAddressListResult: armoracledatabase.VirtualNetworkAddressListResult{
		// 		Value: []*armoracledatabase.VirtualNetworkAddress{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg000/providers/Oracle.Database/cloudVmClusters/cluster1/virtualNetworkAddresses/hostname1"),
		// 				Type: to.Ptr("Oracle.Database/cloudVmClusters/virtualNetworkAddresses"),
		// 				Properties: &armoracledatabase.VirtualNetworkAddressProperties{
		// 					IPAddress: to.Ptr("192.168.0.1"),
		// 					VMOcid: to.Ptr("ocid1..aaaa"),
		// 					Ocid: to.Ptr("ocid1..aaaaaa"),
		// 					Domain: to.Ptr("domain1"),
		// 					LifecycleDetails: to.Ptr("success"),
		// 					ProvisioningState: to.Ptr(armoracledatabase.AzureResourceProvisioningStateSucceeded),
		// 					LifecycleState: to.Ptr(armoracledatabase.VirtualNetworkAddressLifecycleStateAvailable),
		// 					TimeAssigned: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-22T00:27:02.119Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

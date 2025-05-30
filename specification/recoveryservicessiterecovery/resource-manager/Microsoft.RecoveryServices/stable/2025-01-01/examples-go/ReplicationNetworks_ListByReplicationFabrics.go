package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fc983fb08e5fd8a7a821eb6491f5338ce52a9cf/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples/ReplicationNetworks_ListByReplicationFabrics.json
func ExampleReplicationNetworksClient_NewListByReplicationFabricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationNetworksClient().NewListByReplicationFabricsPager("srcBvte2a14C27", "srce2avaultbvtaC27", "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac", nil)
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
		// page.NetworkCollection = armrecoveryservicessiterecovery.NetworkCollection{
		// 	Value: []*armrecoveryservicessiterecovery.Network{
		// 		{
		// 			Name: to.Ptr("93ce99d7-1219-4914-aa61-73fe5023988e"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationNetworks"),
		// 			ID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/srcBvte2a14C27/providers/Microsoft.RecoveryServices/vaults/srce2avaultbvtaC27/replicationFabrics/b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac/replicationNetworks/93ce99d7-1219-4914-aa61-73fe5023988e"),
		// 			Properties: &armrecoveryservicessiterecovery.NetworkProperties{
		// 				FabricType: to.Ptr("VMM"),
		// 				FriendlyName: to.Ptr("VSwitch_VLan"),
		// 				NetworkType: to.Ptr("NoIsolation"),
		// 				Subnets: []*armrecoveryservicessiterecovery.Subnet{
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("b83bf8fd-f304-48d7-82c9-5d74e6215c1b"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationNetworks"),
		// 			ID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/srcBvte2a14C27/providers/Microsoft.RecoveryServices/vaults/srce2avaultbvtaC27/replicationFabrics/b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac/replicationNetworks/b83bf8fd-f304-48d7-82c9-5d74e6215c1b"),
		// 			Properties: &armrecoveryservicessiterecovery.NetworkProperties{
		// 				FabricType: to.Ptr("VMM"),
		// 				FriendlyName: to.Ptr("VSwitch_NoIso"),
		// 				NetworkType: to.Ptr("NoIsolation"),
		// 				Subnets: []*armrecoveryservicessiterecovery.Subnet{
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("e2267b5c-2650-49bd-ab3f-d66aae694c06"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationNetworks"),
		// 			ID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/srcBvte2a14C27/providers/Microsoft.RecoveryServices/vaults/srce2avaultbvtaC27/replicationFabrics/b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac/replicationNetworks/e2267b5c-2650-49bd-ab3f-d66aae694c06"),
		// 			Properties: &armrecoveryservicessiterecovery.NetworkProperties{
		// 				FabricType: to.Ptr("VMM"),
		// 				FriendlyName: to.Ptr("corp"),
		// 				NetworkType: to.Ptr("NoIsolation"),
		// 				Subnets: []*armrecoveryservicessiterecovery.Subnet{
		// 				},
		// 			},
		// 	}},
		// }
	}
}

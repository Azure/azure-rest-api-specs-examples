package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationNetworkMappings_ListByReplicationNetworks.json
func ExampleReplicationNetworkMappingsClient_NewListByReplicationNetworksPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewReplicationNetworkMappingsClient().NewListByReplicationNetworksPager("srce2avaultbvtaC27", "srcBvte2a14C27", "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac", "e2267b5c-2650-49bd-ab3f-d66aae694c06", nil)
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
		// page.NetworkMappingCollection = armrecoveryservicessiterecovery.NetworkMappingCollection{
		// 	Value: []*armrecoveryservicessiterecovery.NetworkMapping{
		// 		{
		// 			Name: to.Ptr("corpe2amap"),
		// 			Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationNetworks/replicationNetworkMappings"),
		// 			ID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/srcBvte2a14C27/providers/Microsoft.RecoveryServices/vaults/srce2avaultbvtaC27/replicationFabrics/b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac/replicationNetworks/e2267b5c-2650-49bd-ab3f-d66aae694c06/replicationNetworkMappings/corpe2amap"),
		// 			Properties: &armrecoveryservicessiterecovery.NetworkMappingProperties{
		// 				FabricSpecificSettings: &armrecoveryservicessiterecovery.VmmToAzureNetworkMappingSettings{
		// 					InstanceType: to.Ptr("VmmToAzure"),
		// 				},
		// 				PrimaryFabricFriendlyName: to.Ptr("CP-B3L30108-01.ntdev.corp.microsoft.com"),
		// 				PrimaryNetworkFriendlyName: to.Ptr("corp"),
		// 				PrimaryNetworkID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/srcBvte2a14C27/providers/Microsoft.RecoveryServices/vaults/srce2avaultbvtaC27/replicationFabrics/d49858f157601230a6ac5862fbbc6e63bf38d23ecd96cf953767945d457fe9d5/replicationNetworks/e2267b5c-2650-49bd-ab3f-d66aae694c06"),
		// 				RecoveryFabricArmID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/srcBvte2a14C27/providers/Microsoft.RecoveryServices/vaults/srce2avaultbvtaC27/replicationFabrics/d49858f157601230a6ac5862fbbc6e63bf38d23ecd96cf953767945d457fe9d5"),
		// 				RecoveryFabricFriendlyName: to.Ptr("Microsoft Azure"),
		// 				RecoveryNetworkFriendlyName: to.Ptr("vnetavrai"),
		// 				RecoveryNetworkID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai"),
		// 				State: to.Ptr("Paired"),
		// 			},
		// 	}},
		// }
	}
}

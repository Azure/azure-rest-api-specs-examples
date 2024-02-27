package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationNetworkMappings_Get.json
func ExampleReplicationNetworkMappingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationNetworkMappingsClient().Get(ctx, "srce2avaultbvtaC27", "srcBvte2a14C27", "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac", "e2267b5c-2650-49bd-ab3f-d66aae694c06", "corpe2amap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkMapping = armrecoveryservicessiterecovery.NetworkMapping{
	// 	Name: to.Ptr("corpe2amap"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationNetworks/replicationNetworkMappings"),
	// 	ID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/srcBvte2a14C27/providers/Microsoft.RecoveryServices/vaults/srce2avaultbvtaC27/replicationFabrics/b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac/replicationNetworks/e2267b5c-2650-49bd-ab3f-d66aae694c06/replicationNetworkMappings/corpe2amap"),
	// 	Properties: &armrecoveryservicessiterecovery.NetworkMappingProperties{
	// 		FabricSpecificSettings: &armrecoveryservicessiterecovery.VmmToAzureNetworkMappingSettings{
	// 			InstanceType: to.Ptr("VmmToAzure"),
	// 		},
	// 		PrimaryFabricFriendlyName: to.Ptr("CP-B3L30108-01.ntdev.corp.microsoft.com"),
	// 		PrimaryNetworkFriendlyName: to.Ptr("corp"),
	// 		PrimaryNetworkID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/srcBvte2a14C27/providers/Microsoft.RecoveryServices/vaults/srce2avaultbvtaC27/replicationFabrics/d49858f157601230a6ac5862fbbc6e63bf38d23ecd96cf953767945d457fe9d5/replicationNetworks/e2267b5c-2650-49bd-ab3f-d66aae694c06"),
	// 		RecoveryFabricArmID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/srcBvte2a14C27/providers/Microsoft.RecoveryServices/vaults/srce2avaultbvtaC27/replicationFabrics/d49858f157601230a6ac5862fbbc6e63bf38d23ecd96cf953767945d457fe9d5"),
	// 		RecoveryFabricFriendlyName: to.Ptr("Microsoft Azure"),
	// 		RecoveryNetworkFriendlyName: to.Ptr("vnetavrai"),
	// 		RecoveryNetworkID: to.Ptr("/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai"),
	// 		State: to.Ptr("Paired"),
	// 	},
	// }
}

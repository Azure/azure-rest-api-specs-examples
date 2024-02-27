package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationNetworks_Get.json
func ExampleReplicationNetworksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationNetworksClient().Get(ctx, "srce2avaultbvtaC27", "srcBvte2a14C27", "b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac", "93ce99d7-1219-4914-aa61-73fe5023988e", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Network = armrecoveryservicessiterecovery.Network{
	// 	Name: to.Ptr("93ce99d7-1219-4914-aa61-73fe5023988e"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics/replicationNetworks"),
	// 	ID: to.Ptr("/Subscriptions/9112a37f-0f3e-46ec-9c00-060c6edca071/resourceGroups/srcBvte2a14C27/providers/Microsoft.RecoveryServices/vaults/srce2avaultbvtaC27/replicationFabrics/b0cef6e9a4437b81803d0b55ada4f700ab66caae59c35d62723a1589c0cd13ac/replicationNetworks/93ce99d7-1219-4914-aa61-73fe5023988e"),
	// 	Properties: &armrecoveryservicessiterecovery.NetworkProperties{
	// 		FabricType: to.Ptr("VMM"),
	// 		FriendlyName: to.Ptr("VSwitch_VLan"),
	// 		NetworkType: to.Ptr("NoIsolation"),
	// 		Subnets: []*armrecoveryservicessiterecovery.Subnet{
	// 		},
	// 	},
	// }
}

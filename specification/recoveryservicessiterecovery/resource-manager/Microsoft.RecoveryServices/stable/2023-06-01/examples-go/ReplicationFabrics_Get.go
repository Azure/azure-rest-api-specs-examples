package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a29126ca8200a6c981a4e908e41fe55730df4cad/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationFabrics_Get.json
func ExampleReplicationFabricsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewReplicationFabricsClient().Get(ctx, "vault1", "resourceGroupPS1", "cloud1", &armrecoveryservicessiterecovery.ReplicationFabricsClientGetOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Fabric = armrecoveryservicessiterecovery.Fabric{
	// 	Name: to.Ptr("cloud1"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics"),
	// 	ID: to.Ptr("/Subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/resourceGroupPS1/providers/Microsoft.RecoveryServices/vaults/vault1/replicationFabrics/cloud1"),
	// 	Properties: &armrecoveryservicessiterecovery.FabricProperties{
	// 		BcdrState: to.Ptr("Valid"),
	// 		CustomDetails: &armrecoveryservicessiterecovery.HyperVSiteDetails{
	// 			InstanceType: to.Ptr("HyperVSite"),
	// 		},
	// 		EncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 		FriendlyName: to.Ptr("cloud1"),
	// 		Health: to.Ptr("Normal"),
	// 		HealthErrorDetails: []*armrecoveryservicessiterecovery.HealthError{
	// 		},
	// 		InternalIdentifier: to.Ptr("6d224fc6-f326-5d35-96de-fbf51efb3179"),
	// 		RolloverEncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 	},
	// }
}

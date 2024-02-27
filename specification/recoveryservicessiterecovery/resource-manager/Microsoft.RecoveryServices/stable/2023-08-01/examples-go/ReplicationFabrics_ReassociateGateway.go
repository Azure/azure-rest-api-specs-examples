package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples/ReplicationFabrics_ReassociateGateway.json
func ExampleReplicationFabricsClient_BeginReassociateGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicessiterecovery.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewReplicationFabricsClient().BeginReassociateGateway(ctx, "MadhaviVault", "MadhaviVRG", "GRACE-V2A-1", armrecoveryservicessiterecovery.FailoverProcessServerRequest{
		Properties: &armrecoveryservicessiterecovery.FailoverProcessServerRequestProperties{
			ContainerName:         to.Ptr("cloud_1f3c15af-2256-4568-9e06-e1ef4f728f75"),
			SourceProcessServerID: to.Ptr("AFA0EC54-1894-4E44-9CAB02DB8854B117"),
			TargetProcessServerID: to.Ptr("5D3ED340-85AE-C646-B338641E015DA405"),
			UpdateType:            to.Ptr("ServerLevel"),
			VMsToMigrate: []*string{
				to.Ptr("Vm1"),
				to.Ptr("Vm2")},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Fabric = armrecoveryservicessiterecovery.Fabric{
	// 	Name: to.Ptr("bc15edf300344660d9c2965f5d9225593d99734f6580613c997033abc3512a56"),
	// 	Type: to.Ptr("Microsoft.RecoveryServices/vaults/replicationFabrics"),
	// 	ID: to.Ptr("/Subscriptions/7c943c1b-5122-4097-90c8-861411bdd574/resourceGroups/MadhaviVRG/providers/Microsoft.RecoveryServices/vaults/MadhaviVault/replicationFabrics/bc15edf300344660d9c2965f5d9225593d99734f6580613c997033abc3512a56"),
	// 	Properties: &armrecoveryservicessiterecovery.FabricProperties{
	// 		BcdrState: to.Ptr("Valid"),
	// 		CustomDetails: &armrecoveryservicessiterecovery.VMwareDetails{
	// 			InstanceType: to.Ptr("VMware"),
	// 		},
	// 		EncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 		FriendlyName: to.Ptr("GRACE-V2A-1"),
	// 		Health: to.Ptr("Normal"),
	// 		HealthErrorDetails: []*armrecoveryservicessiterecovery.HealthError{
	// 		},
	// 		InternalIdentifier: to.Ptr("1f3c15af-2256-4568-9e06-e1ef4f728f75"),
	// 		RolloverEncryptionDetails: &armrecoveryservicessiterecovery.EncryptionDetails{
	// 			KekState: to.Ptr("None"),
	// 		},
	// 	},
	// }
}

package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/TriggerRecoveryPointMove_Post.json
func ExampleClient_BeginMoveRecoveryPoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicesbackup.NewClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginMoveRecoveryPoint(ctx,
		"testVault",
		"netsdktestrg",
		"Azure",
		"IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
		"VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1",
		"348916168024334",
		armrecoveryservicesbackup.MoveRPAcrossTiersRequest{
			ObjectType:     to.Ptr("MoveRPAcrossTiersRequest"),
			SourceTierType: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeHardenedRP),
			TargetTierType: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeArchivedRP),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

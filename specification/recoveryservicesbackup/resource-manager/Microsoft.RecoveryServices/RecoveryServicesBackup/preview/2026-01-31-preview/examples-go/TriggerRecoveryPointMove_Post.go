package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-01-31-preview/TriggerRecoveryPointMove_Post.json
func ExampleRecoveryServicesClient_BeginMoveRecoveryPoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRecoveryServicesClient().BeginMoveRecoveryPoint(ctx, "testVault", "netsdktestrg", "Azure", "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1", "348916168024334", armrecoveryservicesbackup.MoveRPAcrossTiersRequest{
		ObjectType:     to.Ptr("MoveRPAcrossTiersRequest"),
		SourceTierType: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeHardenedRP),
		TargetTierType: to.Ptr(armrecoveryservicesbackup.RecoveryPointTierTypeArchivedRP),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}

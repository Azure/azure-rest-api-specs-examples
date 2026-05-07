package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
)

// Generated from example definition: 2026-01-31-preview/BackupDataMove/PrepareDataMove_Post.json
func ExampleRecoveryServicesClient_BeginBMSPrepareDataMove() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesbackup.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRecoveryServicesClient().BeginBMSPrepareDataMove(ctx, "source-rsv", "sourceRG", armrecoveryservicesbackup.PrepareDataMoveRequest{
		DataMoveLevel:    to.Ptr(armrecoveryservicesbackup.DataMoveLevelVault),
		TargetRegion:     to.Ptr("USGov Virginia"),
		TargetResourceID: to.Ptr("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/targetRG/providers/Microsoft.RecoveryServices/vaults/target-rsv"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesbackup.RecoveryServicesClientBMSPrepareDataMoveResponse{
	// }
}

package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/BackupDataMove/PrepareDataMove_Post.json
func ExampleRecoveryServicesBackupClient_BeginBMSPrepareDataMove() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewRecoveryServicesBackupClient("<subscription-id>", cred, nil)
	poller, err := client.BeginBMSPrepareDataMove(ctx,
		"<vault-name>",
		"<resource-group-name>",
		armrecoveryservicesbackup.PrepareDataMoveRequest{
			DataMoveLevel:    armrecoveryservicesbackup.DataMoveLevelVault.ToPtr(),
			TargetRegion:     to.StringPtr("<target-region>"),
			TargetResourceID: to.StringPtr("<target-resource-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

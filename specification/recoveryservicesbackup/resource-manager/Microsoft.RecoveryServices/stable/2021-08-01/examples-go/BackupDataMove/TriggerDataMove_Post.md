```go
package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/BackupDataMove/TriggerDataMove_Post.json
func ExampleRecoveryServicesBackupClient_BeginBMSTriggerDataMove() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewRecoveryServicesBackupClient("<subscription-id>", cred, nil)
	poller, err := client.BeginBMSTriggerDataMove(ctx,
		"<vault-name>",
		"<resource-group-name>",
		armrecoveryservicesbackup.TriggerDataMoveRequest{
			CorrelationID:    to.StringPtr("<correlation-id>"),
			DataMoveLevel:    armrecoveryservicesbackup.DataMoveLevelVault.ToPtr(),
			SourceRegion:     to.StringPtr("<source-region>"),
			SourceResourceID: to.StringPtr("<source-resource-id>"),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservicesbackup%2Fv0.1.0/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/README.md) on how to add the SDK to your project and authenticate.

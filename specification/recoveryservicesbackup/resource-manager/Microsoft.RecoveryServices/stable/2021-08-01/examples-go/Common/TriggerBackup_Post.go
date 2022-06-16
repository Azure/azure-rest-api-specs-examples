package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/Common/TriggerBackup_Post.json
func ExampleBackupsClient_Trigger() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewBackupsClient("<subscription-id>", cred, nil)
	_, err = client.Trigger(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<fabric-name>",
		"<container-name>",
		"<protected-item-name>",
		armrecoveryservicesbackup.BackupRequestResource{
			Properties: &armrecoveryservicesbackup.IaasVMBackupRequest{
				BackupRequest: armrecoveryservicesbackup.BackupRequest{
					ObjectType: to.StringPtr("<object-type>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

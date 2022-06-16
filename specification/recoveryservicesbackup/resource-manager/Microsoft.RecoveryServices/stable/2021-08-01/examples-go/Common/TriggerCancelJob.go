package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/Common/TriggerCancelJob.json
func ExampleJobCancellationsClient_Trigger() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewJobCancellationsClient("<subscription-id>", cred, nil)
	_, err = client.Trigger(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<job-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

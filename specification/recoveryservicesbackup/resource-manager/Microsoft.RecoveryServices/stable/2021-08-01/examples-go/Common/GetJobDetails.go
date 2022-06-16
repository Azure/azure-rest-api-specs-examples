package armrecoveryservicesbackup_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup"
)

// x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/Common/GetJobDetails.json
func ExampleJobDetailsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicesbackup.NewJobDetailsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<vault-name>",
		"<resource-group-name>",
		"<job-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("JobResource.ID: %s\n", *res.ID)
}

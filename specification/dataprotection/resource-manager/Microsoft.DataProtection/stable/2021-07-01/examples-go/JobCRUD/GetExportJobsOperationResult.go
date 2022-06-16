package armdataprotection_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection"
)

// x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2021-07-01/examples/JobCRUD/GetExportJobsOperationResult.json
func ExampleExportJobsOperationResultClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdataprotection.NewExportJobsOperationResultClient("<subscription-id>", cred, nil)
	_, err = client.Get(ctx,
		"<resource-group-name>",
		"<vault-name>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

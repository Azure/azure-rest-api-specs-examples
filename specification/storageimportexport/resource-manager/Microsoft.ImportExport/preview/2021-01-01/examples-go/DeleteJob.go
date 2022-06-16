package armstorageimportexport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/DeleteJob.json
func ExampleJobsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorageimportexport.NewJobsClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		nil, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"myJob",
		"myResourceGroup",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

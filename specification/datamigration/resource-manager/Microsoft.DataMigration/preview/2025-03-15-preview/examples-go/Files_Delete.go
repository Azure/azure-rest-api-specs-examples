package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/Files_Delete.json
func ExampleFilesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewFilesClient().Delete(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkProject", "x114d023d8", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

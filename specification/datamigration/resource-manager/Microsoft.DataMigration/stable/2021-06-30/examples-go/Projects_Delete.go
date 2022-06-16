package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Projects_Delete.json
func ExampleProjectsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatamigration.NewProjectsClient("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"DmsSdkRg",
		"DmsSdkService",
		"DmsSdkProject",
		&armdatamigration.ProjectsClientDeleteOptions{DeleteRunningTasks: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

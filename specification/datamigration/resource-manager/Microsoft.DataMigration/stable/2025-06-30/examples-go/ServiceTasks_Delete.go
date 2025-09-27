package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/930e8030f5058d947fea4e2640725baab8a4561a/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2025-06-30/examples/ServiceTasks_Delete.json
func ExampleServiceTasksClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewServiceTasksClient().Delete(ctx, "DmsSdkRg", "DmsSdkService", "DmsSdkTask", &armdatamigration.ServiceTasksClientDeleteOptions{DeleteRunningTasks: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

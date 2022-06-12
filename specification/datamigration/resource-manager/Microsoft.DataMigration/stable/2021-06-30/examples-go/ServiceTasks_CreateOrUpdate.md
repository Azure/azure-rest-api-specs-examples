```go
package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/ServiceTasks_CreateOrUpdate.json
func ExampleServiceTasksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatamigration.NewServiceTasksClient("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"DmsSdkRg",
		"DmsSdkService",
		"DmsSdkTask",
		armdatamigration.ProjectTask{
			Properties: &armdatamigration.CheckOCIDriverTaskProperties{
				TaskType: to.Ptr("Service.Check.OCI"),
				Input: &armdatamigration.CheckOCIDriverTaskInput{
					ServerVersion: to.Ptr("NA"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatamigration%2Farmdatamigration%2Fv1.0.0/sdk/resourcemanager/datamigration/armdatamigration/README.md) on how to add the SDK to your project and authenticate.

package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2021-06-30/examples/Tasks_CreateOrUpdate.json
func ExampleTasksClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatamigration.NewTasksClient("fc04246f-04c5-437e-ac5e-206a19e7193f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"DmsSdkRg",
		"DmsSdkService",
		"DmsSdkProject",
		"DmsSdkTask",
		armdatamigration.ProjectTask{
			Properties: &armdatamigration.ConnectToTargetSQLDbTaskProperties{
				TaskType: to.Ptr("ConnectToTarget.SqlDb"),
				Input: &armdatamigration.ConnectToTargetSQLDbTaskInput{
					TargetConnectionInfo: &armdatamigration.SQLConnectionInfo{
						Type:                   to.Ptr("SqlConnectionInfo"),
						Password:               to.Ptr("testpassword"),
						UserName:               to.Ptr("testuser"),
						Authentication:         to.Ptr(armdatamigration.AuthenticationTypeSQLAuthentication),
						DataSource:             to.Ptr("ssma-test-server.database.windows.net"),
						EncryptConnection:      to.Ptr(true),
						TrustServerCertificate: to.Ptr(true),
					},
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

package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/GetSqlMigrationService.json
func ExampleSQLMigrationServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSQLMigrationServicesClient().Get(ctx, "testrg", "service1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SQLMigrationService = armdatamigration.SQLMigrationService{
	// 	Name: to.Ptr("service1"),
	// 	Type: to.Ptr("Microsoft.DataMigration/sqlMigrationServices"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/service1"),
	// 	Location: to.Ptr("northeurope"),
	// 	Tags: map[string]*string{
	// 		"myTag": to.Ptr("myVal"),
	// 	},
	// 	Properties: &armdatamigration.SQLMigrationServiceProperties{
	// 		IntegrationRuntimeState: to.Ptr("NeedRegistration"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}

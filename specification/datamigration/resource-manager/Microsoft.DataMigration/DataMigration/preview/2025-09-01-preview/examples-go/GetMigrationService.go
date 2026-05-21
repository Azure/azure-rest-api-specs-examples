package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v3"
)

// Generated from example definition: 2025-09-01-preview/GetMigrationService.json
func ExampleMigrationServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMigrationServicesClient().Get(ctx, "testrg", "service1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatamigration.MigrationServicesClientGetResponse{
	// 	MigrationService: armdatamigration.MigrationService{
	// 		Name: to.Ptr("service1"),
	// 		Type: to.Ptr("Microsoft.DataMigration/migrationServices"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/migrationServices/service1"),
	// 		Location: to.Ptr("northeurope"),
	// 		Properties: &armdatamigration.MigrationServiceProperties{
	// 			ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"myTag": to.Ptr("myVal"),
	// 		},
	// 	},
	// }
}

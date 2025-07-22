package armdatamigration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6c548b0bd279f5e233661b1c81fb5b61b19965cd/specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/ListByResourceGroupMigrationService.json
func ExampleMigrationServicesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatamigration.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMigrationServicesClient().NewListByResourceGroupPager("testrg", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.MigrationServiceListResult = armdatamigration.MigrationServiceListResult{
		// 	Value: []*armdatamigration.MigrationService{
		// 		{
		// 			Name: to.Ptr("service1"),
		// 			Type: to.Ptr("Microsoft.DataMigration/migrationServices"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/migrationServices/service1"),
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			Properties: &armdatamigration.MigrationServiceProperties{
		// 				ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("service2"),
		// 			Type: to.Ptr("Microsoft.DataMigration/migrationServices"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/migrationServices/service1"),
		// 			Location: to.Ptr("northeurope"),
		// 			Tags: map[string]*string{
		// 				"mytag": to.Ptr("myval"),
		// 			},
		// 			Properties: &armdatamigration.MigrationServiceProperties{
		// 				ProvisioningState: to.Ptr(armdatamigration.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

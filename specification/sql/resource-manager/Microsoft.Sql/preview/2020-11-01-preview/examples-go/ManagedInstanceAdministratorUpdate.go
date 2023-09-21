package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceAdministratorUpdate.json
func ExampleManagedInstanceAdministratorsClient_BeginCreateOrUpdate_updateAdministratorOfManagedInstance() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstanceAdministratorsClient().BeginCreateOrUpdate(ctx, "Default-SQL-SouthEastAsia", "managedInstance", armsql.AdministratorNameActiveDirectory, armsql.ManagedInstanceAdministrator{
		Properties: &armsql.ManagedInstanceAdministratorProperties{
			AdministratorType: to.Ptr(armsql.ManagedInstanceAdministratorTypeActiveDirectory),
			Login:             to.Ptr("bob@contoso.com"),
			Sid:               to.Ptr("44444444-3333-2222-1111-000000000000"),
			TenantID:          to.Ptr("55555555-4444-3333-2222-111111111111"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ManagedInstanceAdministrator = armsql.ManagedInstanceAdministrator{
	// 	Name: to.Ptr("ActiveDirectory"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/administrators"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/managedInstance/administrators/ActiveDirectory"),
	// 	Properties: &armsql.ManagedInstanceAdministratorProperties{
	// 		AdministratorType: to.Ptr(armsql.ManagedInstanceAdministratorTypeActiveDirectory),
	// 		Login: to.Ptr("bob@contoso.com"),
	// 		Sid: to.Ptr("44444444-3333-2222-1111-000000000000"),
	// 		TenantID: to.Ptr("55555555-4444-3333-2222-111111111111"),
	// 	},
	// }
}

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceAdministratorListByInstance.json
func ExampleManagedInstanceAdministratorsClient_NewListByInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedInstanceAdministratorsClient().NewListByInstancePager("Default-SQL-SouthEastAsia", "managedInstance", nil)
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
		// page.ManagedInstanceAdministratorListResult = armsql.ManagedInstanceAdministratorListResult{
		// 	Value: []*armsql.ManagedInstanceAdministrator{
		// 		{
		// 			Name: to.Ptr("ActiveDirectory"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/administrators"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/managedInstance/administrators/ActiveDirectory"),
		// 			Properties: &armsql.ManagedInstanceAdministratorProperties{
		// 				AdministratorType: to.Ptr(armsql.ManagedInstanceAdministratorTypeActiveDirectory),
		// 				Login: to.Ptr("bob@contoso.com"),
		// 				Sid: to.Ptr("44444444-3333-2222-1111-000000000000"),
		// 				TenantID: to.Ptr("55555555-4444-3333-2222-111111111111"),
		// 			},
		// 	}},
		// }
	}
}

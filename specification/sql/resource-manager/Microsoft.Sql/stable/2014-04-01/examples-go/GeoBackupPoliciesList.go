package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/GeoBackupPoliciesList.json
func ExampleGeoBackupPoliciesClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGeoBackupPoliciesClient().NewListByDatabasePager("sqlcrudtest-4799", "sqlcrudtest-5961", "testdw", nil)
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
		// page.GeoBackupPolicyListResult = armsql.GeoBackupPolicyListResult{
		// 	Value: []*armsql.GeoBackupPolicy{
		// 		{
		// 			Name: to.Ptr("Default"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/databases/geoBackupPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-4799/providers/Microsoft.Sql/servers/sqlcrudtest-5961/databases/testdw/geoBackupPolicies/Default"),
		// 			Location: to.Ptr("Central US"),
		// 			Properties: &armsql.GeoBackupPolicyProperties{
		// 				State: to.Ptr(armsql.GeoBackupPolicyStateEnabled),
		// 				StorageType: to.Ptr("Premium"),
		// 			},
		// 	}},
		// }
	}
}

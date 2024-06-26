package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolGeoBackupPolicies_List.json
func ExampleSQLPoolGeoBackupPoliciesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLPoolGeoBackupPoliciesClient().NewListPager("sqlcrudtest-4799", "sqlcrudtest-5961", "testdw", nil)
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
		// page.GeoBackupPolicyListResult = armsynapse.GeoBackupPolicyListResult{
		// 	Value: []*armsynapse.GeoBackupPolicy{
		// 		{
		// 			Name: to.Ptr("Default"),
		// 			Type: to.Ptr("Microsoft.Synapse/workspaces/sqlPools/geoBackupPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-4799/providers/Microsoft.Synapse/workspaces/sqlcrudtest-5961/sqlPools/testdw/geoBackupPolicies/Default"),
		// 			Location: to.Ptr("Central US"),
		// 			Properties: &armsynapse.GeoBackupPolicyProperties{
		// 				State: to.Ptr(armsynapse.GeoBackupPolicyStateEnabled),
		// 				StorageType: to.Ptr("Premium"),
		// 			},
		// 	}},
		// }
	}
}

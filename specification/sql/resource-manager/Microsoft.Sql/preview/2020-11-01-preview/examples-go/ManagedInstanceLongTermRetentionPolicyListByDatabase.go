package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceLongTermRetentionPolicyListByDatabase.json
func ExampleManagedInstanceLongTermRetentionPoliciesClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedInstanceLongTermRetentionPoliciesClient().NewListByDatabasePager("testResourceGroup", "testInstance", "testDatabase", nil)
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
		// page.ManagedInstanceLongTermRetentionPolicyListResult = armsql.ManagedInstanceLongTermRetentionPolicyListResult{
		// 	Value: []*armsql.ManagedInstanceLongTermRetentionPolicy{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Sql/resourceGroups/managedInstances/databases/backupLongTermRetentionPolicies"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testResourceGroup/providers/Microsoft.Sql/managedInstances/testInstance/databases/testDatabase/backupLongTermRetentionPolicies/default"),
		// 			Properties: &armsql.BaseLongTermRetentionPolicyProperties{
		// 				MonthlyRetention: to.Ptr("P1Y"),
		// 				WeekOfYear: to.Ptr[int32](5),
		// 				WeeklyRetention: to.Ptr("P1M"),
		// 				YearlyRetention: to.Ptr("P5Y"),
		// 			},
		// 	}},
		// }
	}
}

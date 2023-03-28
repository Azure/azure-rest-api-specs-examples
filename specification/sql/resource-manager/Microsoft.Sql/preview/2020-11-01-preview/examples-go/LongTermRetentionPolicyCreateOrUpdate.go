package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/LongTermRetentionPolicyCreateOrUpdate.json
func ExampleLongTermRetentionPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLongTermRetentionPoliciesClient().BeginCreateOrUpdate(ctx, "resourceGroup", "testserver", "testDatabase", armsql.LongTermRetentionPolicyNameDefault, armsql.LongTermRetentionPolicy{
		Properties: &armsql.BaseLongTermRetentionPolicyProperties{
			MonthlyRetention: to.Ptr("P1Y"),
			WeekOfYear:       to.Ptr[int32](5),
			WeeklyRetention:  to.Ptr("P1M"),
			YearlyRetention:  to.Ptr("P5Y"),
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
	// res.LongTermRetentionPolicy = armsql.LongTermRetentionPolicy{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Sql/resourceGroups/servers/databases/backupLongTermRetentionPolicies"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/resourceGroups/resourceGroup/servers/testserver/databases/testDatabase/backupLongTermRetentionPolicies/default"),
	// 	Properties: &armsql.BaseLongTermRetentionPolicyProperties{
	// 		MonthlyRetention: to.Ptr("P1Y"),
	// 		WeekOfYear: to.Ptr[int32](5),
	// 		WeeklyRetention: to.Ptr("P1M"),
	// 		YearlyRetention: to.Ptr("P5Y"),
	// 	},
	// }
}

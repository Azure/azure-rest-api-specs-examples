package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/GetWorkloadGroup.json
func ExampleWorkloadGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkloadGroupsClient().Get(ctx, "Default-SQL-SouthEastAsia", "testsvr", "testdb", "smallrc", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkloadGroup = armsql.WorkloadGroup{
	// 	Name: to.Ptr("smallrc"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/workloadGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb/workloadGroups/smallrc"),
	// 	Properties: &armsql.WorkloadGroupProperties{
	// 		Importance: to.Ptr("normal"),
	// 		MaxResourcePercent: to.Ptr[int32](100),
	// 		MaxResourcePercentPerRequest: to.Ptr[float64](3),
	// 		MinResourcePercent: to.Ptr[int32](0),
	// 		MinResourcePercentPerRequest: to.Ptr[float64](3),
	// 		QueryExecutionTimeout: to.Ptr[int32](0),
	// 	},
	// }
}

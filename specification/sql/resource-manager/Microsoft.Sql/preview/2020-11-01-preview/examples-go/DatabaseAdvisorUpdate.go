package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseAdvisorUpdate.json
func ExampleDatabaseAdvisorsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseAdvisorsClient().Update(ctx, "workloadinsight-demos", "misosisvr", "IndexAdvisor_test_3", "CreateIndex", armsql.Advisor{
		Properties: &armsql.AdvisorProperties{
			AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Advisor = armsql.Advisor{
	// 	Name: to.Ptr("CreateIndex"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/advisors"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/CreateIndex"),
	// 	Properties: &armsql.AdvisorProperties{
	// 		AdvisorStatus: to.Ptr(armsql.AdvisorStatusGA),
	// 		AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 		LastChecked: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-20T00:24:39.000Z"); return t}()),
	// 		RecommendationsStatus: to.Ptr("Ok"),
	// 	},
	// }
}

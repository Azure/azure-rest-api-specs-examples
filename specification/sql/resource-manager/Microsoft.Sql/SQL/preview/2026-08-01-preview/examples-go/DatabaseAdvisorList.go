package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/DatabaseAdvisorList.json
func ExampleDatabaseAdvisorsClient_ListByDatabase_listOfDatabaseAdvisors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseAdvisorsClient().ListByDatabase(ctx, "workloadinsight-demos", "misosisvr", "IndexAdvisor_test_3", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.DatabaseAdvisorsClientListByDatabaseResponse{
	// 	AdvisorArray: []*armsql.Advisor{
	// 		{
	// 			Name: to.Ptr("CreateIndex"),
	// 			Type: to.Ptr("Microsoft.Sql/servers/databases/advisors"),
	// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/CreateIndex"),
	// 			Kind: to.Ptr(""),
	// 			Location: to.Ptr("East Asia"),
	// 			Properties: &armsql.AdvisorProperties{
	// 				AdvisorStatus: to.Ptr(armsql.AdvisorStatusGA),
	// 				AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 				AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDatabase),
	// 				LastChecked: to.Ptr(time.Date(2017, time.June, 20, 0, 24, 39, 0, time.UTC)),
	// 				RecommendationsStatus: to.Ptr("Ok"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("DropIndex"),
	// 			Type: to.Ptr("Microsoft.Sql/servers/databases/advisors"),
	// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/DropIndex"),
	// 			Kind: to.Ptr(""),
	// 			Location: to.Ptr("East Asia"),
	// 			Properties: &armsql.AdvisorProperties{
	// 				AdvisorStatus: to.Ptr(armsql.AdvisorStatusGA),
	// 				AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 				AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDatabase),
	// 				LastChecked: to.Ptr(time.Date(2017, time.June, 19, 20, 42, 43, 0, time.UTC)),
	// 				RecommendationsStatus: to.Ptr("DbSeemsTuned"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("DbParameterization"),
	// 			Type: to.Ptr("Microsoft.Sql/servers/databases/advisors"),
	// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/DbParameterization"),
	// 			Kind: to.Ptr(""),
	// 			Location: to.Ptr("East Asia"),
	// 			Properties: &armsql.AdvisorProperties{
	// 				AdvisorStatus: to.Ptr(armsql.AdvisorStatusGA),
	// 				AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusEnabled),
	// 				AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDatabase),
	// 				LastChecked: to.Ptr(time.Date(2017, time.June, 19, 19, 13, 3, 0, time.UTC)),
	// 				RecommendationsStatus: to.Ptr("DbParameterizationIssue"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("SchemaIssue"),
	// 			Type: to.Ptr("Microsoft.Sql/servers/databases/advisors"),
	// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/SchemaIssue"),
	// 			Kind: to.Ptr(""),
	// 			Location: to.Ptr("East Asia"),
	// 			Properties: &armsql.AdvisorProperties{
	// 				AdvisorStatus: to.Ptr(armsql.AdvisorStatusPublicPreview),
	// 				AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 				AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDefault),
	// 				LastChecked: to.Ptr(time.Date(2017, time.June, 20, 14, 0, 32, 0, time.UTC)),
	// 				RecommendationsStatus: to.Ptr("SchemaIsNotConsistent"),
	// 			},
	// 		},
	// 		{
	// 			Name: to.Ptr("ForceLastGoodPlan"),
	// 			Type: to.Ptr("Microsoft.Sql/servers/databases/advisors"),
	// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/ForceLastGoodPlan"),
	// 			Kind: to.Ptr(""),
	// 			Location: to.Ptr("East Asia"),
	// 			Properties: &armsql.AdvisorProperties{
	// 				AdvisorStatus: to.Ptr(armsql.AdvisorStatusPrivatePreview),
	// 				AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 				AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDefault),
	// 			},
	// 		},
	// 	},
	// }
}

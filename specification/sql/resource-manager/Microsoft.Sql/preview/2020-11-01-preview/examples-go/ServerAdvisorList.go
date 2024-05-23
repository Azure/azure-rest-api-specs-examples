package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ServerAdvisorList.json
func ExampleServerAdvisorsClient_ListByServer_listOfServerAdvisors() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerAdvisorsClient().ListByServer(ctx, "workloadinsight-demos", "misosisvr", &armsql.ServerAdvisorsClientListByServerOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AdvisorArray = []*armsql.Advisor{
	// 	{
	// 		Name: to.Ptr("CreateIndex"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/advisors"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/CreateIndex"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.AdvisorProperties{
	// 			AdvisorStatus: to.Ptr(armsql.AdvisorStatusGA),
	// 			AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 			AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDefault),
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("DropIndex"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/advisors"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/DropIndex"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.AdvisorProperties{
	// 			AdvisorStatus: to.Ptr(armsql.AdvisorStatusGA),
	// 			AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 			AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDefault),
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("DbParameterization"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/advisors"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/DbParameterization"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.AdvisorProperties{
	// 			AdvisorStatus: to.Ptr(armsql.AdvisorStatusGA),
	// 			AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 			AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDefault),
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("SchemaIssue"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/advisors"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/SchemaIssue"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.AdvisorProperties{
	// 			AdvisorStatus: to.Ptr(armsql.AdvisorStatusPublicPreview),
	// 			AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 			AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDefault),
	// 		},
	// 	},
	// 	{
	// 		Name: to.Ptr("ForceLastGoodPlan"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/advisors"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/advisors/ForceLastGoodPlan"),
	// 		Kind: to.Ptr(""),
	// 		Location: to.Ptr("East Asia"),
	// 		Properties: &armsql.AdvisorProperties{
	// 			AdvisorStatus: to.Ptr(armsql.AdvisorStatusPrivatePreview),
	// 			AutoExecuteStatus: to.Ptr(armsql.AutoExecuteStatusDisabled),
	// 			AutoExecuteStatusInheritedFrom: to.Ptr(armsql.AutoExecuteStatusInheritedFromDefault),
	// 		},
	// }}
}

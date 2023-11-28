package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DatabaseRecommendedActionUpdate.json
func ExampleDatabaseRecommendedActionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabaseRecommendedActionsClient().Update(ctx, "workloadinsight-demos", "misosisvr", "IndexAdvisor_test_3", "CreateIndex", "IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB", armsql.RecommendedAction{
		Properties: &armsql.RecommendedActionProperties{
			State: &armsql.RecommendedActionStateInfo{
				CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStatePending),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RecommendedAction = armsql.RecommendedAction{
	// 	Name: to.Ptr("IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/advisors/recommendedActions"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workloadinsight-demos/providers/Microsoft.Sql/servers/misosisvr/databases/IndexAdvisor_test_3/advisors/CreateIndex/recommendedActions/IR_[CRM]_[DataPoints]_4821CD2F9510D98184BB"),
	// 	Kind: to.Ptr(""),
	// 	Location: to.Ptr("East Asia"),
	// 	Properties: &armsql.RecommendedActionProperties{
	// 		ErrorDetails: &armsql.RecommendedActionErrorInfo{
	// 		},
	// 		EstimatedImpact: []*armsql.RecommendedActionImpactRecord{
	// 			{
	// 				AbsoluteValue: to.Ptr[float64](1440),
	// 				DimensionName: to.Ptr("ActionDuration"),
	// 				Unit: to.Ptr("Seconds"),
	// 			},
	// 			{
	// 				AbsoluteValue: to.Ptr[float64](209.3125),
	// 				DimensionName: to.Ptr("SpaceChange"),
	// 				Unit: to.Ptr("Megabytes"),
	// 		}},
	// 		ExecuteActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedByUser),
	// 		ExecuteActionInitiatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-20T15:11:15.000Z"); return t}()),
	// 		ImplementationDetails: &armsql.RecommendedActionImplementationInfo{
	// 			Method: to.Ptr(armsql.ImplementationMethodTSQL),
	// 			Script: to.Ptr("CREATE NONCLUSTERED INDEX [nci_wi_DataPoints_B892614093BAC56295EF6018BD4CB51B] ON [CRM].[DataPoints] ([Name],[Money],[Power]) INCLUDE ([Hour], [System], [LastChanged]) WITH (ONLINE = ON)"),
	// 		},
	// 		IsArchivedAction: to.Ptr(false),
	// 		IsExecutableAction: to.Ptr(true),
	// 		IsRevertableAction: to.Ptr(true),
	// 		LastRefresh: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:04.000Z"); return t}()),
	// 		ObservedImpact: []*armsql.RecommendedActionImpactRecord{
	// 		},
	// 		RecommendationReason: to.Ptr(""),
	// 		Score: to.Ptr[int32](1),
	// 		State: &armsql.RecommendedActionStateInfo{
	// 			ActionInitiatedBy: to.Ptr(armsql.RecommendedActionInitiatedByUser),
	// 			CurrentValue: to.Ptr(armsql.RecommendedActionCurrentStatePending),
	// 			LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-20T15:11:15.000Z"); return t}()),
	// 		},
	// 		TimeSeries: []*armsql.RecommendedActionMetricInfo{
	// 		},
	// 		ValidSince: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-01T14:38:04.000Z"); return t}()),
	// 		Details: map[string]any{
	// 			"schema": "[CRM]",
	// 			"includedColumns": "[Hour], [System], [LastChanged]",
	// 			"indexColumns": "[Name],[Money],[Power]",
	// 			"indexName": "nci_wi_DataPoints_B892614093BAC56295EF6018BD4CB51B",
	// 			"indexType": "NONCLUSTERED",
	// 			"table": "[DataPoints]",
	// 		},
	// 	},
	// }
}

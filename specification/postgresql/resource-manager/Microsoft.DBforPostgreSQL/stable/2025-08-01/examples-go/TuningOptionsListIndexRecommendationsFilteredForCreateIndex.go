package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/TuningOptionsListIndexRecommendationsFilteredForCreateIndex.json
func ExampleTuningOptionsClient_NewListRecommendationsPager_listAvailableIndexRecommendationsFilteredToExclusivelyGetThoseOfCreateIndexType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTuningOptionsClient().NewListRecommendationsPager("exampleresourcegroup", "exampleserver", armpostgresqlflexibleservers.TuningOptionParameterEnumIndex, &armpostgresqlflexibleservers.TuningOptionsClientListRecommendationsOptions{RecommendationType: to.Ptr(armpostgresqlflexibleservers.RecommendationTypeParameterEnumCreateIndex)})
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
		// page.ObjectRecommendationList = armpostgresqlflexibleservers.ObjectRecommendationList{
		// 	Value: []*armpostgresqlflexibleservers.ObjectRecommendation{
		// 		{
		// 			Name: to.Ptr("CreateIndex_ecommerce_public_ps_suppkey_idx"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/index"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/tuningOptions/index/recommendations/1"),
		// 			Kind: to.Ptr(""),
		// 			Properties: &armpostgresqlflexibleservers.ObjectRecommendationProperties{
		// 				AnalyzedWorkload: &armpostgresqlflexibleservers.ObjectRecommendationPropertiesAnalyzedWorkload{
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 					QueryCount: to.Ptr[int32](25),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T19:30:22.123Z"); return t}()),
		// 				},
		// 				EstimatedImpact: []*armpostgresqlflexibleservers.ImpactRecord{
		// 					{
		// 						AbsoluteValue: to.Ptr[float64](15.3671875),
		// 						DimensionName: to.Ptr("IndexSize"),
		// 						Unit: to.Ptr("MB"),
		// 					},
		// 					{
		// 						AbsoluteValue: to.Ptr[float64](99.67668452400453),
		// 						DimensionName: to.Ptr("QueryCostImprovement"),
		// 						QueryID: to.Ptr[int64](-3775242682326862300),
		// 						Unit: to.Ptr("Percentage"),
		// 					},
		// 					{
		// 						AbsoluteValue: to.Ptr[float64](85.56742436827899),
		// 						DimensionName: to.Ptr("QueryCostImprovement"),
		// 						QueryID: to.Ptr[int64](6829938984138799000),
		// 						Unit: to.Ptr("Percentage"),
		// 				}},
		// 				ImplementationDetails: &armpostgresqlflexibleservers.ObjectRecommendationPropertiesImplementationDetails{
		// 					Method: to.Ptr("SQL"),
		// 					Script: to.Ptr("create index concurrently ps_suppkey_idx on public.partsupp(ps_suppkey)"),
		// 				},
		// 				ImprovedQueryIDs: []*int64{
		// 					to.Ptr[int64](-3775242682326862300),
		// 					to.Ptr[int64](6829938984138799000)},
		// 					InitialRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 					LastRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 					RecommendationReason: to.Ptr("Column \"partsupp\".\"ps_suppkey\" appear in Join On clause(s) in query -3775242682326862475; Column \"partsupp\".\"ps_suppkey\" appear in Join On clause(s) in query 6829938984138799352;"),
		// 					RecommendationType: to.Ptr(armpostgresqlflexibleservers.RecommendationTypeEnumCreateIndex),
		// 					TimesRecommended: to.Ptr[int32](1),
		// 					Details: &armpostgresqlflexibleservers.ObjectRecommendationDetails{
		// 						Schema: to.Ptr("public"),
		// 						DatabaseName: to.Ptr("ecommerce"),
		// 						IndexColumns: []*string{
		// 							to.Ptr("\"partsupp\".\"ps_suppkey\"")},
		// 							IndexName: to.Ptr("ps_suppkey_idx"),
		// 							IndexType: to.Ptr("BTREE"),
		// 							Table: to.Ptr("partsupp"),
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("CreateIndex_ecommerce_public_ps_partkey_idx"),
		// 					Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/index"),
		// 					ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/tuningOptions/index/recommendations/2"),
		// 					Kind: to.Ptr(""),
		// 					Properties: &armpostgresqlflexibleservers.ObjectRecommendationProperties{
		// 						AnalyzedWorkload: &armpostgresqlflexibleservers.ObjectRecommendationPropertiesAnalyzedWorkload{
		// 							EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 							QueryCount: to.Ptr[int32](25),
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T19:30:22.123Z"); return t}()),
		// 						},
		// 						EstimatedImpact: []*armpostgresqlflexibleservers.ImpactRecord{
		// 							{
		// 								AbsoluteValue: to.Ptr[float64](15.3671875),
		// 								DimensionName: to.Ptr("IndexSize"),
		// 								Unit: to.Ptr("MB"),
		// 							},
		// 							{
		// 								AbsoluteValue: to.Ptr[float64](99.67668452400453),
		// 								DimensionName: to.Ptr("QueryCostImprovement"),
		// 								QueryID: to.Ptr[int64](-3775242682326862300),
		// 								Unit: to.Ptr("Percentage"),
		// 							},
		// 							{
		// 								AbsoluteValue: to.Ptr[float64](79.06603712430707),
		// 								DimensionName: to.Ptr("QueryCostImprovement"),
		// 								QueryID: to.Ptr[int64](4735984994430715000),
		// 								Unit: to.Ptr("Percentage"),
		// 						}},
		// 						ImplementationDetails: &armpostgresqlflexibleservers.ObjectRecommendationPropertiesImplementationDetails{
		// 							Method: to.Ptr("SQL"),
		// 							Script: to.Ptr("create index concurrently ps_partkey_idx on public.partsupp(ps_partkey)"),
		// 						},
		// 						ImprovedQueryIDs: []*int64{
		// 							to.Ptr[int64](-3775242682326862300),
		// 							to.Ptr[int64](4735984994430715000)},
		// 							InitialRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 							LastRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 							RecommendationReason: to.Ptr("Column \"partsupp\".\"ps_partkey\" appear in Equal Predicate clause(s) in query -3775242682326862475; Column \"partsupp\".\"ps_partkey\" appear in Join On clause(s) in query 4735984994430714735;"),
		// 							RecommendationType: to.Ptr(armpostgresqlflexibleservers.RecommendationTypeEnumCreateIndex),
		// 							TimesRecommended: to.Ptr[int32](1),
		// 							Details: &armpostgresqlflexibleservers.ObjectRecommendationDetails{
		// 								Schema: to.Ptr("public"),
		// 								DatabaseName: to.Ptr("ecommerce"),
		// 								IndexColumns: []*string{
		// 									to.Ptr("\"partsupp\".\"ps_partkey\"")},
		// 									IndexName: to.Ptr("ps_partkey_idx"),
		// 									IndexType: to.Ptr("BTREE"),
		// 									Table: to.Ptr("partsupp"),
		// 								},
		// 							},
		// 					}},
		// 				}
	}
}

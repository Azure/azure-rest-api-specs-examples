package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1f4d539964453ce8008e4b069e59885e12ba441/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2025-01-01-preview/examples/TuningIndex_GetRecommendations.json
func ExampleTuningIndexClient_NewListRecommendationsPager_tuningIndexListRecommendations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTuningIndexClient().NewListRecommendationsPager("testrg", "pgtestsvc4", armpostgresqlflexibleservers.TuningOptionEnumIndex, &armpostgresqlflexibleservers.TuningIndexClientListRecommendationsOptions{RecommendationType: nil})
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
		// page.IndexRecommendationListResult = armpostgresqlflexibleservers.IndexRecommendationListResult{
		// 	Value: []*armpostgresqlflexibleservers.IndexRecommendationResource{
		// 		{
		// 			Name: to.Ptr("CreateIndex_ecommerce_public_c_custkey_idx"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/index"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc4/tuningOptions/index/recommendations/14"),
		// 			Properties: &armpostgresqlflexibleservers.IndexRecommendationResourceProperties{
		// 				AnalyzedWorkload: &armpostgresqlflexibleservers.IndexRecommendationResourcePropertiesAnalyzedWorkload{
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T19:14:34.412Z"); return t}()),
		// 					QueryCount: to.Ptr[int32](25),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T07:14:34.412Z"); return t}()),
		// 				},
		// 				EstimatedImpact: []*armpostgresqlflexibleservers.ImpactRecord{
		// 					{
		// 						AbsoluteValue: to.Ptr[float64](2.890625),
		// 						DimensionName: to.Ptr("IndexSize"),
		// 						Unit: to.Ptr("MB"),
		// 					},
		// 					{
		// 						AbsoluteValue: to.Ptr[float64](56.75424930485546),
		// 						DimensionName: to.Ptr("QueryCostImprovement"),
		// 						QueryID: to.Ptr[int64](-5265741861999699000),
		// 						Unit: to.Ptr("Percentage"),
		// 					},
		// 					{
		// 						AbsoluteValue: to.Ptr[float64](43.67826245896799),
		// 						DimensionName: to.Ptr("QueryCostImprovement"),
		// 						QueryID: to.Ptr[int64](234118651516953380),
		// 						Unit: to.Ptr("Percentage"),
		// 					},
		// 					{
		// 						AbsoluteValue: to.Ptr[float64](10.176593636567572),
		// 						DimensionName: to.Ptr("QueryCostImprovement"),
		// 						QueryID: to.Ptr[int64](-632598438706436700),
		// 						Unit: to.Ptr("Percentage"),
		// 				}},
		// 				ImplementationDetails: &armpostgresqlflexibleservers.IndexRecommendationResourcePropertiesImplementationDetails{
		// 					Method: to.Ptr("SQL"),
		// 					Script: to.Ptr("create index concurrently c_custkey_idx on public.customer(c_custkey)"),
		// 				},
		// 				ImprovedQueryIDs: []*int64{
		// 					to.Ptr[int64](-5265741861999699000),
		// 					to.Ptr[int64](-632598438706436700),
		// 					to.Ptr[int64](234118651516953380)},
		// 					InitialRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T08:13:55.382Z"); return t}()),
		// 					LastRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T19:14:34.412Z"); return t}()),
		// 					RecommendationReason: to.Ptr("Column \"customer\".\"c_custkey\" appear in Join On clause(s) in query -5265741861999699191; Column \"customer\".\"c_custkey\" appear in Join On clause(s) in query -632598438706436788; Column \"customer\".\"c_custkey\" appear in Join On clause(s) in query 234118651516953391;"),
		// 					RecommendationType: to.Ptr(armpostgresqlflexibleservers.RecommendationTypeEnumCreateIndex),
		// 					TimesRecommended: to.Ptr[int32](2),
		// 					Details: &armpostgresqlflexibleservers.IndexRecommendationDetails{
		// 						Schema: to.Ptr("public"),
		// 						DatabaseName: to.Ptr("ecommerce"),
		// 						IndexColumns: []*string{
		// 							to.Ptr("\"customer\".\"c_custkey\"")},
		// 							IndexName: to.Ptr("c_custkey_idx"),
		// 							IndexType: to.Ptr("BTREE"),
		// 							Table: to.Ptr("customer"),
		// 						},
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("CreateIndex_ecommerce_public_l_orderkey_idx"),
		// 					Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/index"),
		// 					ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc4/tuningOptions/index/recommendations/11"),
		// 					Properties: &armpostgresqlflexibleservers.IndexRecommendationResourceProperties{
		// 						AnalyzedWorkload: &armpostgresqlflexibleservers.IndexRecommendationResourcePropertiesAnalyzedWorkload{
		// 							EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T19:14:34.412Z"); return t}()),
		// 							QueryCount: to.Ptr[int32](25),
		// 							StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T07:14:34.412Z"); return t}()),
		// 						},
		// 						EstimatedImpact: []*armpostgresqlflexibleservers.ImpactRecord{
		// 							{
		// 								AbsoluteValue: to.Ptr[float64](115.21875),
		// 								DimensionName: to.Ptr("IndexSize"),
		// 								Unit: to.Ptr("MB"),
		// 							},
		// 							{
		// 								AbsoluteValue: to.Ptr[float64](63.42250763539086),
		// 								DimensionName: to.Ptr("QueryCostImprovement"),
		// 								QueryID: to.Ptr[int64](-1887467119476231700),
		// 								Unit: to.Ptr("Percentage"),
		// 							},
		// 							{
		// 								AbsoluteValue: to.Ptr[float64](56.75424930485546),
		// 								DimensionName: to.Ptr("QueryCostImprovement"),
		// 								QueryID: to.Ptr[int64](-5265741861999699000),
		// 								Unit: to.Ptr("Percentage"),
		// 							},
		// 							{
		// 								AbsoluteValue: to.Ptr[float64](77.23764140869261),
		// 								DimensionName: to.Ptr("QueryCostImprovement"),
		// 								QueryID: to.Ptr[int64](-5259344436787616000),
		// 								Unit: to.Ptr("Percentage"),
		// 							},
		// 							{
		// 								AbsoluteValue: to.Ptr[float64](43.67826245896799),
		// 								DimensionName: to.Ptr("QueryCostImprovement"),
		// 								QueryID: to.Ptr[int64](234118651516953380),
		// 								Unit: to.Ptr("Percentage"),
		// 							},
		// 							{
		// 								AbsoluteValue: to.Ptr[float64](75.73857546652877),
		// 								DimensionName: to.Ptr("QueryCostImprovement"),
		// 								QueryID: to.Ptr[int64](-2174336608863105800),
		// 								Unit: to.Ptr("Percentage"),
		// 						}},
		// 						ImplementationDetails: &armpostgresqlflexibleservers.IndexRecommendationResourcePropertiesImplementationDetails{
		// 							Method: to.Ptr("SQL"),
		// 							Script: to.Ptr("create index concurrently l_orderkey_idx on public.lineitem(l_orderkey)"),
		// 						},
		// 						ImprovedQueryIDs: []*int64{
		// 							to.Ptr[int64](-5265741861999699000),
		// 							to.Ptr[int64](-5259344436787616000),
		// 							to.Ptr[int64](-2174336608863105800),
		// 							to.Ptr[int64](-1887467119476231700),
		// 							to.Ptr[int64](234118651516953380)},
		// 							InitialRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T08:13:55.382Z"); return t}()),
		// 							LastRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T19:14:34.412Z"); return t}()),
		// 							RecommendationReason: to.Ptr("Column \"lineitem\".\"l_orderkey\" appear in Join On clause(s) in query -5265741861999699191; Column \"lineitem\".\"l_orderkey\" appear in Join On clause(s) in query -5259344436787615838; Column \"lineitem\".\"l_orderkey\" appear in Join On clause(s) in query -2174336608863105762; Column \"lineitem\".\"l_orderkey\" appear in Equal Predicate clause(s) in query -1887467119476231661; Column \"lineitem\".\"l_orderkey\" appear in Join On clause(s) in query 234118651516953391;"),
		// 							RecommendationType: to.Ptr(armpostgresqlflexibleservers.RecommendationTypeEnumCreateIndex),
		// 							TimesRecommended: to.Ptr[int32](2),
		// 							Details: &armpostgresqlflexibleservers.IndexRecommendationDetails{
		// 								Schema: to.Ptr("public"),
		// 								DatabaseName: to.Ptr("ecommerce"),
		// 								IndexColumns: []*string{
		// 									to.Ptr("\"lineitem\".\"l_orderkey\"")},
		// 									IndexName: to.Ptr("l_orderkey_idx"),
		// 									IndexType: to.Ptr("BTREE"),
		// 									Table: to.Ptr("lineitem"),
		// 								},
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("CreateIndex_ecommerce_public_l_orderkey_l_shipdate_idx"),
		// 							Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/index"),
		// 							ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/pgtestsvc4/tuningOptions/index/recommendations/18"),
		// 							Properties: &armpostgresqlflexibleservers.IndexRecommendationResourceProperties{
		// 								AnalyzedWorkload: &armpostgresqlflexibleservers.IndexRecommendationResourcePropertiesAnalyzedWorkload{
		// 									EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T19:14:34.412Z"); return t}()),
		// 									QueryCount: to.Ptr[int32](25),
		// 									StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T07:14:34.412Z"); return t}()),
		// 								},
		// 								EstimatedImpact: []*armpostgresqlflexibleservers.ImpactRecord{
		// 									{
		// 										AbsoluteValue: to.Ptr[float64](161.140625),
		// 										DimensionName: to.Ptr("IndexSize"),
		// 										Unit: to.Ptr("MB"),
		// 									},
		// 									{
		// 										AbsoluteValue: to.Ptr[float64](43.07835695205126),
		// 										DimensionName: to.Ptr("QueryCostImprovement"),
		// 										QueryID: to.Ptr[int64](-6263488632059809000),
		// 										Unit: to.Ptr("Percentage"),
		// 									},
		// 									{
		// 										AbsoluteValue: to.Ptr[float64](83.06572136980492),
		// 										DimensionName: to.Ptr("QueryCostImprovement"),
		// 										QueryID: to.Ptr[int64](-2436244725743338500),
		// 										Unit: to.Ptr("Percentage"),
		// 								}},
		// 								ImplementationDetails: &armpostgresqlflexibleservers.IndexRecommendationResourcePropertiesImplementationDetails{
		// 									Method: to.Ptr("SQL"),
		// 									Script: to.Ptr("create index concurrently l_orderkey_l_shipdate_idx on public.lineitem(l_orderkey,l_shipdate)"),
		// 								},
		// 								ImprovedQueryIDs: []*int64{
		// 									to.Ptr[int64](-6263488632059809000),
		// 									to.Ptr[int64](-2436244725743338500)},
		// 									InitialRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T08:13:55.382Z"); return t}()),
		// 									LastRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-05-31T19:14:34.412Z"); return t}()),
		// 									RecommendationReason: to.Ptr("Column \"lineitem\".\"l_orderkey\" appear in Join On clause(s) in query -6263488632059808423; Column \"lineitem\".\"l_shipdate\" appear in Non-Equal Predicate clause(s) in query -6263488632059808423;"),
		// 									RecommendationType: to.Ptr(armpostgresqlflexibleservers.RecommendationTypeEnumCreateIndex),
		// 									TimesRecommended: to.Ptr[int32](2),
		// 									Details: &armpostgresqlflexibleservers.IndexRecommendationDetails{
		// 										Schema: to.Ptr("public"),
		// 										DatabaseName: to.Ptr("ecommerce"),
		// 										IndexColumns: []*string{
		// 											to.Ptr("\"lineitem\".\"l_orderkey\""),
		// 											to.Ptr("\"lineitem\".\"l_shipdate\"")},
		// 											IndexName: to.Ptr("l_orderkey_l_shipdate_idx"),
		// 											IndexType: to.Ptr("BTREE"),
		// 											Table: to.Ptr("lineitem"),
		// 										},
		// 									},
		// 							}},
		// 						}
	}
}

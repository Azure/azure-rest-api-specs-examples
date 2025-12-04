package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/TuningOptionsListTableRecommendationsFilteredForAnalyzeTable.json
func ExampleTuningOptionsClient_NewListRecommendationsPager_listAvailableTableRecommendationsFilteredToExclusivelyGetThoseOfAnalyzeTableType() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTuningOptionsClient().NewListRecommendationsPager("exampleresourcegroup", "exampleserver", armpostgresqlflexibleservers.TuningOptionParameterEnumTable, &armpostgresqlflexibleservers.TuningOptionsClientListRecommendationsOptions{RecommendationType: to.Ptr(armpostgresqlflexibleservers.RecommendationTypeParameterEnumAnalyzeTable)})
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
		// 			Name: to.Ptr("Analyze_postgres_public_nation"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/table"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/tuningOptions/table/recommendations/1"),
		// 			Kind: to.Ptr(""),
		// 			Properties: &armpostgresqlflexibleservers.ObjectRecommendationProperties{
		// 				AnalyzedWorkload: &armpostgresqlflexibleservers.ObjectRecommendationPropertiesAnalyzedWorkload{
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 					QueryCount: to.Ptr[int32](22),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T19:30:22.123Z"); return t}()),
		// 				},
		// 				EstimatedImpact: []*armpostgresqlflexibleservers.ImpactRecord{
		// 				},
		// 				ImplementationDetails: &armpostgresqlflexibleservers.ObjectRecommendationPropertiesImplementationDetails{
		// 					Method: to.Ptr("SQL"),
		// 					Script: to.Ptr("analyze table public.nation"),
		// 				},
		// 				ImprovedQueryIDs: []*int64{
		// 					to.Ptr[int64](2071439792137543700),
		// 					to.Ptr[int64](7860150533486302000),
		// 					to.Ptr[int64](6411979446509506000),
		// 					to.Ptr[int64](3219604056681277400),
		// 					to.Ptr[int64](-360410933364310600),
		// 					to.Ptr[int64](6171467644166225000),
		// 					to.Ptr[int64](3548728559597612500),
		// 					to.Ptr[int64](-4753875211349607000),
		// 					to.Ptr[int64](-8711548294430095000)},
		// 					InitialRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 					LastRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 					RecommendationReason: to.Ptr("Table \"nation\" in schema \"public\" is unanalyzed and appears in the following queries: 2071439792137543669, 7860150533486301820, 6411979446509505239, 3219604056681277471, -360410933364310591, 6171467644166224729, 3548728559597612316, -4753875211349607298, -8711548294430094920"),
		// 					RecommendationType: to.Ptr(armpostgresqlflexibleservers.RecommendationTypeEnum("Analyze")),
		// 					TimesRecommended: to.Ptr[int32](1),
		// 					Details: &armpostgresqlflexibleservers.ObjectRecommendationDetails{
		// 						Schema: to.Ptr("public"),
		// 						DatabaseName: to.Ptr("postgres"),
		// 						Table: to.Ptr("nation"),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Analyze_postgres_public_region"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/tuningOptions/table"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/tuningOptions/table/recommendations/2"),
		// 				Kind: to.Ptr(""),
		// 				Properties: &armpostgresqlflexibleservers.ObjectRecommendationProperties{
		// 					AnalyzedWorkload: &armpostgresqlflexibleservers.ObjectRecommendationPropertiesAnalyzedWorkload{
		// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 						QueryCount: to.Ptr[int32](22),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T19:30:22.123Z"); return t}()),
		// 					},
		// 					EstimatedImpact: []*armpostgresqlflexibleservers.ImpactRecord{
		// 					},
		// 					ImplementationDetails: &armpostgresqlflexibleservers.ObjectRecommendationPropertiesImplementationDetails{
		// 						Method: to.Ptr("SQL"),
		// 						Script: to.Ptr("analyze table public.region"),
		// 					},
		// 					ImprovedQueryIDs: []*int64{
		// 						to.Ptr[int64](3219604056681277400),
		// 						to.Ptr[int64](6171467644166225000),
		// 						to.Ptr[int64](-4753875211349607000)},
		// 						InitialRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 						LastRecommendedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-06-01T20:30:22.123Z"); return t}()),
		// 						RecommendationReason: to.Ptr("Table \"region\" in schema \"public\" is unanalyzed and appears in the following queries: 3219604056681277471, 6171467644166224729, -4753875211349607298"),
		// 						RecommendationType: to.Ptr(armpostgresqlflexibleservers.RecommendationTypeEnum("Analyze")),
		// 						TimesRecommended: to.Ptr[int32](1),
		// 						Details: &armpostgresqlflexibleservers.ObjectRecommendationDetails{
		// 							Schema: to.Ptr("public"),
		// 							DatabaseName: to.Ptr("postgres"),
		// 							Table: to.Ptr("region"),
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}

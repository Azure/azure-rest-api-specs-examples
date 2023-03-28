package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7932c2df6c8435d6c0e5cbebbca79bce627d5f06/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoritesList.json
func ExampleFavoritesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapplicationinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFavoritesClient().List(ctx, "my-resource-group", "my-ai-component", &armapplicationinsights.FavoritesClientListOptions{FavoriteType: nil,
		SourceType:      nil,
		CanFetchContent: nil,
		Tags:            []string{},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ComponentFavoriteArray = []*armapplicationinsights.ComponentFavorite{
	// 	{
	// 		Config: to.Ptr("blah blah"),
	// 		FavoriteID: to.Ptr("c0deea5e-3344-40f2-96f8-6f8e1c3b5722"),
	// 		FavoriteType: to.Ptr(armapplicationinsights.FavoriteTypeShared),
	// 		IsGeneratedFromTemplate: to.Ptr(false),
	// 		Name: to.Ptr("Example Metric Chart Favorite"),
	// 		Tags: []*string{
	// 		},
	// 		TimeModified: to.Ptr("2018-01-31T23:56:16.008902Z"),
	// 		Version: to.Ptr("ME"),
	// 	},
	// 	{
	// 		Config: to.Ptr("{\"TimeSelection\":{\"durationMs\":1800000,\"endTime\":\"2018-01-31T23:56:15.493Z\",\"createdTime\":\"Wed Jan 31 2018 15:58:36 GMT-0800 (Pacific Standard Time)\",\"isInitialTime\":false,\"grain\":1,\"useDashboardTimeRange\":false},\"SearchFilter\":{\"eventTypes\":[1,2],\"typeFacets\":{},\"isPermissive\":false},\"QueryText\":\"*\",\"partId\":\"99e33a16-1b00-4a7d-b98f-a3c1bb3a4df8\"}"),
	// 		FavoriteID: to.Ptr("deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
	// 		FavoriteType: to.Ptr(armapplicationinsights.FavoriteTypeShared),
	// 		IsGeneratedFromTemplate: to.Ptr(false),
	// 		Name: to.Ptr("Example Search Blade Favorite"),
	// 		Tags: []*string{
	// 			to.Ptr("SampleTag01"),
	// 			to.Ptr("SampleTag2")},
	// 			TimeModified: to.Ptr("2018-01-31T23:59:25.4594264Z"),
	// 			Version: to.Ptr("Search"),
	// 	}}
}

package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/transforms-list-all-filter-by-lastmodified.json
func ExampleTransformsClient_NewListPager_listsTheTransformsFilterByLastmodified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTransformsClient().NewListPager("contosoresources", "contosomedia", &armmediaservices.TransformsClientListOptions{Filter: to.Ptr("properties/lastmodified gt 2021-06-01T00:00:00.0000000Z and properties/lastmodified le 2021-06-01T00:00:10.0000000Z"),
		Orderby: to.Ptr("properties/lastmodified desc"),
	})
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
		// page.TransformCollection = armmediaservices.TransformCollection{
		// 	Value: []*armmediaservices.Transform{
		// 		{
		// 			Name: to.Ptr("sampleEncodeAndVideoIndex"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/transforms"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoresources/providers/Microsoft.Media/mediaservices/contosomedia/transforms/sampleEncodeAndVideoIndex"),
		// 			Properties: &armmediaservices.TransformProperties{
		// 				Description: to.Ptr("A sample Transform using the Video Analyzer."),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:10.000Z"); return t}()),
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:10.000Z"); return t}()),
		// 				Outputs: []*armmediaservices.TransformOutput{
		// 					{
		// 						OnError: to.Ptr(armmediaservices.OnErrorTypeStopProcessingJob),
		// 						Preset: &armmediaservices.VideoAnalyzerPreset{
		// 							ODataType: to.Ptr("#Microsoft.Media.VideoAnalyzerPreset"),
		// 							Mode: to.Ptr(armmediaservices.AudioAnalysisModeStandard),
		// 							InsightsToExtract: to.Ptr(armmediaservices.InsightsTypeAllInsights),
		// 						},
		// 						RelativePriority: to.Ptr(armmediaservices.PriorityNormal),
		// 				}},
		// 			},
		// 			SystemData: &armmediaservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:10.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("contoso@microsoft.com"),
		// 				CreatedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:10.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("contoso@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}

package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/jobs-list-all-filter-by-created.json
func ExampleJobsClient_NewListPager_listsJobsForTheTransformFilterByCreated() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewJobsClient().NewListPager("contosoresources", "contosomedia", "exampleTransform", &armmediaservices.JobsClientListOptions{Filter: to.Ptr("properties/created ge 2021-06-01T00:00:10.0000000Z and properties/created le 2021-06-01T00:00:20.0000000Z"),
		Orderby: to.Ptr("properties/created"),
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
		// page.JobCollection = armmediaservices.JobCollection{
		// 	Value: []*armmediaservices.Job{
		// 		{
		// 			Name: to.Ptr("job2"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/transforms/jobs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoresources/providers/Microsoft.Media/mediaservices/contosomedia/transforms/exampleTransform/jobs/job2"),
		// 			Properties: &armmediaservices.JobProperties{
		// 				CorrelationData: map[string]*string{
		// 				},
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:10.000Z"); return t}()),
		// 				Input: &armmediaservices.JobInputs{
		// 					ODataType: to.Ptr("#Microsoft.Media.JobInputs"),
		// 					Inputs: []armmediaservices.JobInputClassification{
		// 						&armmediaservices.JobInputAsset{
		// 							ODataType: to.Ptr("#Microsoft.Media.JobInputAsset"),
		// 							Files: []*string{
		// 							},
		// 							InputDefinitions: []armmediaservices.InputDefinitionClassification{
		// 							},
		// 							AssetName: to.Ptr("job2-InputAsset"),
		// 					}},
		// 				},
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:10.000Z"); return t}()),
		// 				Outputs: []armmediaservices.JobOutputClassification{
		// 					&armmediaservices.JobOutputAsset{
		// 						ODataType: to.Ptr("#Microsoft.Media.JobOutputAsset"),
		// 						Label: to.Ptr("example-custom-label"),
		// 						Progress: to.Ptr[int32](50),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-17T23:04:33.233Z"); return t}()),
		// 						State: to.Ptr(armmediaservices.JobStateProcessing),
		// 						AssetName: to.Ptr("job2-OutputAsset"),
		// 				}},
		// 				Priority: to.Ptr(armmediaservices.PriorityLow),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-17T23:04:33.233Z"); return t}()),
		// 				State: to.Ptr(armmediaservices.JobStateProcessing),
		// 			},
		// 			SystemData: &armmediaservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:10.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("contoso@microsoft.com"),
		// 				CreatedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:10.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("contoso@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("job3"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/transforms/jobs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoresources/providers/Microsoft.Media/mediaservices/contosomedia/transforms/exampleTransform/jobs/job3"),
		// 			Properties: &armmediaservices.JobProperties{
		// 				CorrelationData: map[string]*string{
		// 				},
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:20.000Z"); return t}()),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-17T23:14:33.233Z"); return t}()),
		// 				Input: &armmediaservices.JobInputs{
		// 					ODataType: to.Ptr("#Microsoft.Media.JobInputs"),
		// 					Inputs: []armmediaservices.JobInputClassification{
		// 						&armmediaservices.JobInputAsset{
		// 							ODataType: to.Ptr("#Microsoft.Media.JobInputAsset"),
		// 							Files: []*string{
		// 							},
		// 							InputDefinitions: []armmediaservices.InputDefinitionClassification{
		// 							},
		// 							AssetName: to.Ptr("job3-InputAsset"),
		// 					}},
		// 				},
		// 				LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:20.000Z"); return t}()),
		// 				Outputs: []armmediaservices.JobOutputClassification{
		// 					&armmediaservices.JobOutputAsset{
		// 						ODataType: to.Ptr("#Microsoft.Media.JobOutputAsset"),
		// 						EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-17T23:14:33.233Z"); return t}()),
		// 						Label: to.Ptr("example-custom-label"),
		// 						Progress: to.Ptr[int32](100),
		// 						StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-17T23:04:33.233Z"); return t}()),
		// 						State: to.Ptr(armmediaservices.JobStateFinished),
		// 						AssetName: to.Ptr("job3-OutputAsset"),
		// 				}},
		// 				Priority: to.Ptr(armmediaservices.PriorityLow),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-17T23:04:33.233Z"); return t}()),
		// 				State: to.Ptr(armmediaservices.JobStateFinished),
		// 			},
		// 			SystemData: &armmediaservices.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:20.000Z"); return t}()),
		// 				CreatedBy: to.Ptr("contoso@microsoft.com"),
		// 				CreatedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-01T00:00:20.000Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("contoso@microsoft.com"),
		// 				LastModifiedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
		// 			},
		// 	}},
		// }
	}
}

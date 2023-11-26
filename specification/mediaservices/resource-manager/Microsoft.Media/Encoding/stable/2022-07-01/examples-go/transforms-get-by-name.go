package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Encoding/stable/2022-07-01/examples/transforms-get-by-name.json
func ExampleTransformsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTransformsClient().Get(ctx, "contosoresources", "contosomedia", "sampleTransform", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Transform = armmediaservices.Transform{
	// 	Name: to.Ptr("sampleTransform"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/transforms"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contosoresources/providers/Microsoft.Media/mediaservices/contosomedia/transforms/sampleTransform"),
	// 	Properties: &armmediaservices.TransformProperties{
	// 		Description: to.Ptr("A sample Transform using the Standard Encoder with a built in preset."),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-17T23:14:31.628Z"); return t}()),
	// 		LastModified: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-17T23:14:31.628Z"); return t}()),
	// 		Outputs: []*armmediaservices.TransformOutput{
	// 			{
	// 				OnError: to.Ptr(armmediaservices.OnErrorTypeStopProcessingJob),
	// 				Preset: &armmediaservices.BuiltInStandardEncoderPreset{
	// 					ODataType: to.Ptr("#Microsoft.Media.BuiltInStandardEncoderPreset"),
	// 					PresetName: to.Ptr(armmediaservices.EncoderNamedPresetH264MultipleBitrate1080P),
	// 				},
	// 				RelativePriority: to.Ptr(armmediaservices.PriorityNormal),
	// 		}},
	// 	},
	// 	SystemData: &armmediaservices.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-17T23:14:31.628Z"); return t}()),
	// 		CreatedBy: to.Ptr("contoso@microsoft.com"),
	// 		CreatedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-17T23:14:31.628Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("contoso@microsoft.com"),
	// 		LastModifiedByType: to.Ptr(armmediaservices.CreatedByTypeUser),
	// 	},
	// }
}

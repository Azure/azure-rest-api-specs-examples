package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/transforms-create.json
func ExampleTransformsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewTransformsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "contosoresources", "contosomedia", "createdTransform", armmediaservices.Transform{
		Properties: &armmediaservices.TransformProperties{
			Description: to.Ptr("Example Transform to illustrate create and update."),
			Outputs: []*armmediaservices.TransformOutput{
				{
					Preset: &armmediaservices.BuiltInStandardEncoderPreset{
						ODataType:  to.Ptr("#Microsoft.Media.BuiltInStandardEncoderPreset"),
						PresetName: to.Ptr(armmediaservices.EncoderNamedPresetAdaptiveStreaming),
					},
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices"
)

// x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/transforms-create.json
func ExampleTransformsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmediaservices.NewTransformsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<transform-name>",
		armmediaservices.Transform{
			Properties: &armmediaservices.TransformProperties{
				Description: to.StringPtr("<description>"),
				Outputs: []*armmediaservices.TransformOutput{
					{
						Preset: &armmediaservices.BuiltInStandardEncoderPreset{
							ODataType:  to.StringPtr("<odata-type>"),
							PresetName: armmediaservices.EncoderNamedPreset("AdaptiveStreaming").ToPtr(),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.TransformsClientCreateOrUpdateResult)
}

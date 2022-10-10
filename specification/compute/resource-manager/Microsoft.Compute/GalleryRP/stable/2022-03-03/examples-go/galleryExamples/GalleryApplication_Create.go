package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/galleryExamples/GalleryApplication_Create.json
func ExampleGalleryApplicationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewGalleryApplicationsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryApplicationName", armcompute.GalleryApplication{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryApplicationProperties{
			Description: to.Ptr("This is the gallery application description."),
			CustomActions: []*armcompute.GalleryApplicationCustomAction{
				{
					Name:        to.Ptr("myCustomAction"),
					Description: to.Ptr("This is the custom action description."),
					Parameters: []*armcompute.GalleryApplicationCustomActionParameter{
						{
							Name:         to.Ptr("myCustomActionParameter"),
							Type:         to.Ptr(armcompute.GalleryApplicationCustomActionParameterTypeString),
							Description:  to.Ptr("This is the description of the parameter"),
							DefaultValue: to.Ptr("default value of parameter."),
							Required:     to.Ptr(false),
						}},
					Script: to.Ptr("myCustomActionScript"),
				}},
			Eula:                to.Ptr("This is the gallery application EULA."),
			PrivacyStatementURI: to.Ptr("myPrivacyStatementUri}"),
			ReleaseNoteURI:      to.Ptr("myReleaseNoteUri"),
			SupportedOSType:     to.Ptr(armcompute.OperatingSystemTypesWindows),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

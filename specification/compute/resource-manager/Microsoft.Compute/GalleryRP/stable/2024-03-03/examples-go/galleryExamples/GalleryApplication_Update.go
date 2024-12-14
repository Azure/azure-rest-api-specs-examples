package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/GalleryApplication_Update.json
func ExampleGalleryApplicationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryApplicationsClient().BeginUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryApplicationName", armcompute.GalleryApplicationUpdate{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GalleryApplication = armcompute.GalleryApplication{
	// 	Name: to.Ptr("myGalleryApplicationName"),
	// 	Type: to.Ptr("Microsoft.Compute/galleries"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/applications/myGalleryApplicationName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryApplicationProperties{
	// 		Description: to.Ptr("This is the gallery application description."),
	// 		CustomActions: []*armcompute.GalleryApplicationCustomAction{
	// 			{
	// 				Name: to.Ptr("myCustomAction"),
	// 				Description: to.Ptr("This is the custom action description."),
	// 				Parameters: []*armcompute.GalleryApplicationCustomActionParameter{
	// 					{
	// 						Name: to.Ptr("myCustomActionParameter"),
	// 						Type: to.Ptr(armcompute.GalleryApplicationCustomActionParameterTypeString),
	// 						Description: to.Ptr("This is the description of the parameter"),
	// 						DefaultValue: to.Ptr("default value of parameter."),
	// 						Required: to.Ptr(false),
	// 				}},
	// 				Script: to.Ptr("myCustomActionScript"),
	// 		}},
	// 		Eula: to.Ptr("This is the gallery application EULA."),
	// 		PrivacyStatementURI: to.Ptr("myPrivacyStatementUri}"),
	// 		ReleaseNoteURI: to.Ptr("myReleaseNoteUri"),
	// 		SupportedOSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 	},
	// }
}

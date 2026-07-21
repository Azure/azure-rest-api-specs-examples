package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-12-03/galleryExamples/GalleryApplication_Get.json
func ExampleGalleryApplicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleryApplicationsClient().Get(ctx, "myResourceGroup", "myGalleryName", "myGalleryApplicationName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.GalleryApplicationsClientGetResponse{
	// 	GalleryApplication: armcompute.GalleryApplication{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/applications/myGalleryApplicationName"),
	// 		Properties: &armcompute.GalleryApplicationProperties{
	// 			Description: to.Ptr("This is the gallery application description."),
	// 			Eula: to.Ptr("This is the gallery application EULA."),
	// 			PrivacyStatementURI: to.Ptr("myPrivacyStatementUri}"),
	// 			ReleaseNoteURI: to.Ptr("myReleaseNoteUri"),
	// 			SupportedOSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			CustomActions: []*armcompute.GalleryApplicationCustomAction{
	// 				{
	// 					Name: to.Ptr("myCustomAction"),
	// 					Script: to.Ptr("myCustomActionScript"),
	// 					Description: to.Ptr("This is the custom action description."),
	// 					Parameters: []*armcompute.GalleryApplicationCustomActionParameter{
	// 						{
	// 							Name: to.Ptr("myCustomActionParameter"),
	// 							Required: to.Ptr(false),
	// 							Type: to.Ptr(armcompute.GalleryApplicationCustomActionParameterTypeString),
	// 							DefaultValue: to.Ptr("default value of parameter."),
	// 							Description: to.Ptr("This is the description of the parameter"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("myGalleryApplicationName"),
	// 	},
	// }
}

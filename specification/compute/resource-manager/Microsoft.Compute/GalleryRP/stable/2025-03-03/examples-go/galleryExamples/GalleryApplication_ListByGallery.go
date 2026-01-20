package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/GalleryApplication_ListByGallery.json
func ExampleGalleryApplicationsClient_NewListByGalleryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGalleryApplicationsClient().NewListByGalleryPager("myResourceGroup", "myGalleryName", nil)
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
		// page.GalleryApplicationList = armcompute.GalleryApplicationList{
		// 	Value: []*armcompute.GalleryApplication{
		// 		{
		// 			Name: to.Ptr("myGalleryApplicationName"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/applications/myGalleryApplicationName"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcompute.GalleryApplicationProperties{
		// 				Description: to.Ptr("This is the gallery application description."),
		// 				CustomActions: []*armcompute.GalleryApplicationCustomAction{
		// 					{
		// 						Name: to.Ptr("myCustomAction"),
		// 						Description: to.Ptr("This is the custom action description."),
		// 						Parameters: []*armcompute.GalleryApplicationCustomActionParameter{
		// 							{
		// 								Name: to.Ptr("myCustomActionParameter"),
		// 								Type: to.Ptr(armcompute.GalleryApplicationCustomActionParameterTypeString),
		// 								Description: to.Ptr("This is the description of the parameter"),
		// 								DefaultValue: to.Ptr("default value of parameter."),
		// 								Required: to.Ptr(false),
		// 						}},
		// 						Script: to.Ptr("myCustomActionScript"),
		// 				}},
		// 				Eula: to.Ptr("This is the gallery application EULA."),
		// 				PrivacyStatementURI: to.Ptr("myPrivacyStatementUri}"),
		// 				ReleaseNoteURI: to.Ptr("myReleaseNoteUri"),
		// 				SupportedOSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 			},
		// 	}},
		// }
	}
}

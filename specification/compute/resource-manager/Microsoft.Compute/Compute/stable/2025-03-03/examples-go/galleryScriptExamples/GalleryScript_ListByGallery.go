package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryScriptExamples/GalleryScript_ListByGallery.json
func ExampleGalleryScriptsClient_NewListByGalleryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGalleryScriptsClient().NewListByGalleryPager("myResourceGroup", "myGalleryName", nil)
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
		// page = armcompute.GalleryScriptsClientListByGalleryResponse{
		// 	GalleryScriptList: armcompute.GalleryScriptList{
		// 		Value: []*armcompute.GalleryScript{
		// 			{
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/scripts/myGalleryScriptName"),
		// 				Properties: &armcompute.GalleryScriptProperties{
		// 					Description: to.Ptr("This is the gallery script description."),
		// 					Eula: to.Ptr("This is the gallery script EULA."),
		// 					PrivacyStatementURI: to.Ptr("{myPrivacyStatementUri}"),
		// 					ReleaseNoteURI: to.Ptr("{myReleaseNoteUri}"),
		// 					SupportedOSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 					ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 				},
		// 				Location: to.Ptr("West US"),
		// 				Name: to.Ptr("myGalleryScriptName"),
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/scripts/myGalleryScriptName2"),
		// 				Properties: &armcompute.GalleryScriptProperties{
		// 					Description: to.Ptr("This is the gallery script description."),
		// 					Eula: to.Ptr("This is the gallery script EULA."),
		// 					PrivacyStatementURI: to.Ptr("{myPrivacyStatementUri}"),
		// 					ReleaseNoteURI: to.Ptr("{myReleaseNoteUri}"),
		// 					SupportedOSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 					ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 				},
		// 				Location: to.Ptr("West US"),
		// 				Name: to.Ptr("myGalleryScriptName2"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/nextLinkExample"),
		// 	},
		// }
	}
}

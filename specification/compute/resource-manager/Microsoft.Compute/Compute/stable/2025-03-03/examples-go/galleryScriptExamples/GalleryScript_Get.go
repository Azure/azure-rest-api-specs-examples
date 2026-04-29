package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryScriptExamples/GalleryScript_Get.json
func ExampleGalleryScriptsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleryScriptsClient().Get(ctx, "myResourceGroup", "myGalleryName", "myGalleryScriptName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.GalleryScriptsClientGetResponse{
	// 	GalleryScript: &armcompute.GalleryScript{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/scripts/myGalleryScriptName"),
	// 		Properties: &armcompute.GalleryScriptProperties{
	// 			Description: to.Ptr("This is the gallery script description."),
	// 			Eula: to.Ptr("This is the gallery script EULA."),
	// 			PrivacyStatementURI: to.Ptr("{myPrivacyStatementUri}"),
	// 			ReleaseNoteURI: to.Ptr("{myReleaseNoteUri}"),
	// 			SupportedOSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("myGalleryScriptName"),
	// 	},
	// }
}

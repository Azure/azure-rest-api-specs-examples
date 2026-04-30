package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryResourceProfileExamples/GalleryInVMAccessControlProfile_ListByGallery.json
func ExampleGalleryInVMAccessControlProfilesClient_NewListByGalleryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGalleryInVMAccessControlProfilesClient().NewListByGalleryPager("myResourceGroup", "myGalleryName", nil)
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
		// page = armcompute.GalleryInVMAccessControlProfilesClientListByGalleryResponse{
		// 	GalleryInVMAccessControlProfileList: armcompute.GalleryInVMAccessControlProfileList{
		// 		Value: []*armcompute.GalleryInVMAccessControlProfile{
		// 			{
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGallery/inVMAccessControlProfiles/myInVMAccessControlProfileName"),
		// 				Properties: &armcompute.GalleryInVMAccessControlProfileProperties{
		// 					OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
		// 					ApplicableHostEndpoint: to.Ptr(armcompute.EndpointTypesWireServer),
		// 					ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 				},
		// 				Location: to.Ptr("West US"),
		// 				Name: to.Ptr("myInVMAccessControlProfileName"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://svchost:99/subscriptions/subscriptionId/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/inVMAccessControlProfiles?$skiptoken=token/Subscriptions/subscriptionId/ResourceGroups/myResourceGroup/galleries/myGalleryName/inVMAccessControlProfiles/myInVMAccessControlProfileName"),
		// 	},
		// }
	}
}

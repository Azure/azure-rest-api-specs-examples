package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryResourceProfileExamples/GalleryInVMAccessControlProfile_ListByGallery.json
func ExampleGalleryInVMAccessControlProfilesClient_NewListByGalleryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.GalleryInVMAccessControlProfileList = armcompute.GalleryInVMAccessControlProfileList{
		// 	Value: []*armcompute.GalleryInVMAccessControlProfile{
		// 		{
		// 			Name: to.Ptr("myInVMAccessControlProfileName"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGallery/inVMAccessControlProfiles/myInVMAccessControlProfileName"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcompute.GalleryInVMAccessControlProfileProperties{
		// 				ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 				ApplicableHostEndpoint: to.Ptr(armcompute.EndpointTypesWireServer),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
		// 			},
		// 	}},
		// }
	}
}

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryResourceProfileExamples/GalleryInVMAccessControlProfile_Get.json
func ExampleGalleryInVMAccessControlProfilesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleryInVMAccessControlProfilesClient().Get(ctx, "myResourceGroup", "myGalleryName", "myInVMAccessControlProfileName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GalleryInVMAccessControlProfile = armcompute.GalleryInVMAccessControlProfile{
	// 	Name: to.Ptr("myInVMAccessControlProfileName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGallery/inVMAccessControlProfiles/myInVMAccessControlProfileName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryInVMAccessControlProfileProperties{
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		ApplicableHostEndpoint: to.Ptr(armcompute.EndpointTypesWireServer),
	// 		OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
	// 	},
	// }
}

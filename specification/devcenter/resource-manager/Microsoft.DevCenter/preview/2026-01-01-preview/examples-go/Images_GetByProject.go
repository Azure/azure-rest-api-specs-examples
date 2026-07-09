package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/Images_GetByProject.json
func ExampleImagesClient_GetByProject() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewImagesClient().GetByProject(ctx, "rg1", "myProject", "~gallery~DefaultDevGallery~ContosoBaseImage", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.ImagesClientGetByProjectResponse{
	// 	Image: armdevcenter.Image{
	// 		Name: to.Ptr("ContosoBaseImage"),
	// 		Type: to.Ptr("Microsoft.DevCenter/project/images"),
	// 		ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/project/myProject/images/~gallery~DefaultDevGallery~ContosoBaseImage"),
	// 		Properties: &armdevcenter.ImageProperties{
	// 			Architecture: to.Ptr(armdevcenter.ArchitectureTypeX64),
	// 			Description: to.Ptr("Standard Windows Dev/Test image."),
	// 			Offer: to.Ptr("Finance"),
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 			Publisher: to.Ptr("Contoso"),
	// 			RecommendedMachineConfiguration: &armdevcenter.RecommendedMachineConfiguration{
	// 				Memory: &armdevcenter.ResourceRange{
	// 					Max: to.Ptr[int32](512),
	// 					Min: to.Ptr[int32](256),
	// 				},
	// 				VCPUs: &armdevcenter.ResourceRange{
	// 					Max: to.Ptr[int32](8),
	// 					Min: to.Ptr[int32](4),
	// 				},
	// 			},
	// 			SKU: to.Ptr("Backend"),
	// 		},
	// 		SystemData: &armdevcenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user1"),
	// 			LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}

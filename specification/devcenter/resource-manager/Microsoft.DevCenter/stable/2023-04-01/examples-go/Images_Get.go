package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Images_Get.json
func ExampleImagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewImagesClient().Get(ctx, "rg1", "Contoso", "DefaultDevGallery", "ContosoBaseImage", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Image = armdevcenter.Image{
	// 	Name: to.Ptr("ContosoBaseImage"),
	// 	Type: to.Ptr("Microsoft.DevCenter/devcenters/galleries/images"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso/galleries/DefaultDevGallery/images/ContosoBaseImage"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user1"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Properties: &armdevcenter.ImageProperties{
	// 		Description: to.Ptr("Standard Windows Dev/Test image."),
	// 		Offer: to.Ptr("Finance"),
	// 		ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 		Publisher: to.Ptr("Contoso"),
	// 		RecommendedMachineConfiguration: &armdevcenter.RecommendedMachineConfiguration{
	// 			Memory: &armdevcenter.ResourceRange{
	// 				Max: to.Ptr[int32](512),
	// 				Min: to.Ptr[int32](256),
	// 			},
	// 			VCPUs: &armdevcenter.ResourceRange{
	// 				Max: to.Ptr[int32](8),
	// 				Min: to.Ptr[int32](4),
	// 			},
	// 		},
	// 		SKU: to.Ptr("Backend"),
	// 	},
	// }
}

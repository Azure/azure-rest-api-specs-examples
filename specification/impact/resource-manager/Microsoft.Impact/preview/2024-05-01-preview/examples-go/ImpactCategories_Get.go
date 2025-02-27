package armimpactreporting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
)

// Generated from example definition: 2024-05-01-preview/ImpactCategories_Get.json
func ExampleImpactCategoriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewImpactCategoriesClient().Get(ctx, "ARMOperation.Create", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armimpactreporting.ImpactCategoriesClientGetResponse{
	// 	ImpactCategory: &armimpactreporting.ImpactCategory{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/ImpactCategories/ARMOperation.Create"),
	// 		Name: to.Ptr("ARMOperation.Create"),
	// 		Type: to.Ptr("Microsoft.Impact/impactCategories"),
	// 		Properties: &armimpactreporting.ImpactCategoryProperties{
	// 			Description: to.Ptr("Use this to report problems related to creating a new azure virtual machine such as provisioning or allocation failures."),
	// 			CategoryID: to.Ptr("778f815b-6576-4a36-bea9-bffb3d26d7f4"),
	// 			ParentCategoryID: to.Ptr("36266d25-9c53-40b3-af41-27418b11851e"),
	// 			RequiredImpactProperties: []*armimpactreporting.RequiredImpactProperties{
	// 				{
	// 					Name: to.Ptr("armCorrelations"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}

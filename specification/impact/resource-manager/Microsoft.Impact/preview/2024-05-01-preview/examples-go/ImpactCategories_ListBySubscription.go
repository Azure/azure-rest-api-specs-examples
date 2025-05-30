package armimpactreporting_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/impactreporting/armimpactreporting"
)

// Generated from example definition: 2024-05-01-preview/ImpactCategories_ListBySubscription.json
func ExampleImpactCategoriesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armimpactreporting.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewImpactCategoriesClient().NewListBySubscriptionPager("microsoft.compute/virtualmachines", nil)
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
		// page = armimpactreporting.ImpactCategoriesClientListBySubscriptionResponse{
		// 	ImpactCategoryListResult: armimpactreporting.ImpactCategoryListResult{
		// 		Value: []*armimpactreporting.ImpactCategory{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/ImpactCategories/ARMOperation.Create"),
		// 				Name: to.Ptr("ARMOperation.Create"),
		// 				Type: to.Ptr("Microsoft.Impact/impactCategories"),
		// 				Properties: &armimpactreporting.ImpactCategoryProperties{
		// 					Description: to.Ptr("Use this to report problems related to creating a new azure virtual machine such as provisioning or allocation failures."),
		// 					CategoryID: to.Ptr("778f815b-6576-4a36-bea9-bffb3d26d7f4"),
		// 					ParentCategoryID: to.Ptr("36266d25-9c53-40b3-af41-27418b11851e"),
		// 					RequiredImpactProperties: []*armimpactreporting.RequiredImpactProperties{
		// 						{
		// 							Name: to.Ptr("armCorrelations"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/ImpactCategories/ARMOperation.Delete"),
		// 				Name: to.Ptr("ARMOperation.Delete"),
		// 				Type: to.Ptr("Microsoft.Impact/impactCategories"),
		// 				Properties: &armimpactreporting.ImpactCategoryProperties{
		// 					Description: to.Ptr("Use this to report failures in deleting VMs."),
		// 					CategoryID: to.Ptr("a9a0c6e2-1208-406f-8f4f-1ccc13fb75d5"),
		// 					ParentCategoryID: to.Ptr("36266d25-9c53-40b3-af41-27418b11851e"),
		// 					RequiredImpactProperties: []*armimpactreporting.RequiredImpactProperties{
		// 						{
		// 							Name: to.Ptr("armCorrelations"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/ImpactCategories/ARMOperation.Other"),
		// 				Name: to.Ptr("ARMOperation.Other"),
		// 				Type: to.Ptr("Microsoft.Impact/impactCategories"),
		// 				Properties: &armimpactreporting.ImpactCategoryProperties{
		// 					Description: to.Ptr("Use this to report Control Plane operation failures that don't fall into other ARMOperation categories"),
		// 					CategoryID: to.Ptr("7b71f937-9344-499c-bfbe-d86fb755d891"),
		// 					ParentCategoryID: to.Ptr("36266d25-9c53-40b3-af41-27418b11851e"),
		// 					RequiredImpactProperties: []*armimpactreporting.RequiredImpactProperties{
		// 						{
		// 							Name: to.Ptr("armCorrelations"),
		// 						},
		// 					},
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/ImpactCategories/Resource.Connectivity"),
		// 				Name: to.Ptr("Resource.Connectivity"),
		// 				Type: to.Ptr("Microsoft.Impact/impactCategories"),
		// 				Properties: &armimpactreporting.ImpactCategoryProperties{
		// 					Description: to.Ptr("Use this to report connectivity issues to or from a VM."),
		// 					CategoryID: to.Ptr("95903644-3a15-4e37-b4be-a2ae8651f27b"),
		// 					ParentCategoryID: to.Ptr("a8d1c1ae-5fcb-4a8b-a647-fd0a911e8667"),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/microsoft.Impact/ImpactCategories/Resource.Connectivity.Inbound"),
		// 				Name: to.Ptr("Resource.Connectivity.Inbound"),
		// 				Type: to.Ptr("Microsoft.Impact/impactCategories"),
		// 				Properties: &armimpactreporting.ImpactCategoryProperties{
		// 					Description: to.Ptr("Use this to report inbound connectivity issues to a VM."),
		// 					CategoryID: to.Ptr("940fb706-8586-4a0e-bb18-2e2d0ef0708d"),
		// 					ParentCategoryID: to.Ptr("95903644-3a15-4e37-b4be-a2ae8651f27b"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

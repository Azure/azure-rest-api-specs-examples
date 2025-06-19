package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/AppServiceEnvironments_ListMultiRolePoolSkus.json
func ExampleEnvironmentsClient_NewListMultiRolePoolSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnvironmentsClient().NewListMultiRolePoolSKUsPager("test-rg", "test-ase", nil)
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
		// page.SKUInfoCollection = armappservice.SKUInfoCollection{
		// 	Value: []*armappservice.SKUInfo{
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](247),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/multiRolePools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("S2"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](247),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/multiRolePools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("S3"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](247),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/multiRolePools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("S4"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](247),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/multiRolePools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("O2"),
		// 				Tier: to.Ptr("Optimized"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](247),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/multiRolePools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("O3"),
		// 				Tier: to.Ptr("Optimized"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](247),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/multiRolePools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("O4"),
		// 				Tier: to.Ptr("Optimized"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](247),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/multiRolePools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("Q1"),
		// 				Tier: to.Ptr("Quantum"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](247),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/multiRolePools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("Q2"),
		// 				Tier: to.Ptr("Quantum"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](247),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/multiRolePools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("Q3"),
		// 				Tier: to.Ptr("Quantum"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](247),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/multiRolePools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("Q4"),
		// 				Tier: to.Ptr("Quantum"),
		// 			},
		// 	}},
		// }
	}
}

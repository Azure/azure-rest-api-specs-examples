package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/AppServiceEnvironments_ListWorkerPoolSkus.json
func ExampleEnvironmentsClient_NewListWorkerPoolSKUsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnvironmentsClient().NewListWorkerPoolSKUsPager("test-rg", "test-ase", "workerPool1", nil)
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
		// 				Maximum: to.Ptr[int32](53),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/workerPools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("S1"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](53),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/workerPools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("S2"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](53),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/workerPools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("S3"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](53),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/workerPools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("S4"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Capacity: &armappservice.SKUCapacity{
		// 				Default: to.Ptr[int32](2),
		// 				Maximum: to.Ptr[int32](53),
		// 				Minimum: to.Ptr[int32](2),
		// 				ScaleType: to.Ptr("automatic"),
		// 			},
		// 			ResourceType: to.Ptr("Microsoft.Web/hostingEnvironments/workerPools"),
		// 			SKU: &armappservice.SKUDescription{
		// 				Name: to.Ptr("O1"),
		// 				Tier: to.Ptr("Optimized"),
		// 			},
		// 	}},
		// }
	}
}

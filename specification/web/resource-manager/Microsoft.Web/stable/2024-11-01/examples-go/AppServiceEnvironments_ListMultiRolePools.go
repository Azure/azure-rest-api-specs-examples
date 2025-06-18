package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/82e9c6f9fbfa2d6d47d5e2a6a11c0ad2eb345c43/specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/AppServiceEnvironments_ListMultiRolePools.json
func ExampleEnvironmentsClient_NewListMultiRolePoolsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnvironmentsClient().NewListMultiRolePoolsPager("test-rg", "test-ase", nil)
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
		// page.WorkerPoolCollection = armappservice.WorkerPoolCollection{
		// 	Value: []*armappservice.WorkerPoolResource{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Web/hostingEnvironments/multiRolePools"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/multiRolePools/default"),
		// 			Properties: &armappservice.WorkerPool{
		// 				InstanceNames: []*string{
		// 					to.Ptr("10.7.1.8"),
		// 					to.Ptr("10.7.1.9")},
		// 					WorkerCount: to.Ptr[int32](2),
		// 					WorkerSize: to.Ptr("Standard_D1_V2"),
		// 				},
		// 				SKU: &armappservice.SKUDescription{
		// 					Name: to.Ptr("Q1"),
		// 					Capacity: to.Ptr[int32](2),
		// 					Family: to.Ptr("Q"),
		// 					Size: to.Ptr("Q1"),
		// 					Tier: to.Ptr("Quantum"),
		// 				},
		// 		}},
		// 	}
	}
}

package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/21c2852d62ccc3abe9cc3800c989c6826f8363dc/specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/AppServiceEnvironments_ListWorkerPools.json
func ExampleEnvironmentsClient_NewListWorkerPoolsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnvironmentsClient().NewListWorkerPoolsPager("test-rg", "test-ase", nil)
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
		// 			Name: to.Ptr("workerPool1"),
		// 			Type: to.Ptr("Microsoft.Web/hostingEnvironments/workerPools"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/workerPools/workerPool1"),
		// 			Kind: to.Ptr("ASEV1"),
		// 			Properties: &armappservice.WorkerPool{
		// 				InstanceNames: []*string{
		// 				},
		// 				WorkerCount: to.Ptr[int32](2),
		// 				WorkerSize: to.Ptr("Small"),
		// 				WorkerSizeID: to.Ptr[int32](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("workerPool2"),
		// 			Type: to.Ptr("Microsoft.Web/hostingEnvironments/workerPools"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/workerPools/workerPool2"),
		// 			Kind: to.Ptr("ASEV1"),
		// 			Properties: &armappservice.WorkerPool{
		// 				InstanceNames: []*string{
		// 				},
		// 				WorkerCount: to.Ptr[int32](0),
		// 				WorkerSize: to.Ptr("Small"),
		// 				WorkerSizeID: to.Ptr[int32](1),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("workerPool3"),
		// 			Type: to.Ptr("Microsoft.Web/hostingEnvironments/workerPools"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/workerPools/workerPool3"),
		// 			Kind: to.Ptr("ASEV1"),
		// 			Properties: &armappservice.WorkerPool{
		// 				InstanceNames: []*string{
		// 				},
		// 				WorkerCount: to.Ptr[int32](0),
		// 				WorkerSize: to.Ptr("Small"),
		// 				WorkerSizeID: to.Ptr[int32](2),
		// 			},
		// 	}},
		// }
	}
}

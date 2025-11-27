package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/AppServiceEnvironments_GetWorkerPool.json
func ExampleEnvironmentsClient_GetWorkerPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEnvironmentsClient().GetWorkerPool(ctx, "test-rg", "test-ase", "workerPool1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkerPoolResource = armappservice.WorkerPoolResource{
	// 	Name: to.Ptr("workerPool1"),
	// 	Type: to.Ptr("Microsoft.Web/hostingEnvironments/workerPools"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-rg/providers/Microsoft.Web/hostingEnvironments/test-ase/workerPools/workerPool1"),
	// 	Kind: to.Ptr("ASEV1"),
	// 	Properties: &armappservice.WorkerPool{
	// 		InstanceNames: []*string{
	// 		},
	// 		WorkerCount: to.Ptr[int32](2),
	// 		WorkerSize: to.Ptr("Small"),
	// 		WorkerSizeID: to.Ptr[int32](0),
	// 	},
	// }
}

package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1dd99306d14fd6c602f47652a209a4a6812c368c/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-02-02-preview/examples/OperationStatusResultGet.json
func ExampleOperationStatusResultClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusResultClient().Get(ctx, "rg1", "clustername1", "00000000-0000-0000-0000-000000000001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatusResult = armcontainerservice.OperationStatusResult{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/operations/00000000-0000-0000-0000-000000000001"),
	// 	PercentComplete: to.Ptr[float32](40),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-26T12:14:26.317Z"); return t}()),
	// 	Status: to.Ptr("InProgress"),
	// }
}

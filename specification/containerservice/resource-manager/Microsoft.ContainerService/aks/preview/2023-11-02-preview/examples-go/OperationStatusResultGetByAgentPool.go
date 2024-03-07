package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-11-02-preview/examples/OperationStatusResultGetByAgentPool.json
func ExampleOperationStatusResultClient_GetByAgentPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusResultClient().GetByAgentPool(ctx, "rg1", "clustername1", "agentpool1", "00000000-0000-0000-0000-000000000001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatusResult = armcontainerservice.OperationStatusResult{
	// 	Name: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/agentpool1/operations/00000000-0000-0000-0000-000000000001"),
	// 	PercentComplete: to.Ptr[float32](40),
	// 	StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-26T12:14:26.317Z"); return t}()),
	// 	Status: to.Ptr("InProgress"),
	// }
}

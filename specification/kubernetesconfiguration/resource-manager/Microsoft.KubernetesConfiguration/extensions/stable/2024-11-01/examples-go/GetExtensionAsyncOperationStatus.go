package armextensions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armextensions"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ba0c086df0ebe03a61579485c1c10de0d17804b2/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/extensions/stable/2024-11-01/examples/GetExtensionAsyncOperationStatus.json
func ExampleOperationStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextensions.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationStatusClient().Get(ctx, "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "ClusterMonitor", "99999999-9999-9999-9999-999999999999", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OperationStatusResult = armextensions.OperationStatusResult{
	// 	Name: to.Ptr("99999999-9999-9999-9999-999999999999"),
	// 	ID: to.Ptr("/subscriptions/subId1/resourceGroups/rg1/providers/Microsoft.Kubernetes/connectedClusters/clusterName1/providers/Microsoft.KubernetesConfiguration/extensions/ClusterMonitor/operations/99999999-9999-9999-9999-999999999999"),
	// 	Properties: map[string]*string{
	// 	},
	// 	Status: to.Ptr("Succeeded"),
	// }
}

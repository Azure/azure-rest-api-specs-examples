package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/GetFluxConfigurationAsyncOperationStatus.json
func ExampleFluxConfigOperationStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkubernetesconfiguration.NewFluxConfigOperationStatusClient("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"rg1",
		"Microsoft.Kubernetes",
		"connectedClusters",
		"clusterName1",
		"srs-fluxconfig",
		"99999999-9999-9999-9999-999999999999",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

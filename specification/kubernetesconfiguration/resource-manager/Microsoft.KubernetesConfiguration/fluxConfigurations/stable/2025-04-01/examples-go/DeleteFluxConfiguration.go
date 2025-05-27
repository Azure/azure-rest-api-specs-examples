package armfluxconfigurations_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armfluxconfigurations"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ba0c086df0ebe03a61579485c1c10de0d17804b2/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/fluxConfigurations/stable/2025-04-01/examples/DeleteFluxConfiguration.json
func ExampleClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armfluxconfigurations.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClient().BeginDelete(ctx, "rg1", "Microsoft.Kubernetes", "connectedClusters", "clusterName1", "srs-fluxconfig", &armfluxconfigurations.ClientBeginDeleteOptions{ForceDelete: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

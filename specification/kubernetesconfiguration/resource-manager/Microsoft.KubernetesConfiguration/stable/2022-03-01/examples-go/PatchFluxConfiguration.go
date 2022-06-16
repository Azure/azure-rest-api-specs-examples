package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/PatchFluxConfiguration.json
func ExampleFluxConfigurationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkubernetesconfiguration.NewFluxConfigurationsClient("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"rg1",
		"Microsoft.Kubernetes",
		"connectedClusters",
		"clusterName1",
		"srs-fluxconfig",
		armkubernetesconfiguration.FluxConfigurationPatch{
			Properties: &armkubernetesconfiguration.FluxConfigurationPatchProperties{
				GitRepository: &armkubernetesconfiguration.GitRepositoryPatchDefinition{
					URL: to.Ptr("https://github.com/jonathan-innis/flux2-kustomize-helm-example.git"),
				},
				Kustomizations: map[string]*armkubernetesconfiguration.KustomizationPatchDefinition{
					"srs-kustomization1": nil,
					"srs-kustomization2": {
						Path:                  to.Ptr("./test/alt-path"),
						SyncIntervalInSeconds: to.Ptr[int64](300),
					},
					"srs-kustomization3": {
						Path:                  to.Ptr("./test/another-path"),
						SyncIntervalInSeconds: to.Ptr[int64](300),
					},
				},
				Suspend: to.Ptr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

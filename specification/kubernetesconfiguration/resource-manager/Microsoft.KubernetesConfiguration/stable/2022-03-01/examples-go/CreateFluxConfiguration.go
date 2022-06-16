package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/CreateFluxConfiguration.json
func ExampleFluxConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkubernetesconfiguration.NewFluxConfigurationsClient("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"Microsoft.Kubernetes",
		"connectedClusters",
		"clusterName1",
		"srs-fluxconfig",
		armkubernetesconfiguration.FluxConfiguration{
			Properties: &armkubernetesconfiguration.FluxConfigurationProperties{
				GitRepository: &armkubernetesconfiguration.GitRepositoryDefinition{
					HTTPSCACert: to.Ptr("ZXhhbXBsZWNlcnRpZmljYXRl"),
					RepositoryRef: &armkubernetesconfiguration.RepositoryRefDefinition{
						Branch: to.Ptr("master"),
					},
					SyncIntervalInSeconds: to.Ptr[int64](600),
					TimeoutInSeconds:      to.Ptr[int64](600),
					URL:                   to.Ptr("https://github.com/Azure/arc-k8s-demo"),
				},
				Kustomizations: map[string]*armkubernetesconfiguration.KustomizationDefinition{
					"srs-kustomization1": {
						Path:                  to.Ptr("./test/path"),
						DependsOn:             []*string{},
						SyncIntervalInSeconds: to.Ptr[int64](600),
						TimeoutInSeconds:      to.Ptr[int64](600),
					},
					"srs-kustomization2": {
						Path: to.Ptr("./other/test/path"),
						DependsOn: []*string{
							to.Ptr("srs-kustomization1")},
						Prune:                  to.Ptr(false),
						RetryIntervalInSeconds: to.Ptr[int64](600),
						SyncIntervalInSeconds:  to.Ptr[int64](600),
						TimeoutInSeconds:       to.Ptr[int64](600),
					},
				},
				Namespace:  to.Ptr("srs-namespace"),
				Scope:      to.Ptr(armkubernetesconfiguration.ScopeTypeCluster),
				SourceKind: to.Ptr(armkubernetesconfiguration.SourceKindTypeGitRepository),
				Suspend:    to.Ptr(false),
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

package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/preview/2022-01-01-preview/examples/CreateFluxConfiguration.json
func ExampleFluxConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkubernetesconfiguration.NewFluxConfigurationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		armkubernetesconfiguration.ExtensionsClusterRp("Microsoft.Kubernetes"),
		armkubernetesconfiguration.ExtensionsClusterResourceName("connectedClusters"),
		"<cluster-name>",
		"<flux-configuration-name>",
		armkubernetesconfiguration.FluxConfiguration{
			Properties: &armkubernetesconfiguration.FluxConfigurationProperties{
				GitRepository: &armkubernetesconfiguration.GitRepositoryDefinition{
					HTTPSCACert: to.StringPtr("<httpscacert>"),
					RepositoryRef: &armkubernetesconfiguration.RepositoryRefDefinition{
						Branch: to.StringPtr("<branch>"),
					},
					SyncIntervalInSeconds: to.Int64Ptr(600),
					TimeoutInSeconds:      to.Int64Ptr(600),
					URL:                   to.StringPtr("<url>"),
				},
				Kustomizations: map[string]*armkubernetesconfiguration.KustomizationDefinition{
					"srs-kustomization1": {
						Path:                  to.StringPtr("<path>"),
						DependsOn:             []*armkubernetesconfiguration.DependsOnDefinition{},
						SyncIntervalInSeconds: to.Int64Ptr(600),
						TimeoutInSeconds:      to.Int64Ptr(600),
					},
					"srs-kustomization2": {
						Path: to.StringPtr("<path>"),
						DependsOn: []*armkubernetesconfiguration.DependsOnDefinition{
							{
								KustomizationName: to.StringPtr("<kustomization-name>"),
							}},
						Prune:                  to.BoolPtr(false),
						RetryIntervalInSeconds: to.Int64Ptr(600),
						SyncIntervalInSeconds:  to.Int64Ptr(600),
						TimeoutInSeconds:       to.Int64Ptr(600),
					},
				},
				Namespace:  to.StringPtr("<namespace>"),
				Scope:      armkubernetesconfiguration.ScopeType("cluster").ToPtr(),
				SourceKind: armkubernetesconfiguration.SourceKindType("GitRepository").ToPtr(),
				Suspend:    to.BoolPtr(false),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FluxConfigurationsClientCreateOrUpdateResult)
}

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkubernetesconfiguration%2Farmkubernetesconfiguration%2Fv0.5.0/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.

```go
package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/CreateFluxConfiguration.json
func ExampleFluxConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armkubernetesconfiguration.NewFluxConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-rp>",
		"<cluster-resource-name>",
		"<cluster-name>",
		"<flux-configuration-name>",
		armkubernetesconfiguration.FluxConfiguration{
			Properties: &armkubernetesconfiguration.FluxConfigurationProperties{
				GitRepository: &armkubernetesconfiguration.GitRepositoryDefinition{
					HTTPSCACert: to.Ptr("<httpscacert>"),
					RepositoryRef: &armkubernetesconfiguration.RepositoryRefDefinition{
						Branch: to.Ptr("<branch>"),
					},
					SyncIntervalInSeconds: to.Ptr[int64](600),
					TimeoutInSeconds:      to.Ptr[int64](600),
					URL:                   to.Ptr("<url>"),
				},
				Kustomizations: map[string]*armkubernetesconfiguration.KustomizationDefinition{
					"srs-kustomization1": {
						Path:                  to.Ptr("<path>"),
						DependsOn:             []*string{},
						SyncIntervalInSeconds: to.Ptr[int64](600),
						TimeoutInSeconds:      to.Ptr[int64](600),
					},
					"srs-kustomization2": {
						Path: to.Ptr("<path>"),
						DependsOn: []*string{
							to.Ptr("srs-kustomization1")},
						Prune:                  to.Ptr(false),
						RetryIntervalInSeconds: to.Ptr[int64](600),
						SyncIntervalInSeconds:  to.Ptr[int64](600),
						TimeoutInSeconds:       to.Ptr[int64](600),
					},
				},
				Namespace:  to.Ptr("<namespace>"),
				Scope:      to.Ptr(armkubernetesconfiguration.ScopeTypeCluster),
				SourceKind: to.Ptr(armkubernetesconfiguration.SourceKindTypeGitRepository),
				Suspend:    to.Ptr(false),
			},
		},
		&armkubernetesconfiguration.FluxConfigurationsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

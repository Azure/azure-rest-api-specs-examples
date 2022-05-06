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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/PatchFluxConfiguration.json
func ExampleFluxConfigurationsClient_BeginUpdate() {
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
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-rp>",
		"<cluster-resource-name>",
		"<cluster-name>",
		"<flux-configuration-name>",
		armkubernetesconfiguration.FluxConfigurationPatch{
			Properties: &armkubernetesconfiguration.FluxConfigurationPatchProperties{
				GitRepository: &armkubernetesconfiguration.GitRepositoryPatchDefinition{
					URL: to.Ptr("<url>"),
				},
				Kustomizations: map[string]*armkubernetesconfiguration.KustomizationPatchDefinition{
					"srs-kustomization1": nil,
					"srs-kustomization2": {
						Path:                  to.Ptr("<path>"),
						SyncIntervalInSeconds: to.Ptr[int64](300),
					},
					"srs-kustomization3": {
						Path:                  to.Ptr("<path>"),
						SyncIntervalInSeconds: to.Ptr[int64](300),
					},
				},
				Suspend: to.Ptr(true),
			},
		},
		&armkubernetesconfiguration.FluxConfigurationsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```

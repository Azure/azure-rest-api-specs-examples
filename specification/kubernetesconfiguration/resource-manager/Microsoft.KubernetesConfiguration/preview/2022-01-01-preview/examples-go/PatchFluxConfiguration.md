Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkubernetesconfiguration%2Farmkubernetesconfiguration%2Fv0.2.1/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/preview/2022-01-01-preview/examples/PatchFluxConfiguration.json
func ExampleFluxConfigurationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkubernetesconfiguration.NewFluxConfigurationsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		armkubernetesconfiguration.Enum0("Microsoft.Kubernetes"),
		armkubernetesconfiguration.Enum1("connectedClusters"),
		"<cluster-name>",
		"<flux-configuration-name>",
		armkubernetesconfiguration.FluxConfigurationPatch{
			Properties: &armkubernetesconfiguration.FluxConfigurationPatchProperties{
				GitRepository: &armkubernetesconfiguration.GitRepositoryPatchDefinition{
					URL: to.StringPtr("<url>"),
				},
				Kustomizations: map[string]*armkubernetesconfiguration.KustomizationPatchDefinition{
					"srs-kustomization1": nil,
					"srs-kustomization2": {
						Path:                  to.StringPtr("<path>"),
						SyncIntervalInSeconds: to.Int64Ptr(300),
					},
					"srs-kustomization3": {
						Path:                  to.StringPtr("<path>"),
						SyncIntervalInSeconds: to.Int64Ptr(300),
					},
				},
				Suspend: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkubernetesconfiguration%2Farmkubernetesconfiguration%2Fv0.1.0/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.

```go
package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/preview/2022-01-01-preview/examples/GetFluxConfigurationAsyncOperationStatus.json
func ExampleFluxConfigOperationStatusClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkubernetesconfiguration.NewFluxConfigOperationStatusClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		armkubernetesconfiguration.Enum0MicrosoftKubernetes,
		armkubernetesconfiguration.Enum1ConnectedClusters,
		"<cluster-name>",
		"<flux-configuration-name>",
		"<operation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OperationStatusResult.ID: %s\n", *res.ID)
}
```

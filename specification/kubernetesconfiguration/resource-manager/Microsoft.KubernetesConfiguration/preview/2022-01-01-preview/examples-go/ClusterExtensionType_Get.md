Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkubernetesconfiguration%2Farmkubernetesconfiguration%2Fv0.1.0/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.

```go
package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/preview/2022-01-01-preview/examples/ClusterExtensionType_Get.json
func ExampleClusterExtensionTypeClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkubernetesconfiguration.NewClusterExtensionTypeClient("<subscription-id>", cred, nil)
	_, err = client.Get(ctx,
		"<resource-group-name>",
		armkubernetesconfiguration.Enum0MicrosoftContainerService,
		armkubernetesconfiguration.Enum1ManagedClusters,
		"<cluster-name>",
		"<extension-type-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/PatchExtension.json
func ExampleExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armkubernetesconfiguration.NewExtensionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<cluster-rp>",
		"<cluster-resource-name>",
		"<cluster-name>",
		"<extension-name>",
		armkubernetesconfiguration.PatchExtension{
			Properties: &armkubernetesconfiguration.PatchExtensionProperties{
				AutoUpgradeMinorVersion: to.Ptr(true),
				ConfigurationProtectedSettings: map[string]*string{
					"omsagent.secret.key": to.Ptr("secretKeyValue01"),
				},
				ConfigurationSettings: map[string]*string{
					"omsagent.env.clusterName": to.Ptr("clusterName1"),
					"omsagent.secret.wsid":     to.Ptr("a38cef99-5a89-52ed-b6db-22095c23664b"),
				},
				ReleaseTrain: to.Ptr("<release-train>"),
			},
		},
		&armkubernetesconfiguration.ExtensionsClientBeginUpdateOptions{ResumeToken: ""})
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

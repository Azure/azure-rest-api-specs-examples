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

// x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/preview/2022-01-01-preview/examples/CreateExtension.json
func ExampleExtensionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkubernetesconfiguration.NewExtensionsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		armkubernetesconfiguration.Enum0("Microsoft.Kubernetes"),
		armkubernetesconfiguration.Enum1("connectedClusters"),
		"<cluster-name>",
		"<extension-name>",
		armkubernetesconfiguration.Extension{
			Properties: &armkubernetesconfiguration.ExtensionProperties{
				AutoUpgradeMinorVersion: to.BoolPtr(true),
				ConfigurationProtectedSettings: map[string]*string{
					"omsagent.secret.key": to.StringPtr("secretKeyValue01"),
				},
				ConfigurationSettings: map[string]*string{
					"omsagent.env.clusterName": to.StringPtr("clusterName1"),
					"omsagent.secret.wsid":     to.StringPtr("a38cef99-5a89-52ed-b6db-22095c23664b"),
				},
				ExtensionType: to.StringPtr("<extension-type>"),
				ReleaseTrain:  to.StringPtr("<release-train>"),
				Scope: &armkubernetesconfiguration.Scope{
					Cluster: &armkubernetesconfiguration.ScopeCluster{
						ReleaseNamespace: to.StringPtr("<release-namespace>"),
					},
				},
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
	log.Printf("Response result: %#v\n", res.ExtensionsClientCreateResult)
}
```

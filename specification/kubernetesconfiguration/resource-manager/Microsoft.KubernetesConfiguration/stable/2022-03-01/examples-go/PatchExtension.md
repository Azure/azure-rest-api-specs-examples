```go
package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/PatchExtension.json
func ExampleExtensionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkubernetesconfiguration.NewExtensionsClient("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"rg1",
		"Microsoft.Kubernetes",
		"connectedClusters",
		"clusterName1",
		"ClusterMonitor",
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
				ReleaseTrain: to.Ptr("Preview"),
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkubernetesconfiguration%2Farmkubernetesconfiguration%2Fv1.0.0/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.

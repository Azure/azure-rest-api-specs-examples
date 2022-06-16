package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/CreateExtension.json
func ExampleExtensionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkubernetesconfiguration.NewExtensionsClient("subId1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"rg1",
		"Microsoft.Kubernetes",
		"connectedClusters",
		"clusterName1",
		"ClusterMonitor",
		armkubernetesconfiguration.Extension{
			Properties: &armkubernetesconfiguration.ExtensionProperties{
				AutoUpgradeMinorVersion: to.Ptr(true),
				ConfigurationProtectedSettings: map[string]*string{
					"omsagent.secret.key": to.Ptr("secretKeyValue01"),
				},
				ConfigurationSettings: map[string]*string{
					"omsagent.env.clusterName": to.Ptr("clusterName1"),
					"omsagent.secret.wsid":     to.Ptr("a38cef99-5a89-52ed-b6db-22095c23664b"),
				},
				ExtensionType: to.Ptr("azuremonitor-containers"),
				ReleaseTrain:  to.Ptr("Preview"),
				Scope: &armkubernetesconfiguration.Scope{
					Cluster: &armkubernetesconfiguration.ScopeCluster{
						ReleaseNamespace: to.Ptr("kube-system"),
					},
				},
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

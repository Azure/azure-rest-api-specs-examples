```go
package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// x-ms-original-file: specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/preview/2022-01-01-preview/examples/CreateSourceControlConfiguration.json
func ExampleSourceControlConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkubernetesconfiguration.NewSourceControlConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		armkubernetesconfiguration.ExtensionsClusterRp("Microsoft.Kubernetes"),
		armkubernetesconfiguration.ExtensionsClusterResourceName("connectedClusters"),
		"<cluster-name>",
		"<source-control-configuration-name>",
		armkubernetesconfiguration.SourceControlConfiguration{
			Properties: &armkubernetesconfiguration.SourceControlConfigurationProperties{
				ConfigurationProtectedSettings: map[string]*string{
					"protectedSetting1Key": to.StringPtr("protectedSetting1Value"),
				},
				EnableHelmOperator: to.BoolPtr(true),
				HelmOperatorProperties: &armkubernetesconfiguration.HelmOperatorProperties{
					ChartValues:  to.StringPtr("<chart-values>"),
					ChartVersion: to.StringPtr("<chart-version>"),
				},
				OperatorInstanceName:  to.StringPtr("<operator-instance-name>"),
				OperatorNamespace:     to.StringPtr("<operator-namespace>"),
				OperatorParams:        to.StringPtr("<operator-params>"),
				OperatorScope:         armkubernetesconfiguration.OperatorScopeType("namespace").ToPtr(),
				OperatorType:          armkubernetesconfiguration.OperatorType("Flux").ToPtr(),
				RepositoryURL:         to.StringPtr("<repository-url>"),
				SSHKnownHostsContents: to.StringPtr("<sshknown-hosts-contents>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SourceControlConfigurationsClientCreateOrUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkubernetesconfiguration%2Farmkubernetesconfiguration%2Fv0.3.0/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.

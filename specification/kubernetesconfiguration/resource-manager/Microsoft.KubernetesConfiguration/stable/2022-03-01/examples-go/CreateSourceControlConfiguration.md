Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkubernetesconfiguration%2Farmkubernetesconfiguration%2Fv0.5.0/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration/README.md) on how to add the SDK to your project and authenticate.

```go
package armkubernetesconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/stable/2022-03-01/examples/CreateSourceControlConfiguration.json
func ExampleSourceControlConfigurationsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armkubernetesconfiguration.NewSourceControlConfigurationsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-rp>",
		"<cluster-resource-name>",
		"<cluster-name>",
		"<source-control-configuration-name>",
		armkubernetesconfiguration.SourceControlConfiguration{
			Properties: &armkubernetesconfiguration.SourceControlConfigurationProperties{
				ConfigurationProtectedSettings: map[string]*string{
					"protectedSetting1Key": to.Ptr("protectedSetting1Value"),
				},
				EnableHelmOperator: to.Ptr(true),
				HelmOperatorProperties: &armkubernetesconfiguration.HelmOperatorProperties{
					ChartValues:  to.Ptr("<chart-values>"),
					ChartVersion: to.Ptr("<chart-version>"),
				},
				OperatorInstanceName:  to.Ptr("<operator-instance-name>"),
				OperatorNamespace:     to.Ptr("<operator-namespace>"),
				OperatorParams:        to.Ptr("<operator-params>"),
				OperatorScope:         to.Ptr(armkubernetesconfiguration.OperatorScopeTypeNamespace),
				OperatorType:          to.Ptr(armkubernetesconfiguration.OperatorTypeFlux),
				RepositoryURL:         to.Ptr("<repository-url>"),
				SSHKnownHostsContents: to.Ptr("<sshknown-hosts-contents>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

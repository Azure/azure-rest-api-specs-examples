Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.7.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateAwsAssumeRoleConnectorSubscription_example.json
func ExampleAccountConnectorsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurity.NewAccountConnectorsClient("20ff7fc3-e762-44dd-bd96-b71116dcdc23", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"aws_dev2",
		armsecurity.ConnectorSetting{
			Properties: &armsecurity.ConnectorSettingProperties{
				AuthenticationDetails: &armsecurity.AwAssumeRoleAuthenticationDetailsProperties{
					AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeAwsAssumeRole),
					AwsAssumeRoleArn:   to.Ptr("arn:aws:iam::81231569658:role/AscConnector"),
					AwsExternalID:      to.Ptr("20ff7fc3-e762-44dd-bd96-b71116dcdc23"),
				},
				HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
					AutoProvision: to.Ptr(armsecurity.AutoProvisionOn),
					ProxyServer: &armsecurity.ProxyServerProperties{
						IP:   to.Ptr("167.220.197.140"),
						Port: to.Ptr("34"),
					},
					Region:            to.Ptr("West US 2"),
					ResourceGroupName: to.Ptr("AwsConnectorRG"),
					ServicePrincipal: &armsecurity.ServicePrincipalProperties{
						ApplicationID: to.Ptr("ad9bcd79-be9c-45ab-abd8-80ca1654a7d1"),
						Secret:        to.Ptr("<secret>"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

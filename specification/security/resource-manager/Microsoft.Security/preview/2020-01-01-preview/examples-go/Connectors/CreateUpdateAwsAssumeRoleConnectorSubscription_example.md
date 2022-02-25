Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.4.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2020-01-01-preview/examples/Connectors/CreateUpdateAwsAssumeRoleConnectorSubscription_example.json
func ExampleAccountConnectorsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAccountConnectorsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<connector-name>",
		armsecurity.ConnectorSetting{
			Properties: &armsecurity.ConnectorSettingProperties{
				AuthenticationDetails: &armsecurity.AwAssumeRoleAuthenticationDetailsProperties{
					AuthenticationType: armsecurity.AuthenticationType("awsAssumeRole").ToPtr(),
					AwsAssumeRoleArn:   to.StringPtr("<aws-assume-role-arn>"),
					AwsExternalID:      to.StringPtr("<aws-external-id>"),
				},
				HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
					AutoProvision: armsecurity.AutoProvision("On").ToPtr(),
					ProxyServer: &armsecurity.ProxyServerProperties{
						IP:   to.StringPtr("<ip>"),
						Port: to.StringPtr("<port>"),
					},
					Region:            to.StringPtr("<region>"),
					ResourceGroupName: to.StringPtr("<resource-group-name>"),
					ServicePrincipal: &armsecurity.ServicePrincipalProperties{
						ApplicationID: to.StringPtr("<application-id>"),
						Secret:        to.StringPtr("<secret>"),
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountConnectorsClientCreateOrUpdateResult)
}
```

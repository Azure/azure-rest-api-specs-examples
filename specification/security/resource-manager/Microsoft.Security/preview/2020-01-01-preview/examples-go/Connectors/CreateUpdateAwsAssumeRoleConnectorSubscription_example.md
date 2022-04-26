Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.6.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

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
		return
	}
	ctx := context.Background()
	client, err := armsecurity.NewAccountConnectorsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<connector-name>",
		armsecurity.ConnectorSetting{
			Properties: &armsecurity.ConnectorSettingProperties{
				AuthenticationDetails: &armsecurity.AwAssumeRoleAuthenticationDetailsProperties{
					AuthenticationType: to.Ptr(armsecurity.AuthenticationTypeAwsAssumeRole),
					AwsAssumeRoleArn:   to.Ptr("<aws-assume-role-arn>"),
					AwsExternalID:      to.Ptr("<aws-external-id>"),
				},
				HybridComputeSettings: &armsecurity.HybridComputeSettingsProperties{
					AutoProvision: to.Ptr(armsecurity.AutoProvisionOn),
					ProxyServer: &armsecurity.ProxyServerProperties{
						IP:   to.Ptr("<ip>"),
						Port: to.Ptr("<port>"),
					},
					Region:            to.Ptr("<region>"),
					ResourceGroupName: to.Ptr("<resource-group-name>"),
					ServicePrincipal: &armsecurity.ServicePrincipalProperties{
						ApplicationID: to.Ptr("<application-id>"),
						Secret:        to.Ptr("<secret>"),
					},
				},
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

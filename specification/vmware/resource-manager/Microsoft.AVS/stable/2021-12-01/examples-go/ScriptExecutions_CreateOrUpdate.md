Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Favs%2Farmavs%2Fv1.0.0/sdk/resourcemanager/avs/armavs/README.md) on how to add the SDK to your project and authenticate.

```go
package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/ScriptExecutions_CreateOrUpdate.json
func ExampleScriptExecutionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armavs.NewScriptExecutionsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"group1",
		"cloud1",
		"addSsoServer",
		armavs.ScriptExecution{
			Properties: &armavs.ScriptExecutionProperties{
				HiddenParameters: []armavs.ScriptExecutionParameterClassification{
					&armavs.ScriptSecureStringExecutionParameter{
						Name:        to.Ptr("Password"),
						Type:        to.Ptr(armavs.ScriptExecutionParameterTypeSecureValue),
						SecureValue: to.Ptr("PlaceholderPassword"),
					}},
				Parameters: []armavs.ScriptExecutionParameterClassification{
					&armavs.ScriptStringExecutionParameter{
						Name:  to.Ptr("DomainName"),
						Type:  to.Ptr(armavs.ScriptExecutionParameterTypeValue),
						Value: to.Ptr("placeholderDomain.local"),
					},
					&armavs.ScriptStringExecutionParameter{
						Name:  to.Ptr("BaseUserDN"),
						Type:  to.Ptr(armavs.ScriptExecutionParameterTypeValue),
						Value: to.Ptr("DC=placeholder, DC=placeholder"),
					}},
				Retention:      to.Ptr("P0Y0M60DT0H60M60S"),
				ScriptCmdletID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/AVS.PowerCommands@1.0.0/scriptCmdlets/New-SsoExternalIdentitySource"),
				Timeout:        to.Ptr("P0Y0M0DT0H60M60S"),
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

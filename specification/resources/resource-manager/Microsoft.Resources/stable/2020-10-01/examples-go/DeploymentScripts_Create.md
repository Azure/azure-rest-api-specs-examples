Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmdeploymentscripts%2Fv0.5.0/sdk/resourcemanager/resources/armdeploymentscripts/README.md) on how to add the SDK to your project and authenticate.

```go
package armdeploymentscripts_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentscripts"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2020-10-01/examples/DeploymentScripts_Create.json
func ExampleClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdeploymentscripts.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<script-name>",
		&armdeploymentscripts.AzurePowerShellScript{
			Identity: &armdeploymentscripts.ManagedServiceIdentity{
				Type: to.Ptr(armdeploymentscripts.ManagedServiceIdentityTypeUserAssigned),
				UserAssignedIdentities: map[string]*armdeploymentscripts.UserAssignedIdentity{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scriptRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uai": {},
				},
			},
			Kind:     to.Ptr(armdeploymentscripts.ScriptTypeAzurePowerShell),
			Location: to.Ptr("<location>"),
			Properties: &armdeploymentscripts.AzurePowerShellScriptProperties{
				CleanupPreference: to.Ptr(armdeploymentscripts.CleanupOptionsAlways),
				Arguments:         to.Ptr("<arguments>"),
				RetentionInterval: to.Ptr("<retention-interval>"),
				ScriptContent:     to.Ptr("<script-content>"),
				SupportingScriptUris: []*string{
					to.Ptr("https://uri1.to.supporting.script"),
					to.Ptr("https://uri2.to.supporting.script")},
				Timeout:             to.Ptr("<timeout>"),
				AzPowerShellVersion: to.Ptr("<az-power-shell-version>"),
			},
		},
		&armdeploymentscripts.ClientBeginCreateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmdeploymentscripts%2Fv0.1.1/sdk/resourcemanager/resources/armdeploymentscripts/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2020-10-01/examples/DeploymentScripts_Create.json
func ExampleDeploymentScriptsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdeploymentscripts.NewDeploymentScriptsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<script-name>",
		&armdeploymentscripts.AzurePowerShellScript{
			DeploymentScript: armdeploymentscripts.DeploymentScript{
				Identity: &armdeploymentscripts.ManagedServiceIdentity{
					Type: armdeploymentscripts.ManagedServiceIdentityTypeUserAssigned.ToPtr(),
					UserAssignedIdentities: map[string]*armdeploymentscripts.UserAssignedIdentity{
						"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scriptRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uai": {},
					},
				},
				Kind:     armdeploymentscripts.ScriptTypeAzurePowerShell.ToPtr(),
				Location: to.StringPtr("<location>"),
			},
			Properties: &armdeploymentscripts.AzurePowerShellScriptProperties{
				DeploymentScriptPropertiesBase: armdeploymentscripts.DeploymentScriptPropertiesBase{
					CleanupPreference: armdeploymentscripts.CleanupOptionsAlways.ToPtr(),
				},
				ScriptConfigurationBase: armdeploymentscripts.ScriptConfigurationBase{
					Arguments:         to.StringPtr("<arguments>"),
					RetentionInterval: to.StringPtr("<retention-interval>"),
					ScriptContent:     to.StringPtr("<script-content>"),
					SupportingScriptUris: []*string{
						to.StringPtr("https://uri1.to.supporting.script"),
						to.StringPtr("https://uri2.to.supporting.script")},
					Timeout: to.StringPtr("<timeout>"),
				},
				AzPowerShellVersion: to.StringPtr("<az-power-shell-version>"),
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
	log.Printf("DeploymentScriptClassification.GetDeploymentScript().ID: %s\n", *res.GetDeploymentScript().ID)
}
```

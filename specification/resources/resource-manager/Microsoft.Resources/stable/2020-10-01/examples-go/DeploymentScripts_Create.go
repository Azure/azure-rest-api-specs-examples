package armdeploymentscripts_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentscripts"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2020-10-01/examples/DeploymentScripts_Create.json
func ExampleClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdeploymentscripts.NewClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"script-rg",
		"MyDeploymentScript",
		&armdeploymentscripts.AzurePowerShellScript{
			Identity: &armdeploymentscripts.ManagedServiceIdentity{
				Type: to.Ptr(armdeploymentscripts.ManagedServiceIdentityTypeUserAssigned),
				UserAssignedIdentities: map[string]*armdeploymentscripts.UserAssignedIdentity{
					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scriptRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uai": {},
				},
			},
			Kind:     to.Ptr(armdeploymentscripts.ScriptTypeAzurePowerShell),
			Location: to.Ptr("westus"),
			Properties: &armdeploymentscripts.AzurePowerShellScriptProperties{
				CleanupPreference: to.Ptr(armdeploymentscripts.CleanupOptionsAlways),
				Arguments:         to.Ptr("-Location 'westus' -Name \"*rg2\""),
				RetentionInterval: to.Ptr("PT7D"),
				ScriptContent:     to.Ptr("Param([string]$Location,[string]$Name) $deploymentScriptOutputs['test'] = 'value' Get-AzResourceGroup -Location $Location -Name $Name"),
				SupportingScriptUris: []*string{
					to.Ptr("https://uri1.to.supporting.script"),
					to.Ptr("https://uri2.to.supporting.script")},
				Timeout:             to.Ptr("PT1H"),
				AzPowerShellVersion: to.Ptr("1.7.0"),
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

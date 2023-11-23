package armdeploymentscripts_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentscripts/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/resources/resource-manager/Microsoft.Resources/stable/2020-10-01/examples/DeploymentScripts_ListBySubscription.json
func ExampleClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentscripts.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DeploymentScriptListResult = armdeploymentscripts.DeploymentScriptListResult{
		// 	Value: []armdeploymentscripts.DeploymentScriptClassification{
		// 		&armdeploymentscripts.AzurePowerShellScript{
		// 			Identity: &armdeploymentscripts.ManagedServiceIdentity{
		// 				Type: to.Ptr(armdeploymentscripts.ManagedServiceIdentityTypeUserAssigned),
		// 				UserAssignedIdentities: map[string]*armdeploymentscripts.UserAssignedIdentity{
		// 					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scriptRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uai": &armdeploymentscripts.UserAssignedIdentity{
		// 					},
		// 				},
		// 			},
		// 			Kind: to.Ptr(armdeploymentscripts.ScriptTypeAzurePowerShell),
		// 			Location: to.Ptr("westus"),
		// 			SystemData: &armdeploymentscripts.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armdeploymentscripts.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.107Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armdeploymentscripts.CreatedByTypeApplication),
		// 			},
		// 			Properties: &armdeploymentscripts.AzurePowerShellScriptProperties{
		// 				CleanupPreference: to.Ptr(armdeploymentscripts.CleanupOptionsAlways),
		// 				Outputs: map[string]any{
		// 					"output1": "value1",
		// 				},
		// 				ProvisioningState: to.Ptr(armdeploymentscripts.ScriptProvisioningStateSucceeded),
		// 				Status: &armdeploymentscripts.ScriptStatus{
		// 					ContainerInstanceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scriptRG/providers/Microsoft.ContainerInstance/containerGroups/scriptContainer"),
		// 					EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-13T23:19:45.000Z"); return t}()),
		// 					ExpirationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-13T23:19:45.000Z"); return t}()),
		// 					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-13T23:19:45.000Z"); return t}()),
		// 					StorageAccountID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scriptRG/providers/Microsoft.Storage/storageAccounts/scriptStorage"),
		// 				},
		// 				Arguments: to.Ptr("-Location 'westus' -Name \"*rg2\""),
		// 				RetentionInterval: to.Ptr("PT7D"),
		// 				ScriptContent: to.Ptr("Param([string]$Location,[string]$Name) $deploymentScriptOutputs['test'] = 'value' Get-AzResourceGroup -Location $Location -Name $Name"),
		// 				Timeout: to.Ptr("PT1H"),
		// 				AzPowerShellVersion: to.Ptr("1.7.0"),
		// 			},
		// 	}},
		// }
	}
}

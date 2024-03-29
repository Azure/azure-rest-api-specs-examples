package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c71a66dab813061f1d09982c2748a09317fe0860/specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptCmdlets_List.json
func ExampleScriptCmdletsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScriptCmdletsClient().NewListPager("group1", "{privateCloudName}", "{scriptPackageName}", nil)
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
		// page.ScriptCmdletsList = armavs.ScriptCmdletsList{
		// 	Value: []*armavs.ScriptCmdlet{
		// 		{
		// 			Name: to.Ptr("Set-AvsStoragePolicy"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds/scriptPackages/scriptCmdlets"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/{scriptPackageName}/scriptCmdlets/Set-AvsStoragePolicy"),
		// 			Properties: &armavs.ScriptCmdletProperties{
		// 				Description: to.Ptr("Allow user to set the storage policy of the specified VM"),
		// 				Parameters: []*armavs.ScriptParameter{
		// 					{
		// 						Name: to.Ptr("VM"),
		// 						Type: to.Ptr(armavs.ScriptParameterTypesString),
		// 						Description: to.Ptr("VM to set the storage policy on"),
		// 						Optional: to.Ptr(armavs.OptionalParamEnumRequired),
		// 						Visibility: to.Ptr(armavs.VisibilityParameterEnumVisible),
		// 					},
		// 					{
		// 						Name: to.Ptr("StoragePolicyName"),
		// 						Type: to.Ptr(armavs.ScriptParameterTypesString),
		// 						Description: to.Ptr("Name of the storage policy to set"),
		// 						Optional: to.Ptr(armavs.OptionalParamEnumRequired),
		// 						Visibility: to.Ptr(armavs.VisibilityParameterEnumVisible),
		// 				}},
		// 				Timeout: to.Ptr("P0Y0M0DT0H60M0S"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("New-ExternalSsoDomain"),
		// 			Type: to.Ptr("Microsoft.AVS/privateClouds/scriptPackages/scriptCmdlets"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/{scriptPackageName}/scriptCmdlets/New-ExternalSsoDomain"),
		// 			Properties: &armavs.ScriptCmdletProperties{
		// 				Description: to.Ptr("Add an external Sso domain to their vCenter"),
		// 				Parameters: []*armavs.ScriptParameter{
		// 					{
		// 						Name: to.Ptr("DomainName"),
		// 						Type: to.Ptr(armavs.ScriptParameterTypesString),
		// 						Description: to.Ptr("Domain name of the Server"),
		// 						Optional: to.Ptr(armavs.OptionalParamEnumRequired),
		// 						Visibility: to.Ptr(armavs.VisibilityParameterEnumVisible),
		// 					},
		// 					{
		// 						Name: to.Ptr("BaseUserDN"),
		// 						Type: to.Ptr(armavs.ScriptParameterTypesString),
		// 						Description: to.Ptr("Base User DN of the Server"),
		// 						Optional: to.Ptr(armavs.OptionalParamEnumRequired),
		// 						Visibility: to.Ptr(armavs.VisibilityParameterEnumVisible),
		// 					},
		// 					{
		// 						Name: to.Ptr("Password"),
		// 						Type: to.Ptr(armavs.ScriptParameterTypesSecureString),
		// 						Description: to.Ptr("Password for authenticating to the server"),
		// 						Optional: to.Ptr(armavs.OptionalParamEnumRequired),
		// 						Visibility: to.Ptr(armavs.VisibilityParameterEnumHidden),
		// 				}},
		// 				Timeout: to.Ptr("P0Y0M0DT0H60M0S"),
		// 			},
		// 	}},
		// }
	}
}

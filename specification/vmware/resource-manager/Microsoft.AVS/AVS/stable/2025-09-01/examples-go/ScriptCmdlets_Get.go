package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
)

// Generated from example definition: 2025-09-01/ScriptCmdlets_Get.json
func ExampleScriptCmdletsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armavs.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScriptCmdletsClient().Get(ctx, "group1", "cloud1", "package@1.0.2", "New-ExternalSsoDomain", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armavs.ScriptCmdletsClientGetResponse{
	// 	ScriptCmdlet: &armavs.ScriptCmdlet{
	// 		Name: to.Ptr("New-ExternalSsoDomain"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptPackages/package@1.0.2/scriptCmdlets/New-ExternalSsoDomain"),
	// 		Properties: &armavs.ScriptCmdletProperties{
	// 			Description: to.Ptr("Add an external Sso domain to their vCenter"),
	// 			Timeout: to.Ptr("P0Y0M0DT0H60M0S"),
	// 			Audience: to.Ptr(armavs.ScriptCmdletAudienceAny),
	// 			Parameters: []*armavs.ScriptParameter{
	// 				{
	// 					Name: to.Ptr("DomainName"),
	// 					Description: to.Ptr("Domain name of the Server"),
	// 					Type: to.Ptr(armavs.ScriptParameterTypesString),
	// 					Visibility: to.Ptr(armavs.VisibilityParameterEnumVisible),
	// 					Optional: to.Ptr(armavs.OptionalParamEnumRequired),
	// 				},
	// 				{
	// 					Name: to.Ptr("BaseUserDN"),
	// 					Description: to.Ptr("Base User DN of the Server"),
	// 					Type: to.Ptr(armavs.ScriptParameterTypesString),
	// 					Visibility: to.Ptr(armavs.VisibilityParameterEnumVisible),
	// 					Optional: to.Ptr(armavs.OptionalParamEnumRequired),
	// 				},
	// 				{
	// 					Name: to.Ptr("Password"),
	// 					Description: to.Ptr("Password for authenticating to the server"),
	// 					Type: to.Ptr(armavs.ScriptParameterTypesSecureString),
	// 					Visibility: to.Ptr(armavs.VisibilityParameterEnumHidden),
	// 					Optional: to.Ptr(armavs.OptionalParamEnumRequired),
	// 				},
	// 			},
	// 		},
	// 		Type: to.Ptr("Microsoft.AVS/privateClouds/scriptPackages/scriptCmdlets"),
	// 	},
	// }
}

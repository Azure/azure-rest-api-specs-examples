package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/listApplicationDefinitionsByResourceGroup.json
func ExampleApplicationDefinitionsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewApplicationDefinitionsClient().NewListByResourceGroupPager("rg", nil)
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
		// page.ApplicationDefinitionListResult = armmanagedapplications.ApplicationDefinitionListResult{
		// 	Value: []*armmanagedapplications.ApplicationDefinition{
		// 		{
		// 			Name: to.Ptr("myManagedApplicationDef"),
		// 			Type: to.Ptr("Microsoft.Solutions/applicationDefinitions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myManagedApplicationDef"),
		// 			Location: to.Ptr("East US 2"),
		// 			Properties: &armmanagedapplications.ApplicationDefinitionProperties{
		// 				Description: to.Ptr("myManagedApplicationDef description"),
		// 				Artifacts: []*armmanagedapplications.ApplicationArtifact{
		// 					{
		// 						Name: to.Ptr("CreateUiDefinition"),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeCustom),
		// 						URI: to.Ptr("https://path/to/managedApplicationCreateUiDefinition.json"),
		// 					},
		// 					{
		// 						Name: to.Ptr("ManagedApplicationResourceTemplate"),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeTemplate),
		// 						URI: to.Ptr("https://path/to/mainTemplate.json"),
		// 				}},
		// 				Authorizations: []*armmanagedapplications.ApplicationProviderAuthorization{
		// 					{
		// 						PrincipalID: to.Ptr("validprincipalguid"),
		// 						RoleDefinitionID: to.Ptr("validroleguid"),
		// 				}},
		// 				DisplayName: to.Ptr("myManagedApplicationDef"),
		// 				LockLevel: to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
		// 				PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myManagedApplicationDef2"),
		// 			Type: to.Ptr("Microsoft.Solutions/applicationDefinitions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myManagedApplicationDef2"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armmanagedapplications.ApplicationDefinitionProperties{
		// 				Description: to.Ptr("myManagedApplicationDef2 description"),
		// 				Artifacts: []*armmanagedapplications.ApplicationArtifact{
		// 					{
		// 						Name: to.Ptr("CreateUiDefinition"),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeCustom),
		// 						URI: to.Ptr("https://path/to/managedApplicationCreateUiDefinition.json"),
		// 					},
		// 					{
		// 						Name: to.Ptr("ManagedApplicationResourceTemplate"),
		// 						Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeTemplate),
		// 						URI: to.Ptr("https://path/to/mainTemplate.json"),
		// 				}},
		// 				Authorizations: []*armmanagedapplications.ApplicationProviderAuthorization{
		// 					{
		// 						PrincipalID: to.Ptr("validprincipalguid"),
		// 						RoleDefinitionID: to.Ptr("validroleguid"),
		// 				}},
		// 				DisplayName: to.Ptr("myManagedApplicationDef2"),
		// 				LockLevel: to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
		// 				PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
		// 			},
		// 	}},
		// }
	}
}

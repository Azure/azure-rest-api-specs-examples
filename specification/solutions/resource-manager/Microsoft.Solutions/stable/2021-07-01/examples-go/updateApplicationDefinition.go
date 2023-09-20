package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/updateApplicationDefinition.json
func ExampleApplicationDefinitionsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationDefinitionsClient().Update(ctx, "rg", "myManagedApplicationDef", armmanagedapplications.ApplicationDefinitionPatchable{
		Tags: map[string]*string{
			"department": to.Ptr("Finance"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationDefinition = armmanagedapplications.ApplicationDefinition{
	// 	Name: to.Ptr("myManagedApplicationDef"),
	// 	Type: to.Ptr("Microsoft.Solutions/applicationDefinitions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myManagedApplicationDef"),
	// 	Properties: &armmanagedapplications.ApplicationDefinitionProperties{
	// 		Description: to.Ptr("myManagedApplicationDef description"),
	// 		Artifacts: []*armmanagedapplications.ApplicationDefinitionArtifact{
	// 			{
	// 				Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameCreateUIDefinition),
	// 				Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeCustom),
	// 				URI: to.Ptr("https://path/to/managedApplicationCreateUiDefinition.json"),
	// 			},
	// 			{
	// 				Name: to.Ptr(armmanagedapplications.ApplicationDefinitionArtifactNameApplicationResourceTemplate),
	// 				Type: to.Ptr(armmanagedapplications.ApplicationArtifactTypeTemplate),
	// 				URI: to.Ptr("https://path/to/mainTemplate.json"),
	// 		}},
	// 		Authorizations: []*armmanagedapplications.ApplicationAuthorization{
	// 			{
	// 				PrincipalID: to.Ptr("validprincipalguid"),
	// 				RoleDefinitionID: to.Ptr("validroleguid"),
	// 		}},
	// 		DisplayName: to.Ptr("myManagedApplicationDef"),
	// 		LockLevel: to.Ptr(armmanagedapplications.ApplicationLockLevelNone),
	// 		PackageFileURI: to.Ptr("https://path/to/packagezipfile"),
	// 	},
	// }
}

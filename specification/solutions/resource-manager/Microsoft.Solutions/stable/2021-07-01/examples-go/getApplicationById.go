package armmanagedapplications_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5fb045bd44f143bae17da2e01552ae531f77d0ba/specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/getApplicationById.json
func ExampleApplicationsClient_GetByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedapplications.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().GetByID(ctx, "subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Application = armmanagedapplications.Application{
	// 	Name: to.Ptr("myManagedApplication"),
	// 	Type: to.Ptr("Microsoft.Solutions/applications"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication"),
	// 	Kind: to.Ptr("ServiceCatalog"),
	// 	Properties: &armmanagedapplications.ApplicationProperties{
	// 		ApplicationDefinitionID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/managedApplicationDefinitions/myAppDef"),
	// 		Artifacts: []*armmanagedapplications.ApplicationArtifact{
	// 		},
	// 		Authorizations: []*armmanagedapplications.ApplicationAuthorization{
	// 			{
	// 				PrincipalID: to.Ptr("validprincipalguid"),
	// 				RoleDefinitionID: to.Ptr("validroleguid"),
	// 		}},
	// 		CreatedBy: &armmanagedapplications.ApplicationClientDetails{
	// 			ApplicationID: to.Ptr("ClientApplicationId"),
	// 			Oid: to.Ptr("ClientOid"),
	// 			Puid: to.Ptr("ClientPuid"),
	// 		},
	// 		ManagedResourceGroupID: to.Ptr("/subscriptions/subid/resourceGroups/myManagedRG"),
	// 		ManagementMode: to.Ptr(armmanagedapplications.ApplicationManagementModeManaged),
	// 		ProvisioningState: to.Ptr(armmanagedapplications.ProvisioningStateSucceeded),
	// 		UpdatedBy: &armmanagedapplications.ApplicationClientDetails{
	// 			ApplicationID: to.Ptr("ClientApplicationId"),
	// 			Oid: to.Ptr("ClientOid"),
	// 			Puid: to.Ptr("ClientPuid"),
	// 		},
	// 	},
	// }
}

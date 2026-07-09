package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v3"
)

// Generated from example definition: 2026-01-01-preview/Projects_PutWithCustomizationSettings.json
func ExampleProjectsClient_BeginCreateOrUpdate_projectsCreateOrUpdateWithCustomizationSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("0ac520ee-14c0-480f-b6c9-0a90c58fffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProjectsClient().BeginCreateOrUpdate(ctx, "rg1", "DevProject", armdevcenter.Project{
		Identity: &armdevcenter.ManagedServiceIdentity{
			Type: to.Ptr(armdevcenter.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armdevcenter.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1": {},
			},
		},
		Location: to.Ptr("centralus"),
		Properties: &armdevcenter.ProjectProperties{
			Description: to.Ptr("This is my first project."),
			CustomizationSettings: &armdevcenter.ProjectCustomizationSettings{
				Identities: []*armdevcenter.ProjectCustomizationManagedIdentity{
					{
						IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1"),
						IdentityType:       to.Ptr(armdevcenter.ProjectCustomizationIdentityTypeUserAssignedIdentity),
					},
				},
			},
			DevCenterID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso"),
		},
		Tags: map[string]*string{
			"CostCenter": to.Ptr("R&D"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdevcenter.ProjectsClientCreateOrUpdateResponse{
	// 	Project: armdevcenter.Project{
	// 		Name: to.Ptr("DevProject"),
	// 		Type: to.Ptr("Microsoft.DevCenter/projects"),
	// 		ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/DevProject"),
	// 		Identity: &armdevcenter.ManagedServiceIdentity{
	// 			Type: to.Ptr(armdevcenter.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armdevcenter.UserAssignedIdentity{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1": &armdevcenter.UserAssignedIdentity{
	// 					ClientID: to.Ptr("e35621a5-f615-4a20-940e-de8a84b15abc"),
	// 					PrincipalID: to.Ptr("2111b8fc-e123-485a-b408-bf1153189494"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("centralus"),
	// 		Properties: &armdevcenter.ProjectProperties{
	// 			Description: to.Ptr("This is my first project."),
	// 			CustomizationSettings: &armdevcenter.ProjectCustomizationSettings{
	// 				Identities: []*armdevcenter.ProjectCustomizationManagedIdentity{
	// 					{
	// 						IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1"),
	// 						IdentityType: to.Ptr(armdevcenter.ProjectCustomizationIdentityTypeUserAssignedIdentity),
	// 					},
	// 				},
	// 			},
	// 			DevCenterID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/devcenters/Contoso"),
	// 			DevCenterURI: to.Ptr("https://4c7c8922-78e9-4928-aa6f-75ba59355371-contoso.centralus.devcenter.azure.com"),
	// 			ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armdevcenter.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:00:36.993Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:30:36.993Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user1"),
	// 			LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"CostCenter": to.Ptr("R&D"),
	// 		},
	// 	},
	// }
}

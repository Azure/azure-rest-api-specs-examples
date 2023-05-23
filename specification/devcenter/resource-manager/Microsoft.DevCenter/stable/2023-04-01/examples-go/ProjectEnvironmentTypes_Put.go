package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/ProjectEnvironmentTypes_Put.json
func ExampleProjectEnvironmentTypesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdevcenter.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectEnvironmentTypesClient().CreateOrUpdate(ctx, "rg1", "ContosoProj", "DevTest", armdevcenter.ProjectEnvironmentType{
		Identity: &armdevcenter.ManagedServiceIdentity{
			Type: to.Ptr(armdevcenter.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armdevcenter.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1": {},
			},
		},
		Properties: &armdevcenter.ProjectEnvironmentTypeProperties{
			CreatorRoleAssignment: &armdevcenter.ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment{
				Roles: map[string]*armdevcenter.EnvironmentRole{
					"4cbf0b6c-e750-441c-98a7-10da8387e4d6": {},
				},
			},
			DeploymentTargetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
			Status:             to.Ptr(armdevcenter.EnvironmentTypeEnableStatusEnabled),
			UserRoleAssignments: map[string]*armdevcenter.UserRoleAssignmentValue{
				"e45e3m7c-176e-416a-b466-0c5ec8298f8a": {
					Roles: map[string]*armdevcenter.EnvironmentRole{
						"4cbf0b6c-e750-441c-98a7-10da8387e4d6": {},
					},
				},
			},
		},
		Tags: map[string]*string{
			"CostCenter": to.Ptr("RnD"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProjectEnvironmentType = armdevcenter.ProjectEnvironmentType{
	// 	Name: to.Ptr("DevTest"),
	// 	Type: to.Ptr("Microsoft.DevCenter/projects/environmentTypes"),
	// 	ID: to.Ptr("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.DevCenter/projects/ContosoProj/environmentTypes/DevTest"),
	// 	SystemData: &armdevcenter.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 		CreatedBy: to.Ptr("User1@contoso.com"),
	// 		CreatedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-18T18:24:24.818Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("User1@contoso.com"),
	// 		LastModifiedByType: to.Ptr(armdevcenter.CreatedByTypeUser),
	// 	},
	// 	Identity: &armdevcenter.ManagedServiceIdentity{
	// 		Type: to.Ptr(armdevcenter.ManagedServiceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armdevcenter.UserAssignedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1": &armdevcenter.UserAssignedIdentity{
	// 				ClientID: to.Ptr("e35621a5-f615-4a20-940e-de8a84b15abc"),
	// 				PrincipalID: to.Ptr("2111b8fc-e123-485a-b408-bf1153189494"),
	// 			},
	// 		},
	// 	},
	// 	Location: to.Ptr("centralus"),
	// 	Properties: &armdevcenter.ProjectEnvironmentTypeProperties{
	// 		CreatorRoleAssignment: &armdevcenter.ProjectEnvironmentTypeUpdatePropertiesCreatorRoleAssignment{
	// 			Roles: map[string]*armdevcenter.EnvironmentRole{
	// 				"4cbf0b6c-e750-441c-98a7-10da8387e4d6": &armdevcenter.EnvironmentRole{
	// 					Description: to.Ptr("Allows Developer access to project virtual machine resources."),
	// 					RoleName: to.Ptr("Developer"),
	// 				},
	// 			},
	// 		},
	// 		DeploymentTargetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
	// 		Status: to.Ptr(armdevcenter.EnvironmentTypeEnableStatusEnabled),
	// 		UserRoleAssignments: map[string]*armdevcenter.UserRoleAssignmentValue{
	// 			"e45e3m7c-176e-416a-b466-0c5ec8298f8a": &armdevcenter.UserRoleAssignmentValue{
	// 				Roles: map[string]*armdevcenter.EnvironmentRole{
	// 					"4cbf0b6c-e750-441c-98a7-10da8387e4d6": &armdevcenter.EnvironmentRole{
	// 						Description: to.Ptr("Allows Developer access to project virtual machine resources."),
	// 						RoleName: to.Ptr("Developer"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr(armdevcenter.ProvisioningStateSucceeded),
	// 	},
	// 	Tags: map[string]*string{
	// 		"CostCenter": to.Ptr("RnD"),
	// 	},
	// }
}

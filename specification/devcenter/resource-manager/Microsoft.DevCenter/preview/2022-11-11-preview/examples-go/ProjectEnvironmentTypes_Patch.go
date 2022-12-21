package armdevcenter_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/ProjectEnvironmentTypes_Patch.json
func ExampleProjectEnvironmentTypesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdevcenter.NewProjectEnvironmentTypesClient("0ac520ee-14c0-480f-b6c9-0a90c58ffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx, "rg1", "ContosoProj", "DevTest", armdevcenter.ProjectEnvironmentTypeUpdate{
		Identity: &armdevcenter.ManagedServiceIdentity{
			Type: to.Ptr(armdevcenter.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armdevcenter.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/identityGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testidentity1": {},
			},
		},
		Properties: &armdevcenter.ProjectEnvironmentTypeUpdateProperties{
			DeploymentTargetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000"),
			Status:             to.Ptr(armdevcenter.EnableStatusEnabled),
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
	// TODO: use response item
	_ = res
}

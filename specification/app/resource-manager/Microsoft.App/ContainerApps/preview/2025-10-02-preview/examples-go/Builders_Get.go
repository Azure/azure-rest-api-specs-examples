package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/Builders_Get.json
func ExampleBuildersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBuildersClient().Get(ctx, "rg", "testBuilder", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armappcontainers.BuildersClientGetResponse{
	// 	BuilderResource: armappcontainers.BuilderResource{
	// 		Name: to.Ptr("testBuilder"),
	// 		Type: to.Ptr("Microsoft.App/builders"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/Microsoft.App/builders/testBuilder"),
	// 		Identity: &armappcontainers.ManagedServiceIdentity{
	// 			Type: to.Ptr(armappcontainers.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 			PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			UserAssignedIdentities: map[string]*armappcontainers.UserAssignedIdentity{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armappcontainers.UserAssignedIdentity{
	// 					ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 					PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armappcontainers.BuilderProperties{
	// 			ContainerRegistries: []*armappcontainers.ContainerRegistry{
	// 				{
	// 					ContainerRegistryServer: to.Ptr("test.azurecr.io"),
	// 					IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
	// 				},
	// 				{
	// 					ContainerRegistryServer: to.Ptr("test2.azurecr.io"),
	// 					IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"),
	// 				},
	// 			},
	// 			EnvironmentID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg/providers/Microsoft.App/managedEnvironments/testEnv"),
	// 			ProvisioningState: to.Ptr(armappcontainers.BuilderProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armappcontainers.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.4940669Z"); return t}()),
	// 			CreatedBy: to.Ptr("sample@microsoft.com"),
	// 			CreatedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-10-11T11:05:51.4940669Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("sample@microsoft.com"),
	// 			LastModifiedByType: to.Ptr(armappcontainers.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"key": to.Ptr("value"),
	// 		},
	// 	},
	// }
}

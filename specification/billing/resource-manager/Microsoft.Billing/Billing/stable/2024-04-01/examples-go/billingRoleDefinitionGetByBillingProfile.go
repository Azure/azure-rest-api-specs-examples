package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling/v2"
)

// Generated from example definition: 2024-04-01/billingRoleDefinitionGetByBillingProfile.json
func ExampleRoleDefinitionClient_GetByBillingProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRoleDefinitionClient().GetByBillingProfile(ctx, "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx", "40000000-aaaa-bbbb-cccc-100000000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armbilling.RoleDefinitionClientGetByBillingProfileResponse{
	// 	RoleDefinition: armbilling.RoleDefinition{
	// 		Name: to.Ptr("40000000-aaaa-bbbb-cccc-100000000000"),
	// 		Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/billingRoleDefinitions"),
	// 		ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingProfiles/xxxx-xxxx-xxx-xxx/billingRoleDefinitions/40000000-aaaa-bbbb-cccc-100000000000"),
	// 		Properties: &armbilling.RoleDefinitionProperties{
	// 			Description: to.Ptr("The Owner role gives the user all permissions including access management rights to the billing profile."),
	// 			Permissions: []*armbilling.Permission{
	// 				{
	// 					Actions: []*string{
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000000"),
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000001"),
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000002"),
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000003"),
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000004"),
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000005"),
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000008"),
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000011"),
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000012"),
	// 						to.Ptr("40000000-aaaa-bbbb-cccc-200000000013"),
	// 					},
	// 				},
	// 			},
	// 			RoleName: to.Ptr("Billing profile owner"),
	// 		},
	// 	},
	// }
}

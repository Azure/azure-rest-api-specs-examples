package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleDefinitionListByBillingAccount.json
func ExampleRoleDefinitionClient_NewListByBillingAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleDefinitionClient().NewListByBillingAccountPager("10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", nil)
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
		// page.RoleDefinitionListResult = armbilling.RoleDefinitionListResult{
		// 	Value: []*armbilling.RoleDefinition{
		// 		{
		// 			Name: to.Ptr("50000000-aaaa-bbbb-cccc-100000000000"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingRoleDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingRoleDefinitions/50000000-aaaa-bbbb-cccc-100000000000"),
		// 			Properties: &armbilling.RoleDefinitionProperties{
		// 				Permissions: []*armbilling.Permission{
		// 					{
		// 						Actions: []*string{
		// 							to.Ptr("50000000-aaaa-bbbb-cccc-200000000000"),
		// 							to.Ptr("50000000-aaaa-bbbb-cccc-200000000001"),
		// 							to.Ptr("50000000-aaaa-bbbb-cccc-200000000002"),
		// 							to.Ptr("50000000-aaaa-bbbb-cccc-200000000003"),
		// 							to.Ptr("50000000-aaaa-bbbb-cccc-200000000004"),
		// 							to.Ptr("50000000-aaaa-bbbb-cccc-200000000006"),
		// 							to.Ptr("50000000-aaaa-bbbb-cccc-200000000007"),
		// 							to.Ptr("50000000-aaaa-bbbb-cccc-200000000008"),
		// 							to.Ptr("10000000-aaaa-bbbb-cccc-200000000001"),
		// 							to.Ptr("10000000-aaaa-bbbb-cccc-200000000002"),
		// 							to.Ptr("10000000-aaaa-bbbb-cccc-200000000003"),
		// 							to.Ptr("10000000-aaaa-bbbb-cccc-200000000006"),
		// 							to.Ptr("10000000-aaaa-bbbb-cccc-200000000007"),
		// 							to.Ptr("50000000-aaaa-bbbb-cccc-200000000005"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000000"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000001"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000017"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000018"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000000"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000008"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000012"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000013"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000004"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000001"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000005"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000002"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000003"),
		// 							to.Ptr("40000000-aaaa-bbbb-cccc-200000000011"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000013"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000011"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000006"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000003"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000005"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000012"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000008"),
		// 							to.Ptr("30000000-aaaa-bbbb-cccc-200000000016")},
		// 					}},
		// 					RoleName: to.Ptr("Billing account owner"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("50000000-aaaa-bbbb-cccc-100000000001"),
		// 				Type: to.Ptr("Microsoft.Billing/billingAccounts/billingRoleDefinitions"),
		// 				ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingRoleDefinitions/50000000-aaaa-bbbb-cccc-100000000001"),
		// 				Properties: &armbilling.RoleDefinitionProperties{
		// 					Description: to.Ptr("The Contributor role gives the user all permissions except access management on the organization."),
		// 					Permissions: []*armbilling.Permission{
		// 						{
		// 							Actions: []*string{
		// 								to.Ptr("50000000-aaaa-bbbb-cccc-200000000001"),
		// 								to.Ptr("50000000-aaaa-bbbb-cccc-200000000002"),
		// 								to.Ptr("50000000-aaaa-bbbb-cccc-200000000003"),
		// 								to.Ptr("50000000-aaaa-bbbb-cccc-200000000004"),
		// 								to.Ptr("50000000-aaaa-bbbb-cccc-200000000006"),
		// 								to.Ptr("50000000-aaaa-bbbb-cccc-200000000007"),
		// 								to.Ptr("50000000-aaaa-bbbb-cccc-200000000008"),
		// 								to.Ptr("10000000-aaaa-bbbb-cccc-200000000001"),
		// 								to.Ptr("10000000-aaaa-bbbb-cccc-200000000002"),
		// 								to.Ptr("10000000-aaaa-bbbb-cccc-200000000003"),
		// 								to.Ptr("10000000-aaaa-bbbb-cccc-200000000006"),
		// 								to.Ptr("10000000-aaaa-bbbb-cccc-200000000007"),
		// 								to.Ptr("50000000-aaaa-bbbb-cccc-200000000005"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000001"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000009"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000017"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000018"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000008"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000012"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000013"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000004"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000001"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000005"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000002"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000003"),
		// 								to.Ptr("40000000-aaaa-bbbb-cccc-200000000011"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000013"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000011"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000006"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000003"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000005"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000012"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000008"),
		// 								to.Ptr("30000000-aaaa-bbbb-cccc-200000000016")},
		// 						}},
		// 						RoleName: to.Ptr("Billing account contributor"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("50000000-aaaa-bbbb-cccc-100000000002"),
		// 					Type: to.Ptr("Microsoft.Billing/billingAccounts/billingRoleDefinitions"),
		// 					ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingRoleDefinitions/50000000-aaaa-bbbb-cccc-100000000002"),
		// 					Properties: &armbilling.RoleDefinitionProperties{
		// 						Description: to.Ptr("The Reader role gives the user read permissions to an organization."),
		// 						Permissions: []*armbilling.Permission{
		// 							{
		// 								Actions: []*string{
		// 									to.Ptr("50000000-aaaa-bbbb-cccc-200000000001"),
		// 									to.Ptr("50000000-aaaa-bbbb-cccc-200000000006"),
		// 									to.Ptr("50000000-aaaa-bbbb-cccc-200000000007"),
		// 									to.Ptr("10000000-aaaa-bbbb-cccc-200000000003"),
		// 									to.Ptr("50000000-aaaa-bbbb-cccc-200000000005"),
		// 									to.Ptr("30000000-aaaa-bbbb-cccc-200000000001"),
		// 									to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 									to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 									to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 									to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 									to.Ptr("30000000-aaaa-bbbb-cccc-200000000017"),
		// 									to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
		// 									to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
		// 									to.Ptr("40000000-aaaa-bbbb-cccc-200000000008"),
		// 									to.Ptr("40000000-aaaa-bbbb-cccc-200000000012")},
		// 							}},
		// 							RoleName: to.Ptr("Billing account reader"),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("50000000-aaaa-bbbb-cccc-100000000002"),
		// 						Type: to.Ptr("Microsoft.Billing/billingAccounts/billingRoleDefinitions"),
		// 						ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingRoleDefinitions/50000000-aaaa-bbbb-cccc-100000000002"),
		// 						Properties: &armbilling.RoleDefinitionProperties{
		// 							Description: to.Ptr("The Reader role gives the user read permissions to an organization."),
		// 							Permissions: []*armbilling.Permission{
		// 								{
		// 									Actions: []*string{
		// 										to.Ptr("50000000-aaaa-bbbb-cccc-200000000001"),
		// 										to.Ptr("50000000-aaaa-bbbb-cccc-200000000006"),
		// 										to.Ptr("50000000-aaaa-bbbb-cccc-200000000007"),
		// 										to.Ptr("10000000-aaaa-bbbb-cccc-200000000003"),
		// 										to.Ptr("50000000-aaaa-bbbb-cccc-200000000005"),
		// 										to.Ptr("30000000-aaaa-bbbb-cccc-200000000001"),
		// 										to.Ptr("30000000-aaaa-bbbb-cccc-200000000004"),
		// 										to.Ptr("30000000-aaaa-bbbb-cccc-200000000007"),
		// 										to.Ptr("30000000-aaaa-bbbb-cccc-200000000010"),
		// 										to.Ptr("30000000-aaaa-bbbb-cccc-200000000015"),
		// 										to.Ptr("30000000-aaaa-bbbb-cccc-200000000017"),
		// 										to.Ptr("40000000-aaaa-bbbb-cccc-200000000006"),
		// 										to.Ptr("40000000-aaaa-bbbb-cccc-200000000007"),
		// 										to.Ptr("40000000-aaaa-bbbb-cccc-200000000008"),
		// 										to.Ptr("40000000-aaaa-bbbb-cccc-200000000012")},
		// 								}},
		// 								RoleName: to.Ptr("Billing account reader"),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("50000000-aaaa-bbbb-cccc-100000000003"),
		// 							Type: to.Ptr("Microsoft.Billing/billingAccounts/billingRoleDefinitions"),
		// 							ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31/billingRoleDefinitions/50000000-aaaa-bbbb-cccc-100000000003"),
		// 							Properties: &armbilling.RoleDefinitionProperties{
		// 								Description: to.Ptr("The Purchaser role gives the user permissions to make purchases by creating new Projects and Billing Groups in an organization."),
		// 								Permissions: []*armbilling.Permission{
		// 									{
		// 										Actions: []*string{
		// 											to.Ptr("10000000-aaaa-bbbb-cccc-200000000006"),
		// 											to.Ptr("10000000-aaaa-bbbb-cccc-200000000007"),
		// 											to.Ptr("50000000-aaaa-bbbb-cccc-200000000001"),
		// 											to.Ptr("50000000-aaaa-bbbb-cccc-200000000006")},
		// 									}},
		// 									RoleName: to.Ptr("Basic Purchaser"),
		// 								},
		// 						}},
		// 					}
	}
}

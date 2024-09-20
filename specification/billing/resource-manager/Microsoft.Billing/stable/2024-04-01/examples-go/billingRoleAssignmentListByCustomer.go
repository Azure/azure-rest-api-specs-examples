package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentListByCustomer.json
func ExampleRoleAssignmentsClient_NewListByCustomerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleAssignmentsClient().NewListByCustomerPager("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", "BKM6-54VH-BG7-PGB", "703ab484-dda2-4402-827b-a74513b61e2d", &armbilling.RoleAssignmentsClientListByCustomerOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
	})
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
		// page.RoleAssignmentListResult = armbilling.RoleAssignmentListResult{
		// 	Value: []*armbilling.RoleAssignment{
		// 		{
		// 			Name: to.Ptr("30000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/customers/billingRoleAssignments"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfiles/BKM6-54VH-BG7-PGB/customers/703ab484-dda2-4402-827b-a74513b61e2d/billingRoleAssignments/30000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9"),
		// 			Properties: &armbilling.RoleAssignmentProperties{
		// 				CreatedByPrincipalID: to.Ptr(""),
		// 				CreatedByPrincipalTenantID: to.Ptr("076915e7-de10-4323-bb34-a58c904068bb"),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-19T20:10:50.102Z"); return t}()),
		// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				PrincipalType: to.Ptr(armbilling.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfiles/BKM6-54VH-BG7-PGB/customers/703ab484-dda2-4402-827b-a74513b61e2d/billingRoleDefinitions/30000000-aaaa-bbbb-cccc-100000000000"),
		// 				Scope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfiles/BKM6-54VH-BG7-PGB/customers/703ab484-dda2-4402-827b-a74513b61e2d"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("30000000-aaaa-bbbb-cccc-100000000000_b1839933-b3ac-42ca-8112-d29c43f3ab47"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/billingProfiles/customers/billingRoleAssignments"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfiles/BKM6-54VH-BG7-PGB/customers/703ab484-dda2-4402-827b-a74513b61e2d/billingRoleAssignments/30000000-aaaa-bbbb-cccc-100000000000_b1839933-b3ac-42ca-8112-d29c43f3ab47"),
		// 			Properties: &armbilling.RoleAssignmentProperties{
		// 				CreatedByPrincipalID: to.Ptr(""),
		// 				CreatedByPrincipalTenantID: to.Ptr("076915e7-de10-4323-bb34-a58c904068bb"),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-19T20:10:50.102Z"); return t}()),
		// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				PrincipalType: to.Ptr(armbilling.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfiles/BKM6-54VH-BG7-PGB/customers/703ab484-dda2-4402-827b-a74513b61e2d/billingRoleDefinitions/30000000-aaaa-bbbb-cccc-100000000000"),
		// 				Scope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30/billingProfiles/BKM6-54VH-BG7-PGB/customers/703ab484-dda2-4402-827b-a74513b61e2d"),
		// 			},
		// 	}},
		// }
	}
}

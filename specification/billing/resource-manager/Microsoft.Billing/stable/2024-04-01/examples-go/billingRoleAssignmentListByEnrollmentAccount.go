package armbilling_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billing/armbilling"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c08ac9813477921ad8295b98ced8f82d11b8f913/specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingRoleAssignmentListByEnrollmentAccount.json
func ExampleRoleAssignmentsClient_NewListByEnrollmentAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbilling.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRoleAssignmentsClient().NewListByEnrollmentAccountPager("6100092", "123456", nil)
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
		// 			Name: to.Ptr("10000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingRoleAssignments"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6100092/enrollmentAccounts/123456/billingRoleAssignments/10000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9"),
		// 			Properties: &armbilling.RoleAssignmentProperties{
		// 				CreatedByPrincipalID: to.Ptr("00000000-0000-0000-0000-000000000001"),
		// 				CreatedByPrincipalTenantID: to.Ptr("076915e7-de10-4323-bb34-a58c904068bb"),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-19T20:10:50.102Z"); return t}()),
		// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				PrincipalType: to.Ptr(armbilling.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6100092/enrollmentAccounts/123456/billingRoleDefinitions/50000000-0000-0000-0000-000000000000"),
		// 				Scope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6100092/enrollmentAccounts/123456"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("10000000-aaaa-bbbb-cccc-100000000000_b1839933-b3ac-42ca-8112-d29c43f3ab47"),
		// 			Type: to.Ptr("Microsoft.Billing/billingAccounts/enrollmentAccounts/billingRoleAssignments"),
		// 			ID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6100092/enrollmentAccounts/123456/billingRoleAssignments/10000000-aaaa-bbbb-cccc-100000000000_b1839933-b3ac-42ca-8112-d29c43f3ab47"),
		// 			Properties: &armbilling.RoleAssignmentProperties{
		// 				CreatedByPrincipalID: to.Ptr("00000000-0000-0000-0000-000000000001"),
		// 				CreatedByPrincipalTenantID: to.Ptr("076915e7-de10-4323-bb34-a58c904068bb"),
		// 				CreatedOn: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-19T20:10:50.102Z"); return t}()),
		// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				PrincipalType: to.Ptr(armbilling.PrincipalTypeUser),
		// 				RoleDefinitionID: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6100092/enrollmentAccounts/123456/billingRoleDefinitions/50000000-0000-0000-0000-000000000001"),
		// 				Scope: to.Ptr("/providers/Microsoft.Billing/billingAccounts/6100092/enrollmentAccounts/123456"),
		// 			},
		// 	}},
		// }
	}
}

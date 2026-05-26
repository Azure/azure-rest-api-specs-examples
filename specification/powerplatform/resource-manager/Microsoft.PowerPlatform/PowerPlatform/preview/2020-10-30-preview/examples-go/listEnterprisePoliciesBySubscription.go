package armpowerplatform_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform"
)

// Generated from example definition: 2020-10-30-preview/listEnterprisePoliciesBySubscription.json
func ExampleEnterprisePoliciesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpowerplatform.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEnterprisePoliciesClient().NewListBySubscriptionPager(nil)
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
		// page = armpowerplatform.EnterprisePoliciesClientListBySubscriptionResponse{
		// 	EnterprisePolicyList: armpowerplatform.EnterprisePolicyList{
		// 		Value: []*armpowerplatform.EnterprisePolicy{
		// 			{
		// 				Name: to.Ptr("enterprisePolicy1"),
		// 				Type: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.PowerPlatform/enterprisePolicies/enterprisePolicy1"),
		// 				Identity: &armpowerplatform.EnterprisePolicyIdentity{
		// 					Type: to.Ptr(armpowerplatform.ResourceIdentityTypeSystemAssigned),
		// 					SystemAssignedIdentityPrincipalID: to.Ptr("principalId"),
		// 					TenantID: to.Ptr("tenantId"),
		// 				},
		// 				Kind: to.Ptr(armpowerplatform.EnterprisePolicyKindLockbox),
		// 				Location: to.Ptr("East US"),
		// 				SystemData: &armpowerplatform.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("enterprisePolicy2"),
		// 				Type: to.Ptr("Microsoft.PowerPlatform/enterprisePolicies"),
		// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.PowerPlatform/enterprisePolicies/enterprisePolicy2"),
		// 				Identity: &armpowerplatform.EnterprisePolicyIdentity{
		// 					Type: to.Ptr(armpowerplatform.ResourceIdentityTypeSystemAssigned),
		// 					SystemAssignedIdentityPrincipalID: to.Ptr("principalId"),
		// 					TenantID: to.Ptr("tenantId"),
		// 				},
		// 				Kind: to.Ptr(armpowerplatform.EnterprisePolicyKindLockbox),
		// 				Location: to.Ptr("East US"),
		// 				SystemData: &armpowerplatform.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armpowerplatform.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8a9dbb28e788355a47dc5bad3ea5f8da212b4bf6/specification/network/resource-manager/Microsoft.Network/stable/2024-10-01/examples/NspAccessRuleList.json
func ExampleSecurityPerimeterAccessRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityPerimeterAccessRulesClient().NewListPager("rg1", "nsp1", "profile1", &armnetwork.SecurityPerimeterAccessRulesClientListOptions{Top: nil,
		SkipToken: nil,
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
		// page.NspAccessRuleListResult = armnetwork.NspAccessRuleListResult{
		// 	Value: []*armnetwork.NspAccessRule{
		// 		{
		// 			Name: to.Ptr("accessRule1"),
		// 			Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/profiles/accessRules"),
		// 			ID: to.Ptr("/subscriptions/rg1/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/profile1/accessRules/accessRule1"),
		// 			SystemData: &armnetwork.SecurityPerimeterSystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.344Z"); return t}()),
		// 				CreatedBy: to.Ptr("user"),
		// 				CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.344Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user"),
		// 				LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 			},
		// 			Properties: &armnetwork.NspAccessRuleProperties{
		// 				AddressPrefixes: []*string{
		// 					to.Ptr("10.11.0.0/16"),
		// 					to.Ptr("10.10.1.0/24")},
		// 					Direction: to.Ptr(armnetwork.AccessRuleDirectionInbound),
		// 					EmailAddresses: []*string{
		// 					},
		// 					FullyQualifiedDomainNames: []*string{
		// 					},
		// 					NetworkSecurityPerimeters: []*armnetwork.PerimeterBasedAccessRule{
		// 					},
		// 					PhoneNumbers: []*string{
		// 					},
		// 					ProvisioningState: to.Ptr(armnetwork.NspProvisioningStateSucceeded),
		// 					ServiceTags: []*string{
		// 					},
		// 					Subscriptions: []*armnetwork.SubscriptionID{
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("accessRule2"),
		// 				Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/profiles/accessRules"),
		// 				ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/profile1/accessRules/accessRule2"),
		// 				SystemData: &armnetwork.SecurityPerimeterSystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.344Z"); return t}()),
		// 					CreatedBy: to.Ptr("user"),
		// 					CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.344Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user"),
		// 					LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				},
		// 				Properties: &armnetwork.NspAccessRuleProperties{
		// 					AddressPrefixes: []*string{
		// 					},
		// 					Direction: to.Ptr(armnetwork.AccessRuleDirectionOutbound),
		// 					EmailAddresses: []*string{
		// 					},
		// 					FullyQualifiedDomainNames: []*string{
		// 						to.Ptr("paasrp1.contoso.org"),
		// 						to.Ptr("paasrp2.contoso.org")},
		// 						NetworkSecurityPerimeters: []*armnetwork.PerimeterBasedAccessRule{
		// 						},
		// 						PhoneNumbers: []*string{
		// 						},
		// 						ProvisioningState: to.Ptr(armnetwork.NspProvisioningStateSucceeded),
		// 						ServiceTags: []*string{
		// 						},
		// 						Subscriptions: []*armnetwork.SubscriptionID{
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}

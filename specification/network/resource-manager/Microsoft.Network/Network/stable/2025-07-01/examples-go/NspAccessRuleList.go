package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/NspAccessRuleList.json
func ExampleSecurityPerimeterAccessRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSecurityPerimeterAccessRulesClient().NewListPager("rg1", "nsp1", "profile1", nil)
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
		// page = armnetwork.SecurityPerimeterAccessRulesClientListResponse{
		// 	NspAccessRuleListResult: armnetwork.NspAccessRuleListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/profile1/accessRules?api-version=2025-07-01&$skipToken=10"),
		// 		Value: []*armnetwork.NspAccessRule{
		// 			{
		// 				Name: to.Ptr("accessRule1"),
		// 				Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/profiles/accessRules"),
		// 				ID: to.Ptr("/subscriptions/rg1/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/profile1/accessRules/accessRule1"),
		// 				Properties: &armnetwork.NspAccessRuleProperties{
		// 					AddressPrefixes: []*string{
		// 						to.Ptr("10.11.0.0/16"),
		// 						to.Ptr("10.10.1.0/24"),
		// 					},
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
		// 				SystemData: &armnetwork.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.3446713Z"); return t}()),
		// 					CreatedBy: to.Ptr("user"),
		// 					CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.3446713Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user"),
		// 					LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("accessRule2"),
		// 				Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/profiles/accessRules"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/profile1/accessRules/accessRule2"),
		// 				Properties: &armnetwork.NspAccessRuleProperties{
		// 					AddressPrefixes: []*string{
		// 					},
		// 					Direction: to.Ptr(armnetwork.AccessRuleDirectionOutbound),
		// 					EmailAddresses: []*string{
		// 					},
		// 					FullyQualifiedDomainNames: []*string{
		// 						to.Ptr("paasrp1.contoso.org"),
		// 						to.Ptr("paasrp2.contoso.org"),
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
		// 				SystemData: &armnetwork.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.3446713Z"); return t}()),
		// 					CreatedBy: to.Ptr("user"),
		// 					CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.3446713Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user"),
		// 					LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

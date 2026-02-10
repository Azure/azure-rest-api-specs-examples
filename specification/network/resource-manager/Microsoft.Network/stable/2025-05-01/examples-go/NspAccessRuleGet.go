package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NspAccessRuleGet.json
func ExampleSecurityPerimeterAccessRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSecurityPerimeterAccessRulesClient().Get(ctx, "rg1", "nsp1", "profile1", "accessRule1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NspAccessRule = armnetwork.NspAccessRule{
	// 	Name: to.Ptr("accessRule1"),
	// 	Type: to.Ptr("Microsoft.Network/networkSecurityPerimeters/profiles/accessRules"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityPerimeters/nsp1/profiles/profile1/accessRules/accessRule1"),
	// 	SystemData: &armnetwork.SecurityPerimeterSystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.344Z"); return t}()),
	// 		CreatedBy: to.Ptr("user"),
	// 		CreatedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-07T18:07:36.344Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user"),
	// 		LastModifiedByType: to.Ptr(armnetwork.CreatedByTypeUser),
	// 	},
	// 	Properties: &armnetwork.NspAccessRuleProperties{
	// 		AddressPrefixes: []*string{
	// 		},
	// 		Direction: to.Ptr(armnetwork.AccessRuleDirectionOutbound),
	// 		EmailAddresses: []*string{
	// 		},
	// 		FullyQualifiedDomainNames: []*string{
	// 			to.Ptr("paasrp1.contoso.org"),
	// 			to.Ptr("paasrp2.contoso.org")},
	// 			NetworkSecurityPerimeters: []*armnetwork.PerimeterBasedAccessRule{
	// 			},
	// 			PhoneNumbers: []*string{
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.NspProvisioningStateSucceeded),
	// 			ServiceTags: []*string{
	// 			},
	// 			Subscriptions: []*armnetwork.SubscriptionID{
	// 			},
	// 		},
	// 	}
}

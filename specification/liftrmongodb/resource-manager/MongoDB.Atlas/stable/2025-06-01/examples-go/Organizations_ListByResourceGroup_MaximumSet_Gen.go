package armmongodbatlas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongodbatlas/armmongodbatlas"
)

// Generated from example definition: 2025-06-01/Organizations_ListByResourceGroup_MaximumSet_Gen.json
func ExampleOrganizationsClient_NewListByResourceGroupPager_organizationsListByResourceGroupMaximumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongodbatlas.NewClientFactory("4AFC1287-D389-4265-B2D4-59B96A45CACC", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationsClient().NewListByResourceGroupPager("rgopenapi", nil)
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
		// page = armmongodbatlas.OrganizationsClientListByResourceGroupResponse{
		// 	OrganizationResourceListResult: armmongodbatlas.OrganizationResourceListResult{
		// 		Value: []*armmongodbatlas.OrganizationResource{
		// 			{
		// 				Properties: &armmongodbatlas.OrganizationProperties{
		// 					Marketplace: &armmongodbatlas.MarketplaceDetails{
		// 						SubscriptionID: to.Ptr("isxzpjgmsgvwmqaridgqgnz"),
		// 						SubscriptionStatus: to.Ptr(armmongodbatlas.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 						OfferDetails: &armmongodbatlas.OfferDetails{
		// 							PublisherID: to.Ptr("fpojslpbvfsqvckbafx"),
		// 							OfferID: to.Ptr("armdqpexqptdprwn"),
		// 							PlanID: to.Ptr("gnbu"),
		// 							PlanName: to.Ptr("tajluzdogtzyep"),
		// 							TermUnit: to.Ptr("wzbofhcxshswsdtq"),
		// 							TermID: to.Ptr("lfppyeno"),
		// 						},
		// 					},
		// 					User: &armmongodbatlas.UserDetails{
		// 						FirstName: to.Ptr("xyiejuwmngdgwkuj"),
		// 						LastName: to.Ptr("qvvxiyucekcxfhkyoasgiluqnk"),
		// 						EmailAddress: to.Ptr(".K_@e7N-g1.xjqnbPs"),
		// 						Upn: to.Ptr("mlcjscht"),
		// 						PhoneNumber: to.Ptr("hkfsjzqbqsskl"),
		// 						CompanyName: to.Ptr("iupxjimdyphrdncsikbuqkeics"),
		// 					},
		// 					ProvisioningState: to.Ptr(armmongodbatlas.ResourceProvisioningStateSucceeded),
		// 					PartnerProperties: &armmongodbatlas.PartnerProperties{
		// 						OrganizationID: to.Ptr("nshpooeewedeppupwedibg"),
		// 						RedirectURL: to.Ptr("cbxwtehraetlluocdihfgchvjzockn"),
		// 						OrganizationName: to.Ptr("U.1-:7"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("tjmlissxvwsk"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/MongoDB.Atlas/organizations/testorg"),
		// 				Name: to.Ptr("ginfhwxvil"),
		// 				Type: to.Ptr("zmnehscunk"),
		// 				SystemData: &armmongodbatlas.SystemData{
		// 					CreatedBy: to.Ptr("hetkdvvtqdzyaxbgvpgvthdhw"),
		// 					CreatedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-19T10:17:41.372Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("aghbfhfodcoagjdcmutdqwh"),
		// 					LastModifiedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-19T10:17:41.372Z"); return t}()),
		// 				},
		// 				Identity: &armmongodbatlas.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("123e4567-e89b-12d3-a456-426614174000"),
		// 					TenantID: to.Ptr("123e4567-e89b-12d3-a456-426614174000"),
		// 					Type: to.Ptr(armmongodbatlas.ManagedServiceIdentityTypeNone),
		// 					UserAssignedIdentities: map[string]*armmongodbatlas.UserAssignedIdentity{
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

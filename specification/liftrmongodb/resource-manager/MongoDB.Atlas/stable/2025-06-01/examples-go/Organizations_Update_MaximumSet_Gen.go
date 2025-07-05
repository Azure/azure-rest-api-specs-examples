package armmongodbatlas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongodbatlas/armmongodbatlas"
)

// Generated from example definition: 2025-06-01/Organizations_Update_MaximumSet_Gen.json
func ExampleOrganizationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongodbatlas.NewClientFactory("422A4D59-A5BC-4DBB-8831-EC666633F64F", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationsClient().BeginUpdate(ctx, "rgopenapi", "U.1-:7", armmongodbatlas.OrganizationResourceUpdate{
		Tags: map[string]*string{},
		Properties: &armmongodbatlas.OrganizationResourceUpdateProperties{
			User: &armmongodbatlas.UserDetails{
				FirstName:    to.Ptr("btyhwmlbzzihjfimviefebg"),
				LastName:     to.Ptr("xx"),
				EmailAddress: to.Ptr(".K_@e7N-g1.xjqnbPs"),
				Upn:          to.Ptr("mxtbogd"),
				PhoneNumber:  to.Ptr("isvc"),
				CompanyName:  to.Ptr("oztteysco"),
			},
			PartnerProperties: &armmongodbatlas.PartnerProperties{
				OrganizationID:   to.Ptr("vugtqrobendjkinziswxlqueouo"),
				RedirectURL:      to.Ptr("cbxwtehraetlluocdihfgchvjzockn"),
				OrganizationName: to.Ptr("U.1-:7"),
			},
		},
		Identity: &armmongodbatlas.ManagedServiceIdentity{
			Type:                   to.Ptr(armmongodbatlas.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armmongodbatlas.UserAssignedIdentity{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmongodbatlas.OrganizationsClientUpdateResponse{
	// 	OrganizationResource: &armmongodbatlas.OrganizationResource{
	// 		Properties: &armmongodbatlas.OrganizationProperties{
	// 			Marketplace: &armmongodbatlas.MarketplaceDetails{
	// 				SubscriptionID: to.Ptr("isxzpjgmsgvwmqaridgqgnz"),
	// 				SubscriptionStatus: to.Ptr(armmongodbatlas.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 				OfferDetails: &armmongodbatlas.OfferDetails{
	// 					PublisherID: to.Ptr("fpojslpbvfsqvckbafx"),
	// 					OfferID: to.Ptr("armdqpexqptdprwn"),
	// 					PlanID: to.Ptr("gnbu"),
	// 					PlanName: to.Ptr("tajluzdogtzyep"),
	// 					TermUnit: to.Ptr("wzbofhcxshswsdtq"),
	// 					TermID: to.Ptr("lfppyeno"),
	// 				},
	// 			},
	// 			User: &armmongodbatlas.UserDetails{
	// 				FirstName: to.Ptr("btyhwmlbzzihjfimviefebg"),
	// 				LastName: to.Ptr("xx"),
	// 				EmailAddress: to.Ptr(".K_@e7N-g1.xjqnbPs"),
	// 				Upn: to.Ptr("mxtbogd"),
	// 				PhoneNumber: to.Ptr("isvc"),
	// 				CompanyName: to.Ptr("oztteysco"),
	// 			},
	// 			PartnerProperties: &armmongodbatlas.PartnerProperties{
	// 				OrganizationID: to.Ptr("vugtqrobendjkinziswxlqueouo"),
	// 				RedirectURL: to.Ptr("cbxwtehraetlluocdihfgchvjzockn"),
	// 				OrganizationName: to.Ptr("U.1-:7"),
	// 			},
	// 			ProvisioningState: to.Ptr(armmongodbatlas.ResourceProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("tjmlissxvwsk"),
	// 		Identity: &armmongodbatlas.ManagedServiceIdentity{
	// 			Type: to.Ptr(armmongodbatlas.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armmongodbatlas.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("123e4567-e89b-12d3-a456-426614174000"),
	// 			TenantID: to.Ptr("123e4567-e89b-12d3-a456-426614174000"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"),
	// 		Name: to.Ptr("ginfhwxvil"),
	// 		Type: to.Ptr("zmnehscunk"),
	// 		SystemData: &armmongodbatlas.SystemData{
	// 			CreatedBy: to.Ptr("hetkdvvtqdzyaxbgvpgvthdhw"),
	// 			CreatedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-19T10:17:41.372Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("aghbfhfodcoagjdcmutdqwh"),
	// 			LastModifiedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-19T10:17:41.372Z"); return t}()),
	// 		},
	// 	},
	// }
}

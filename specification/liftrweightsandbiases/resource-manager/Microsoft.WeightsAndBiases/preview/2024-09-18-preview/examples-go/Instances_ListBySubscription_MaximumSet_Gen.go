package armweightsandbiases_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/weightsandbiases/armweightsandbiases"
)

// Generated from example definition: 2024-09-18-preview/Instances_ListBySubscription_MaximumSet_Gen.json
func ExampleInstancesClient_NewListBySubscriptionPager_instancesListBySubscriptionGeneratedByMaximumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armweightsandbiases.NewClientFactory("0BCB047F-CB71-4DFE-B0BD-88519F411C2F", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewInstancesClient().NewListBySubscriptionPager(nil)
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
		// page = armweightsandbiases.InstancesClientListBySubscriptionResponse{
		// 	InstanceResourceListResult: armweightsandbiases.InstanceResourceListResult{
		// 		Value: []*armweightsandbiases.InstanceResource{
		// 			{
		// 				Properties: &armweightsandbiases.InstanceProperties{
		// 					Marketplace: &armweightsandbiases.MarketplaceDetails{
		// 						SubscriptionID: to.Ptr("0BCB047F-CB71-4DFE-B0BD-88519F411C2F"),
		// 						SubscriptionStatus: to.Ptr(armweightsandbiases.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 						OfferDetails: &armweightsandbiases.OfferDetails{
		// 							PublisherID: to.Ptr("0BCB047F-CB71-4DFE-B0BD-88519F411C2F"),
		// 							OfferID: to.Ptr("0BCB047F-CB71-4DFE-B0BD-88519F411C2F"),
		// 							PlanID: to.Ptr("0BCB047F-CB71-4DFE-B0BD-88519F411C2F"),
		// 							PlanName: to.Ptr("adysakgqlryufffz"),
		// 							TermUnit: to.Ptr("dgrkojow"),
		// 							TermID: to.Ptr("0BCB047F-CB71-4DFE-B0BD-88519F411C2F"),
		// 						},
		// 					},
		// 					User: &armweightsandbiases.UserDetails{
		// 						FirstName: to.Ptr("kiiehcojcldrlndoid"),
		// 						LastName: to.Ptr("zhkvsfqvthwkfkvgxcruyud"),
		// 						EmailAddress: to.Ptr("user@outlook.com"),
		// 						Upn: to.Ptr("rmjpgqchpbw"),
		// 						PhoneNumber: to.Ptr("cogmqmuwfcpstkwbzgkgo"),
		// 					},
		// 					ProvisioningState: to.Ptr(armweightsandbiases.ResourceProvisioningStateSucceeded),
		// 					PartnerProperties: &armweightsandbiases.PartnerProperties{
		// 						Region: to.Ptr(armweightsandbiases.RegionEastus),
		// 						Subdomain: to.Ptr("xkecokwnwtkwnkxfgucmzybzzb"),
		// 					},
		// 					SingleSignOnProperties: &armweightsandbiases.SingleSignOnPropertiesV2{
		// 						Type: to.Ptr(armweightsandbiases.SingleSignOnTypeSaml),
		// 						State: to.Ptr(armweightsandbiases.SingleSignOnStatesInitial),
		// 						EnterpriseAppID: to.Ptr("hkxtmpv"),
		// 						URL: to.Ptr("iqlemoksqdygqyxpp"),
		// 						AADDomains: []*string{
		// 							to.Ptr("mxnw"),
		// 						},
		// 					},
		// 				},
		// 				Identity: &armweightsandbiases.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("0BCB047F-CB71-4DFE-B0BD-88519F411C2F"),
		// 					TenantID: to.Ptr("0BCB047F-CB71-4DFE-B0BD-88519F411C2F"),
		// 					Type: to.Ptr(armweightsandbiases.ManagedServiceIdentityTypeNone),
		// 					UserAssignedIdentities: map[string]*armweightsandbiases.UserAssignedIdentity{
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 				},
		// 				Location: to.Ptr("pudewmshbcvbt"),
		// 				ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Microsoft.WeightsAndBiases/instance/myinstance"),
		// 				Name: to.Ptr("khwxivnjhglqlobab"),
		// 				Type: to.Ptr("xulngwcujdafqwmdbjdkwo"),
		// 				SystemData: &armweightsandbiases.SystemData{
		// 					CreatedBy: to.Ptr("lhixxxtbajbjlaytguxuewxe"),
		// 					CreatedByType: to.Ptr(armweightsandbiases.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-18T18:08:48.680Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("qqyff"),
		// 					LastModifiedByType: to.Ptr(armweightsandbiases.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-18T18:08:48.680Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}

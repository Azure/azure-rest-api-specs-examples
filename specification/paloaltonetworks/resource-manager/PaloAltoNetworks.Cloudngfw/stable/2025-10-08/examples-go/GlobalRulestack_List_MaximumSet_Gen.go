package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/GlobalRulestack_List_MaximumSet_Gen.json
func ExampleGlobalRulestackClient_NewListPager_globalRulestackListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGlobalRulestackClient().NewListPager(nil)
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
		// page = armpanngfw.GlobalRulestackClientListResponse{
		// 	GlobalRulestackResourceListResult: armpanngfw.GlobalRulestackResourceListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/subid/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks?api-version=2025-10-08&$skiptoken=xyz"),
		// 		Value: []*armpanngfw.GlobalRulestackResource{
		// 			{
		// 				Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				Type: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 				ID: to.Ptr("aaaaaaaaaa"),
		// 				Identity: &armpanngfw.AzureResourceManagerManagedIdentityProperties{
		// 					Type: to.Ptr(armpanngfw.ManagedIdentityTypeNone),
		// 					PrincipalID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					TenantID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					UserAssignedIdentities: map[string]*armpanngfw.AzureResourceManagerUserAssignedIdentity{
		// 						"key16": &armpanngfw.AzureResourceManagerUserAssignedIdentity{
		// 							ClientID: to.Ptr("aaaa"),
		// 							PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
		// 						},
		// 					},
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armpanngfw.RulestackProperties{
		// 					Description: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 					AssociatedSubscriptions: []*string{
		// 						to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					},
		// 					DefaultMode: to.Ptr(armpanngfw.DefaultModeIPS),
		// 					MinAppIDVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		// 					PanEtag: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c12"),
		// 					PanLocation: to.Ptr("eastus"),
		// 					ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
		// 					Scope: to.Ptr(armpanngfw.ScopeTypeLOCAL),
		// 					SecurityServices: &armpanngfw.SecurityServices{
		// 						AntiSpywareProfile: to.Ptr("aaaaaaaaaa"),
		// 						AntiVirusProfile: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 						DNSSubscription: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						FileBlockingProfile: to.Ptr("aaaaa"),
		// 						OutboundTrustCertificate: to.Ptr("aaaaaa"),
		// 						OutboundUnTrustCertificate: to.Ptr("aaaaaaaa"),
		// 						URLFilteringProfile: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						VulnerabilityProfile: to.Ptr("aaaaaaaaaa"),
		// 					},
		// 				},
		// 				SystemData: &armpanngfw.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 					CreatedBy: to.Ptr("praval"),
		// 					CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("praval"),
		// 					LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

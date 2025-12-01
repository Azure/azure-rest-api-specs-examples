package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/LocalRulestacks_CreateOrUpdate_MaximumSet_Gen.json
func ExampleLocalRulestacksClient_BeginCreateOrUpdate_localRulestacksCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("2bf4a339-294d-4c25-b0b2-ef649e9f5c27", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLocalRulestacksClient().BeginCreateOrUpdate(ctx, "rgopenapi", "lrs1", armpanngfw.LocalRulestackResource{
		Identity: &armpanngfw.AzureResourceManagerManagedIdentityProperties{
			Type: to.Ptr(armpanngfw.ManagedIdentityTypeNone),
			UserAssignedIdentities: map[string]*armpanngfw.AzureResourceManagerUserAssignedIdentity{
				"key16": {
					ClientID:    to.Ptr("aaaa"),
					PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
				},
			},
		},
		Location: to.Ptr("eastus"),
		Properties: &armpanngfw.RulestackProperties{
			Description: to.Ptr("local rulestacks"),
			AssociatedSubscriptions: []*string{
				to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c27"),
			},
			DefaultMode:       to.Ptr(armpanngfw.DefaultModeIPS),
			MinAppIDVersion:   to.Ptr("8.5.3"),
			PanEtag:           to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c12"),
			PanLocation:       to.Ptr("eastus"),
			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
			Scope:             to.Ptr(armpanngfw.ScopeTypeLOCAL),
			SecurityServices: &armpanngfw.SecurityServices{
				AntiSpywareProfile:         to.Ptr("default"),
				AntiVirusProfile:           to.Ptr("default"),
				DNSSubscription:            to.Ptr("default"),
				FileBlockingProfile:        to.Ptr("default"),
				OutboundTrustCertificate:   to.Ptr("default"),
				OutboundUnTrustCertificate: to.Ptr("default"),
				URLFilteringProfile:        to.Ptr("default"),
				VulnerabilityProfile:       to.Ptr("default"),
			},
		},
		Tags: map[string]*string{
			"tagName": to.Ptr("value"),
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
	// res = armpanngfw.LocalRulestacksClientCreateOrUpdateResponse{
	// 	LocalRulestackResource: &armpanngfw.LocalRulestackResource{
	// 		Name: to.Ptr("armid1"),
	// 		Type: to.Ptr("localRulestacks"),
	// 		ID: to.Ptr("/subscriptions/2bf4a339-294d-4c25-b0b2-ef649e9f5c27/resourceGroups/firewall-rg/providers/PaloAltoNetworks.Cloudngfw/localrulestacks/armid1"),
	// 		Identity: &armpanngfw.AzureResourceManagerManagedIdentityProperties{
	// 			Type: to.Ptr(armpanngfw.ManagedIdentityTypeNone),
	// 			PrincipalID: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c27"),
	// 			TenantID: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c27"),
	// 			UserAssignedIdentities: map[string]*armpanngfw.AzureResourceManagerUserAssignedIdentity{
	// 				"key16": &armpanngfw.AzureResourceManagerUserAssignedIdentity{
	// 					ClientID: to.Ptr("aaaa"),
	// 					PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armpanngfw.RulestackProperties{
	// 			Description: to.Ptr("local rulestacks"),
	// 			AssociatedSubscriptions: []*string{
	// 				to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c27"),
	// 			},
	// 			DefaultMode: to.Ptr(armpanngfw.DefaultModeIPS),
	// 			MinAppIDVersion: to.Ptr("8.5.3"),
	// 			PanEtag: to.Ptr("2bf4a339-294d-4c25-b0b2-ef649e9f5c12"),
	// 			PanLocation: to.Ptr("eastus"),
	// 			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
	// 			Scope: to.Ptr(armpanngfw.ScopeTypeLOCAL),
	// 			SecurityServices: &armpanngfw.SecurityServices{
	// 				AntiSpywareProfile: to.Ptr("default"),
	// 				AntiVirusProfile: to.Ptr("default"),
	// 				DNSSubscription: to.Ptr("default"),
	// 				FileBlockingProfile: to.Ptr("default"),
	// 				OutboundTrustCertificate: to.Ptr("default"),
	// 				OutboundUnTrustCertificate: to.Ptr("default"),
	// 				URLFilteringProfile: to.Ptr("default"),
	// 				VulnerabilityProfile: to.Ptr("default"),
	// 			},
	// 		},
	// 		SystemData: &armpanngfw.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 			CreatedBy: to.Ptr("praval"),
	// 			CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("praval"),
	// 			LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tagName": to.Ptr("value"),
	// 		},
	// 	},
	// }
}

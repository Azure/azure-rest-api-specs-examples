package armredhatopenshifthcp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshifthcp/armredhatopenshifthcp"
)

// Generated from example definition: 2026-09-01-preview/ExternalAuths_Update_MaximumSet_Gen.json
func ExampleExternalAuthsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshifthcp.NewClientFactory("FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExternalAuthsClient().BeginUpdate(ctx, "rgopenapi", "hcpCluster-name", "my-cool-auth", armredhatopenshifthcp.ExternalAuth{
		Properties: &armredhatopenshifthcp.ExternalAuthProperties{
			Issuer: &armredhatopenshifthcp.TokenIssuerProfile{
				URL: to.Ptr("https://microsoft.com/a"),
				Audiences: []*string{
					to.Ptr("audience1"),
					to.Ptr("audience2"),
					to.Ptr("audience3"),
					to.Ptr("audience4"),
					to.Ptr("audience5"),
				},
				CA: to.Ptr("rgmklhpshpjkbpjskqxtyfwetjjxr"),
			},
			Clients: []*armredhatopenshifthcp.ExternalAuthClientProfile{
				{
					Component: &armredhatopenshifthcp.ExternalAuthClientComponentProfile{
						Name:                to.Ptr("my-cool-component"),
						AuthClientNamespace: to.Ptr("my-cool-namespace"),
					},
					ClientID: to.Ptr("vobxtzobefgl"),
					ExtraScopes: []*string{
						to.Ptr("ejmvezdxvoozyiickteiqnvpxqciep"),
					},
					Type: to.Ptr(armredhatopenshifthcp.ExternalAuthClientTypeConfidential),
				},
			},
			Claim: &armredhatopenshifthcp.ExternalAuthClaimProfile{
				Mappings: &armredhatopenshifthcp.TokenClaimMappingsProfile{
					Username: &armredhatopenshifthcp.UsernameClaimProfile{
						Claim:        to.Ptr("nmaleeslaspkxdurlxhdntydjdcdqmwizhqpgtywqzzykfvxnouqlewuwqyqlejnddtlmudupjlndnogagnkbnupmpxjsplsfbpoknppcbsjbymnlqmmtukbaiaipzevwugtrgxuxqgwlevtdtabxbcauvuwjqzngklgovnnjwcliigxeedcum"),
						Prefix:       to.Ptr("krxszffgjhffwcszyzttmujlinm"),
						PrefixPolicy: to.Ptr(armredhatopenshifthcp.UsernameClaimPrefixPolicy("grjqszciuqlznueyltsmgec")),
					},
					Groups: &armredhatopenshifthcp.GroupClaimProfile{
						Claim:  to.Ptr("yrqawnseinzjlcevwxetagxeqkxoepjoctyrvddrfozociinj"),
						Prefix: to.Ptr("ajnojzalbh"),
					},
				},
				ValidationRules: []*armredhatopenshifthcp.TokenClaimValidationRule{
					{
						Type: to.Ptr(armredhatopenshifthcp.TokenValidationRuleTypeRequiredClaim),
						RequiredClaim: &armredhatopenshifthcp.TokenRequiredClaim{
							Claim:         to.Ptr("ciapdmvrnfitudpx"),
							RequiredValue: to.Ptr("mqzzjiozgxfgflhdrnwawpke"),
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armredhatopenshifthcp.ExternalAuthsClientUpdateResponse{
	// 	ExternalAuth: armredhatopenshifthcp.ExternalAuth{
	// 		Properties: &armredhatopenshifthcp.ExternalAuthProperties{
	// 			ProvisioningState: to.Ptr(armredhatopenshifthcp.ExternalAuthProvisioningStateSucceeded),
	// 			Issuer: &armredhatopenshifthcp.TokenIssuerProfile{
	// 				URL: to.Ptr("https://microsoft.com/a"),
	// 				Audiences: []*string{
	// 					to.Ptr("audience1"),
	// 					to.Ptr("audience2"),
	// 					to.Ptr("audience3"),
	// 					to.Ptr("audience4"),
	// 					to.Ptr("audience5"),
	// 				},
	// 				CA: to.Ptr("rgmklhpshpjkbpjskqxtyfwetjjxr"),
	// 			},
	// 			Clients: []*armredhatopenshifthcp.ExternalAuthClientProfile{
	// 				{
	// 					Component: &armredhatopenshifthcp.ExternalAuthClientComponentProfile{
	// 						Name: to.Ptr("my-cool-component"),
	// 						AuthClientNamespace: to.Ptr("my-cool-namespace"),
	// 					},
	// 					ClientID: to.Ptr("vobxtzobefgl"),
	// 					ExtraScopes: []*string{
	// 						to.Ptr("ejmvezdxvoozyiickteiqnvpxqciep"),
	// 					},
	// 					Type: to.Ptr(armredhatopenshifthcp.ExternalAuthClientTypeConfidential),
	// 				},
	// 			},
	// 			Claim: &armredhatopenshifthcp.ExternalAuthClaimProfile{
	// 				Mappings: &armredhatopenshifthcp.TokenClaimMappingsProfile{
	// 					Username: &armredhatopenshifthcp.UsernameClaimProfile{
	// 						Claim: to.Ptr("nmaleeslaspkxdurlxhdntydjdcdqmwizhqpgtywqzzykfvxnouqlewuwqyqlejnddtlmudupjlndnogagnkbnupmpxjsplsfbpoknppcbsjbymnlqmmtukbaiaipzevwugtrgxuxqgwlevtdtabxbcauvuwjqzngklgovnnjwcliigxeedcum"),
	// 						Prefix: to.Ptr("krxszffgjhffwcszyzttmujlinm"),
	// 						PrefixPolicy: to.Ptr(armredhatopenshifthcp.UsernameClaimPrefixPolicy("grjqszciuqlznueyltsmgec")),
	// 					},
	// 					Groups: &armredhatopenshifthcp.GroupClaimProfile{
	// 						Claim: to.Ptr("yrqawnseinzjlcevwxetagxeqkxoepjoctyrvddrfozociinj"),
	// 						Prefix: to.Ptr("ajnojzalbh"),
	// 					},
	// 				},
	// 				ValidationRules: []*armredhatopenshifthcp.TokenClaimValidationRule{
	// 					{
	// 						Type: to.Ptr(armredhatopenshifthcp.TokenValidationRuleTypeRequiredClaim),
	// 						RequiredClaim: &armredhatopenshifthcp.TokenRequiredClaim{
	// 							Claim: to.Ptr("ciapdmvrnfitudpx"),
	// 							RequiredValue: to.Ptr("mqzzjiozgxfgflhdrnwawpke"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/FDEA43EA-0230-4A7D-BDEE-F3AFF2183B1D/resourceGroups/resourceGroupName/providers/Microsoft.RedHatOpenShift/resourceType/resourceName"),
	// 		Name: to.Ptr("cabt"),
	// 		Type: to.Ptr("x"),
	// 		SystemData: &armredhatopenshifthcp.SystemData{
	// 			CreatedBy: to.Ptr("lsrkqcuijqfp"),
	// 			CreatedByType: to.Ptr(armredhatopenshifthcp.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2024, time.March, 27, 14, 57, 32, 578000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("tgpmwu"),
	// 			LastModifiedByType: to.Ptr(armredhatopenshifthcp.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2024, time.March, 27, 14, 57, 32, 578000000, time.UTC)),
	// 		},
	// 	},
	// }
}

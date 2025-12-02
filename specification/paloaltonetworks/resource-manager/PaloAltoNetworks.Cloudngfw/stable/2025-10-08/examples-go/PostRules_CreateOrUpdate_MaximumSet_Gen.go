package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/PostRules_CreateOrUpdate_MaximumSet_Gen.json
func ExamplePostRulesClient_BeginCreateOrUpdate_postRulesCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPostRulesClient().BeginCreateOrUpdate(ctx, "lrs1", "1", armpanngfw.PostRulesResource{
		Properties: &armpanngfw.RuleEntry{
			Description: to.Ptr("description of post rule"),
			ActionType:  to.Ptr(armpanngfw.ActionEnumAllow),
			Applications: []*string{
				to.Ptr("app1"),
			},
			AuditComment: to.Ptr("example comment"),
			Category: &armpanngfw.Category{
				Feeds: []*string{
					to.Ptr("feed"),
				},
				URLCustom: []*string{
					to.Ptr("https://microsoft.com"),
				},
			},
			DecryptionRuleType: to.Ptr(armpanngfw.DecryptionRuleTypeEnumSSLOutboundInspection),
			Destination: &armpanngfw.DestinationAddr{
				Cidrs: []*string{
					to.Ptr("1.0.0.1/10"),
				},
				Countries: []*string{
					to.Ptr("India"),
				},
				Feeds: []*string{
					to.Ptr("feed"),
				},
				FqdnLists: []*string{
					to.Ptr("FQDN1"),
				},
				PrefixLists: []*string{
					to.Ptr("PL1"),
				},
			},
			EnableLogging:                to.Ptr(armpanngfw.StateEnumDISABLED),
			Etag:                         to.Ptr("c18e6eef-ba3e-49ee-8a85-2b36c863a9d0"),
			InboundInspectionCertificate: to.Ptr("cert1"),
			NegateDestination:            to.Ptr(armpanngfw.BooleanEnumTRUE),
			NegateSource:                 to.Ptr(armpanngfw.BooleanEnumTRUE),
			ProtocolPortList: []*string{
				to.Ptr("80"),
			},
			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
			RuleName:          to.Ptr("postRule1"),
			RuleState:         to.Ptr(armpanngfw.StateEnumDISABLED),
			Source: &armpanngfw.SourceAddr{
				Cidrs: []*string{
					to.Ptr("1.0.0.1/10"),
				},
				Countries: []*string{
					to.Ptr("India"),
				},
				Feeds: []*string{
					to.Ptr("feed"),
				},
				PrefixLists: []*string{
					to.Ptr("PL1"),
				},
			},
			Tags: []*armpanngfw.TagInfo{
				{
					Key:   to.Ptr("keyName"),
					Value: to.Ptr("value"),
				},
			},
			Protocol: to.Ptr("HTTP"),
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
	// res = armpanngfw.PostRulesClientCreateOrUpdateResponse{
	// 	PostRulesResource: &armpanngfw.PostRulesResource{
	// 		Name: to.Ptr("aaaaaaaaaa"),
	// 		Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 		ID: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 		Properties: &armpanngfw.RuleEntry{
	// 			Description: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			ActionType: to.Ptr(armpanngfw.ActionEnumAllow),
	// 			Applications: []*string{
	// 				to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			},
	// 			AuditComment: to.Ptr("aaa"),
	// 			Category: &armpanngfw.Category{
	// 				Feeds: []*string{
	// 					to.Ptr("aaaaaaaaaaaa"),
	// 				},
	// 				URLCustom: []*string{
	// 					to.Ptr("aaaaa"),
	// 				},
	// 			},
	// 			DecryptionRuleType: to.Ptr(armpanngfw.DecryptionRuleTypeEnumSSLOutboundInspection),
	// 			Destination: &armpanngfw.DestinationAddr{
	// 				Cidrs: []*string{
	// 					to.Ptr("aaaaaaa"),
	// 				},
	// 				Countries: []*string{
	// 					to.Ptr("aaaaaaaaaaaaaa"),
	// 				},
	// 				Feeds: []*string{
	// 					to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				},
	// 				FqdnLists: []*string{
	// 					to.Ptr("aaaaaaaaaaaaa"),
	// 				},
	// 				PrefixLists: []*string{
	// 					to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				},
	// 			},
	// 			EnableLogging: to.Ptr(armpanngfw.StateEnumDISABLED),
	// 			Etag: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 			InboundInspectionCertificate: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			NegateDestination: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 			NegateSource: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 			Priority: to.Ptr[int32](24),
	// 			ProtocolPortList: []*string{
	// 				to.Ptr("aaaaaaaaaaaa"),
	// 			},
	// 			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
	// 			RuleName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			RuleState: to.Ptr(armpanngfw.StateEnumDISABLED),
	// 			Source: &armpanngfw.SourceAddr{
	// 				Cidrs: []*string{
	// 					to.Ptr("aaa"),
	// 				},
	// 				Countries: []*string{
	// 					to.Ptr("aaaaa"),
	// 				},
	// 				Feeds: []*string{
	// 					to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 				},
	// 				PrefixLists: []*string{
	// 					to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 				},
	// 			},
	// 			Tags: []*armpanngfw.TagInfo{
	// 				{
	// 					Key: to.Ptr("keyName"),
	// 					Value: to.Ptr("value"),
	// 				},
	// 			},
	// 			Protocol: to.Ptr("aaaa"),
	// 		},
	// 		SystemData: &armpanngfw.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 			CreatedBy: to.Ptr("praval"),
	// 			CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("praval"),
	// 			LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}

package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/LocalRules_Get_MaximumSet_Gen.json
func ExampleLocalRulesClient_Get_localRulesGetMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("2bf4a339-294d-4c25-b0b2-ef649e9f5c27", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLocalRulesClient().Get(ctx, "firewall-rg", "lrs1", "1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armpanngfw.LocalRulesClientGetResponse{
	// 	LocalRulesResource: &armpanngfw.LocalRulesResource{
	// 		Name: to.Ptr("aaaaaaaaa"),
	// 		Type: to.Ptr("aaaaaaaaa"),
	// 		ID: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 		Properties: &armpanngfw.RuleEntry{
	// 			Description: to.Ptr("aaaaaaa"),
	// 			ActionType: to.Ptr(armpanngfw.ActionEnumAllow),
	// 			Applications: []*string{
	// 				to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 			},
	// 			AuditComment: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 			Category: &armpanngfw.Category{
	// 				Feeds: []*string{
	// 					to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				},
	// 				URLCustom: []*string{
	// 					to.Ptr("aa"),
	// 				},
	// 			},
	// 			DecryptionRuleType: to.Ptr(armpanngfw.DecryptionRuleTypeEnumSSLOutboundInspection),
	// 			Destination: &armpanngfw.DestinationAddr{
	// 				Cidrs: []*string{
	// 					to.Ptr("aaaaaaaaaaaa"),
	// 				},
	// 				Countries: []*string{
	// 					to.Ptr("aaaaa"),
	// 				},
	// 				Feeds: []*string{
	// 					to.Ptr("aaaaaaa"),
	// 				},
	// 				FqdnLists: []*string{
	// 					to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				},
	// 				PrefixLists: []*string{
	// 					to.Ptr("aaaaaaaaaaaaa"),
	// 				},
	// 			},
	// 			EnableLogging: to.Ptr(armpanngfw.StateEnumDISABLED),
	// 			Etag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 			InboundInspectionCertificate: to.Ptr("aaaaaaaaaaaaaaaa"),
	// 			NegateDestination: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 			NegateSource: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 			Priority: to.Ptr[int32](13),
	// 			ProtocolPortList: []*string{
	// 				to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
	// 			},
	// 			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
	// 			RuleName: to.Ptr("aaaaaa"),
	// 			RuleState: to.Ptr(armpanngfw.StateEnumDISABLED),
	// 			Source: &armpanngfw.SourceAddr{
	// 				Cidrs: []*string{
	// 					to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				},
	// 				Countries: []*string{
	// 					to.Ptr("aaaaaaa"),
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
	// 			Protocol: to.Ptr("aaaaaaaaaaaa"),
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

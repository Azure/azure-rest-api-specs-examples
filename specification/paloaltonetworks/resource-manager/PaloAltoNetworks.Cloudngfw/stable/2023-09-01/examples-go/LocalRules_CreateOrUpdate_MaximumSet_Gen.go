package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/LocalRules_CreateOrUpdate_MaximumSet_Gen.json
func ExampleLocalRulesClient_BeginCreateOrUpdate_localRulesCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLocalRulesClient().BeginCreateOrUpdate(ctx, "firewall-rg", "lrs1", "1", armpanngfw.LocalRulesResource{
		Properties: &armpanngfw.RuleEntry{
			Description: to.Ptr("description of local rule"),
			ActionType:  to.Ptr(armpanngfw.ActionEnumAllow),
			Applications: []*string{
				to.Ptr("app1")},
			AuditComment: to.Ptr("example comment"),
			Category: &armpanngfw.Category{
				Feeds: []*string{
					to.Ptr("feed")},
				URLCustom: []*string{
					to.Ptr("https://microsoft.com")},
			},
			DecryptionRuleType: to.Ptr(armpanngfw.DecryptionRuleTypeEnumSSLOutboundInspection),
			Destination: &armpanngfw.DestinationAddr{
				Cidrs: []*string{
					to.Ptr("1.0.0.1/10")},
				Countries: []*string{
					to.Ptr("India")},
				Feeds: []*string{
					to.Ptr("feed")},
				FqdnLists: []*string{
					to.Ptr("FQDN1")},
				PrefixLists: []*string{
					to.Ptr("PL1")},
			},
			EnableLogging:                to.Ptr(armpanngfw.StateEnumDISABLED),
			Etag:                         to.Ptr("c18e6eef-ba3e-49ee-8a85-2b36c863a9d0"),
			InboundInspectionCertificate: to.Ptr("cert1"),
			NegateDestination:            to.Ptr(armpanngfw.BooleanEnumTRUE),
			NegateSource:                 to.Ptr(armpanngfw.BooleanEnumTRUE),
			ProtocolPortList: []*string{
				to.Ptr("80")},
			ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
			RuleName:          to.Ptr("localRule1"),
			RuleState:         to.Ptr(armpanngfw.StateEnumDISABLED),
			Source: &armpanngfw.SourceAddr{
				Cidrs: []*string{
					to.Ptr("1.0.0.1/10")},
				Countries: []*string{
					to.Ptr("India")},
				Feeds: []*string{
					to.Ptr("feed")},
				PrefixLists: []*string{
					to.Ptr("PL1")},
			},
			Tags: []*armpanngfw.TagInfo{
				{
					Key:   to.Ptr("keyName"),
					Value: to.Ptr("value"),
				}},
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
	// res.LocalRulesResource = armpanngfw.LocalRulesResource{
	// 	Name: to.Ptr("aaaaaaaaa"),
	// 	Type: to.Ptr("aaaaaaaaa"),
	// 	ID: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 	SystemData: &armpanngfw.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		CreatedBy: to.Ptr("praval"),
	// 		CreatedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-09T05:08:24.229Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("praval"),
	// 		LastModifiedByType: to.Ptr(armpanngfw.CreatedByTypeUser),
	// 	},
	// 	Properties: &armpanngfw.RuleEntry{
	// 		Description: to.Ptr("aaaaaaa"),
	// 		ActionType: to.Ptr(armpanngfw.ActionEnumAllow),
	// 		Applications: []*string{
	// 			to.Ptr("aaaaaaaaaaaaaaaaaaa")},
	// 			AuditComment: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
	// 			Category: &armpanngfw.Category{
	// 				Feeds: []*string{
	// 					to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa")},
	// 					URLCustom: []*string{
	// 						to.Ptr("aa")},
	// 					},
	// 					DecryptionRuleType: to.Ptr(armpanngfw.DecryptionRuleTypeEnumSSLOutboundInspection),
	// 					Destination: &armpanngfw.DestinationAddr{
	// 						Cidrs: []*string{
	// 							to.Ptr("aaaaaaaaaaaa")},
	// 							Countries: []*string{
	// 								to.Ptr("aaaaa")},
	// 								Feeds: []*string{
	// 									to.Ptr("aaaaaaa")},
	// 									FqdnLists: []*string{
	// 										to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa")},
	// 										PrefixLists: []*string{
	// 											to.Ptr("aaaaaaaaaaaaa")},
	// 										},
	// 										EnableLogging: to.Ptr(armpanngfw.StateEnumDISABLED),
	// 										Etag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 										InboundInspectionCertificate: to.Ptr("aaaaaaaaaaaaaaaa"),
	// 										NegateDestination: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 										NegateSource: to.Ptr(armpanngfw.BooleanEnumTRUE),
	// 										Priority: to.Ptr[int32](13),
	// 										ProtocolPortList: []*string{
	// 											to.Ptr("aaaaaaaaaaaaaaaaaaaaaa")},
	// 											ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateSucceeded),
	// 											RuleName: to.Ptr("aaaaaa"),
	// 											RuleState: to.Ptr(armpanngfw.StateEnumDISABLED),
	// 											Source: &armpanngfw.SourceAddr{
	// 												Cidrs: []*string{
	// 													to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa")},
	// 													Countries: []*string{
	// 														to.Ptr("aaaaaaa")},
	// 														Feeds: []*string{
	// 															to.Ptr("aaaaaaaaaaaaaaaaaaa")},
	// 															PrefixLists: []*string{
	// 																to.Ptr("aaaaaaaaaaaaaaaaaaaa")},
	// 															},
	// 															Tags: []*armpanngfw.TagInfo{
	// 																{
	// 																	Key: to.Ptr("keyName"),
	// 																	Value: to.Ptr("value"),
	// 															}},
	// 															Protocol: to.Ptr("aaaaaaaaaaaa"),
	// 														},
	// 													}
}

package armpanngfw_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw/v2"
)

// Generated from example definition: 2025-10-08/PreRules_List_MaximumSet_Gen.json
func ExamplePreRulesClient_NewListPager_preRulesListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpanngfw.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPreRulesClient().NewListPager("lrs1", nil)
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
		// page = armpanngfw.PreRulesClientListResponse{
		// 	PreRulesResourceListResult: armpanngfw.PreRulesResourceListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/providers/PaloAltoNetworks.Cloudngfw/globalrulestacks/lrs1/preRules?api-version=2025-10-08&$skiptoken=xyz"),
		// 		Value: []*armpanngfw.PreRulesResource{
		// 			{
		// 				Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				Type: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 				ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 				Properties: &armpanngfw.RuleEntry{
		// 					Description: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					ActionType: to.Ptr(armpanngfw.ActionEnumAllow),
		// 					Applications: []*string{
		// 						to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					},
		// 					AuditComment: to.Ptr("aaa"),
		// 					Category: &armpanngfw.Category{
		// 						Feeds: []*string{
		// 							to.Ptr("aaaaaaaaaaaa"),
		// 						},
		// 						URLCustom: []*string{
		// 							to.Ptr("aaaaa"),
		// 						},
		// 					},
		// 					DecryptionRuleType: to.Ptr(armpanngfw.DecryptionRuleTypeEnumSSLOutboundInspection),
		// 					Destination: &armpanngfw.DestinationAddr{
		// 						Cidrs: []*string{
		// 							to.Ptr("aaaaaaa"),
		// 						},
		// 						Countries: []*string{
		// 							to.Ptr("aaaaaaaaaaaaaa"),
		// 						},
		// 						Feeds: []*string{
		// 							to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 						FqdnLists: []*string{
		// 							to.Ptr("aaaaaaaaaaaaa"),
		// 						},
		// 						PrefixLists: []*string{
		// 							to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 					},
		// 					EnableLogging: to.Ptr(armpanngfw.StateEnumDISABLED),
		// 					Etag: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 					InboundInspectionCertificate: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					NegateDestination: to.Ptr(armpanngfw.BooleanEnumTRUE),
		// 					NegateSource: to.Ptr(armpanngfw.BooleanEnumTRUE),
		// 					Priority: to.Ptr[int32](24),
		// 					ProtocolPortList: []*string{
		// 						to.Ptr("aaaaaaaaaaaa"),
		// 					},
		// 					ProvisioningState: to.Ptr(armpanngfw.ProvisioningStateAccepted),
		// 					RuleName: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 					RuleState: to.Ptr(armpanngfw.StateEnumDISABLED),
		// 					Source: &armpanngfw.SourceAddr{
		// 						Cidrs: []*string{
		// 							to.Ptr("aaa"),
		// 						},
		// 						Countries: []*string{
		// 							to.Ptr("aaaaa"),
		// 						},
		// 						Feeds: []*string{
		// 							to.Ptr("aaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 						PrefixLists: []*string{
		// 							to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 						},
		// 					},
		// 					Tags: []*armpanngfw.TagInfo{
		// 						{
		// 							Key: to.Ptr("keyName"),
		// 							Value: to.Ptr("value"),
		// 						},
		// 					},
		// 					Protocol: to.Ptr("aaaa"),
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

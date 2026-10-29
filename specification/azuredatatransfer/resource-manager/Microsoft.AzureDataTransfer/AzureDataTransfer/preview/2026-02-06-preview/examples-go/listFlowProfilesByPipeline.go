package armdatatransfer_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datatransfer/armdatatransfer"
)

// Generated from example definition: 2026-02-06-preview/listFlowProfilesByPipeline.json
func ExampleFlowProfilesClient_NewListByPipelinePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatatransfer.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFlowProfilesClient().NewListByPipelinePager("testRG", "testPipeline", nil)
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
		// page = armdatatransfer.FlowProfilesClientListByPipelineResponse{
		// 	FlowProfileListResult: armdatatransfer.FlowProfileListResult{
		// 		Value: []*armdatatransfer.FlowProfile{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/pipelines/testPipeline/flowProfiles/flowProfileNamea"),
		// 				Name: to.Ptr("flowProfileNamea"),
		// 				Type: to.Ptr("Microsoft.AzureDataTransfer/pipelines/flowProfiles"),
		// 				Location: to.Ptr("West Central US"),
		// 				Properties: &armdatatransfer.FlowProfileProperties{
		// 					ReplicationScenario: to.Ptr(armdatatransfer.DataClassTypeFiles),
		// 					Status: to.Ptr(armdatatransfer.FlowProfileStatusEnabled),
		// 					Description: to.Ptr("Hello world description"),
		// 					Rulesets: &armdatatransfer.FlowProfileRulesets{
		// 						MimeFilters: &armdatatransfer.MimeFilterRuleset{
		// 							Type: to.Ptr(armdatatransfer.FilterTypeAllow),
		// 							Filters: []*armdatatransfer.MimeTypeFilter{
		// 								{
		// 									Media: to.Ptr("application/json"),
		// 									Extensions: []*string{
		// 										to.Ptr(".json"),
		// 										to.Ptr(".app"),
		// 										to.Ptr(".js"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 						TextMatching: &armdatatransfer.TextMatchingRuleset{
		// 							Deny: []*armdatatransfer.TextMatchingRulesetDeny{
		// 								{
		// 									Text: to.Ptr("hello world"),
		// 									MatchType: to.Ptr(armdatatransfer.MatchTypePartial),
		// 									CaseSensitivity: to.Ptr(armdatatransfer.CasingInsensitive),
		// 								},
		// 								{
		// 									Text: to.Ptr("hello"),
		// 									MatchType: to.Ptr(armdatatransfer.MatchTypeComplete),
		// 									CaseSensitivity: to.Ptr(armdatatransfer.CasingSensitive),
		// 								},
		// 							},
		// 						},
		// 						Antivirus: &armdatatransfer.AntivirusRuleset{
		// 							AvSolutions: []*armdatatransfer.AntivirusSolutions{
		// 								to.Ptr(armdatatransfer.AntivirusSolutionsDefender),
		// 								to.Ptr(armdatatransfer.AntivirusSolutionsClamAv),
		// 							},
		// 						},
		// 						Archives: &armdatatransfer.ArchiveRuleset{
		// 							MinimumSizeForExpansion: to.Ptr[int64](1),
		// 							MaximumExpansionSizeLimit: to.Ptr[int64](123456),
		// 							MaximumDepthLimit: to.Ptr[int64](12),
		// 							MaximumCompressionRatioLimit: to.Ptr[float64](123),
		// 						},
		// 						DataSize: &armdatatransfer.DataSizeRuleset{
		// 							Maximum: to.Ptr[int64](123456789),
		// 							Minimum: to.Ptr[int64](1),
		// 						},
		// 						XMLFilters: &armdatatransfer.XMLFilterRuleset{
		// 							DefaultNamespace: to.Ptr("testnamespace"),
		// 							Reference: to.Ptr(armdatatransfer.XMLReferenceTypeInline),
		// 							Schema: to.Ptr("<?xml version=\"1.0\" encoding=\"UTF-8\" ?> <xs:schema xmlns:xs=\"http://www.w3.org/2001/XMLSchema\"></xs:schema>"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armdatatransfer.ProvisioningStateAccepted),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.AzureDataTransfer/pipelines/testPipeline/flowProfiles/flowProfileNameb"),
		// 				Name: to.Ptr("flowProfileNameb"),
		// 				Type: to.Ptr("Microsoft.AzureDataTransfer/pipelines/flowProfiles"),
		// 				Location: to.Ptr("West Central US"),
		// 				Properties: &armdatatransfer.FlowProfileProperties{
		// 					ReplicationScenario: to.Ptr(armdatatransfer.DataClassTypeFiles),
		// 					Status: to.Ptr(armdatatransfer.FlowProfileStatusEnabled),
		// 					Description: to.Ptr("Hello world description"),
		// 					Rulesets: &armdatatransfer.FlowProfileRulesets{
		// 						MimeFilters: &armdatatransfer.MimeFilterRuleset{
		// 							Type: to.Ptr(armdatatransfer.FilterTypeAllow),
		// 							Filters: []*armdatatransfer.MimeTypeFilter{
		// 								{
		// 									Media: to.Ptr("application/json"),
		// 									Extensions: []*string{
		// 										to.Ptr(".json"),
		// 										to.Ptr(".app"),
		// 										to.Ptr(".js"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 						TextMatching: &armdatatransfer.TextMatchingRuleset{
		// 							Deny: []*armdatatransfer.TextMatchingRulesetDeny{
		// 								{
		// 									Text: to.Ptr("hello world"),
		// 									MatchType: to.Ptr(armdatatransfer.MatchTypePartial),
		// 									CaseSensitivity: to.Ptr(armdatatransfer.CasingInsensitive),
		// 								},
		// 								{
		// 									Text: to.Ptr("hello"),
		// 									MatchType: to.Ptr(armdatatransfer.MatchTypeComplete),
		// 									CaseSensitivity: to.Ptr(armdatatransfer.CasingSensitive),
		// 								},
		// 							},
		// 						},
		// 						Antivirus: &armdatatransfer.AntivirusRuleset{
		// 							AvSolutions: []*armdatatransfer.AntivirusSolutions{
		// 								to.Ptr(armdatatransfer.AntivirusSolutionsDefender),
		// 								to.Ptr(armdatatransfer.AntivirusSolutionsClamAv),
		// 							},
		// 						},
		// 						Archives: &armdatatransfer.ArchiveRuleset{
		// 							MinimumSizeForExpansion: to.Ptr[int64](1),
		// 							MaximumExpansionSizeLimit: to.Ptr[int64](123456),
		// 							MaximumDepthLimit: to.Ptr[int64](12),
		// 							MaximumCompressionRatioLimit: to.Ptr[float64](123),
		// 						},
		// 						DataSize: &armdatatransfer.DataSizeRuleset{
		// 							Maximum: to.Ptr[int64](123456789),
		// 							Minimum: to.Ptr[int64](1),
		// 						},
		// 						XMLFilters: &armdatatransfer.XMLFilterRuleset{
		// 							DefaultNamespace: to.Ptr("testnamespace"),
		// 							Reference: to.Ptr(armdatatransfer.XMLReferenceTypeInline),
		// 							Schema: to.Ptr("<?xml version=\"1.0\" encoding=\"UTF-8\" ?> <xs:schema xmlns:xs=\"http://www.w3.org/2001/XMLSchema\"></xs:schema>"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armdatatransfer.ProvisioningStateAccepted),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}

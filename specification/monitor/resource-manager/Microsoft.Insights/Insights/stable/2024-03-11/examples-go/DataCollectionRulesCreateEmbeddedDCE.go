package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-03-11/DataCollectionRulesCreateEmbeddedDCE.json
func ExampleDataCollectionRulesClient_Create_createOrUpdateDataCollectionRuleWithEmbeddedIngestionEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmonitor.NewClientFactory("703362b3-f278-4e4b-9179-c76eaf41ffc2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataCollectionRulesClient().Create(ctx, "myResourceGroup", "myCollectionRule", armmonitor.DataCollectionRuleResource{
		Kind:     to.Ptr(armmonitor.KnownDataCollectionRuleResourceKind(" Direct")),
		Location: to.Ptr("eastus"),
		Properties: &armmonitor.DataCollectionRuleResourceProperties{
			Description: to.Ptr("A Direct Ingestion Rule with builtin ingestion fqdns"),
			DataFlows: []*armmonitor.DataFlow{
				{
					Destinations: []*string{
						to.Ptr("myworkspace"),
					},
					OutputStream: to.Ptr("Custom-LOGS1_CL"),
					Streams: []*armmonitor.KnownDataFlowStreams{
						to.Ptr(armmonitor.KnownDataFlowStreams("Custom-LOGS1_CL")),
					},
					TransformKql: to.Ptr("source | extend jsonContext = parse_json(AdditionalContext) | project TimeGenerated = Time, Computer, AdditionalContext = jsonContext, CounterName=tostring(jsonContext.CounterName), CounterValue=toreal(jsonContext.CounterValue)"),
				},
			},
			Destinations: &armmonitor.DataCollectionRuleDestinations{
				LogAnalytics: []*armmonitor.LogAnalyticsDestination{
					{
						Name:                to.Ptr("centralWorkspace"),
						WorkspaceResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/centralTeamWorkspace"),
					},
				},
			},
			StreamDeclarations: map[string]*armmonitor.StreamDeclaration{
				"Custom-LOGS1_CL": {
					Columns: []*armmonitor.ColumnDefinition{
						{
							Name: to.Ptr("Time"),
							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeDatetime),
						},
						{
							Name: to.Ptr("Computer"),
							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeString),
						},
						{
							Name: to.Ptr("AdditionalContext"),
							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeString),
						},
						{
							Name: to.Ptr("CounterName"),
							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeString),
						},
						{
							Name: to.Ptr("CounterValue"),
							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeReal),
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armmonitor.DataCollectionRulesClientCreateResponse{
	// 	DataCollectionRuleResource: armmonitor.DataCollectionRuleResource{
	// 		Name: to.Ptr("myCollectionRule"),
	// 		Type: to.Ptr("Microsoft.Insights/dataCollectionRules"),
	// 		Etag: to.Ptr("070057da-0000-0000-0000-5ba70d6c0000"),
	// 		ID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionRules/myCollectionRule"),
	// 		Kind: to.Ptr(armmonitor.KnownDataCollectionRuleResourceKind("Direct")),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armmonitor.DataCollectionRuleResourceProperties{
	// 			Description: to.Ptr("A Direct Ingestion Rule with builtin ingestion fqdns"),
	// 			DataFlows: []*armmonitor.DataFlow{
	// 				{
	// 					Destinations: []*string{
	// 						to.Ptr("myworkspace"),
	// 					},
	// 					OutputStream: to.Ptr("Custom-LOGS1_CL"),
	// 					Streams: []*armmonitor.KnownDataFlowStreams{
	// 						to.Ptr(armmonitor.KnownDataFlowStreams("Custom-LOGS1_CL")),
	// 					},
	// 					TransformKql: to.Ptr("source | extend jsonContext = parse_json(AdditionalContext) | project TimeGenerated = Time, Computer, AdditionalContext = jsonContext, CounterName=tostring(jsonContext.CounterName), CounterValue=toreal(jsonContext.CounterValue)"),
	// 				},
	// 			},
	// 			Destinations: &armmonitor.DataCollectionRuleDestinations{
	// 				LogAnalytics: []*armmonitor.LogAnalyticsDestination{
	// 					{
	// 						Name: to.Ptr("centralWorkspace"),
	// 						WorkspaceResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/centralTeamWorkspace"),
	// 					},
	// 				},
	// 			},
	// 			Endpoints: &armmonitor.DataCollectionRuleEndpoints{
	// 				LogsIngestion: to.Ptr("https://mycollectionrule-8ykm-eastus2euap.logs.z1.canary.ingest.monitor.azure.com"),
	// 				MetricsIngestion: to.Ptr("https://mycollectionrule-jcvc-eastus2euap.metrics.z1.canary.ingest.monitor.azure.com"),
	// 			},
	// 			ImmutableID: to.Ptr("dcr-d2a09c11a66243009af059a655750000"),
	// 			ProvisioningState: to.Ptr(armmonitor.KnownDataCollectionRuleProvisioningStateSucceeded),
	// 			StreamDeclarations: map[string]*armmonitor.StreamDeclaration{
	// 				"Custom-LOGS1_CL": &armmonitor.StreamDeclaration{
	// 					Columns: []*armmonitor.ColumnDefinition{
	// 						{
	// 							Name: to.Ptr("Time"),
	// 							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeDatetime),
	// 						},
	// 						{
	// 							Name: to.Ptr("Computer"),
	// 							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeString),
	// 						},
	// 						{
	// 							Name: to.Ptr("AdditionalContext"),
	// 							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeString),
	// 						},
	// 						{
	// 							Name: to.Ptr("CounterName"),
	// 							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeString),
	// 						},
	// 						{
	// 							Name: to.Ptr("CounterValue"),
	// 							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeReal),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		SystemData: &armmonitor.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T17:50:40.5383301Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-01-30T17:50:40.5383301Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user1"),
	// 			LastModifiedByType: to.Ptr(armmonitor.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}

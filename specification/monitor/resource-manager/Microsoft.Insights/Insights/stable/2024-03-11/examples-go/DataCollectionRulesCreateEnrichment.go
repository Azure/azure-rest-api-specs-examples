package armmonitor_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
)

// Generated from example definition: 2024-03-11/DataCollectionRulesCreateEnrichment.json
func ExampleDataCollectionRulesClient_Create_createOrUpdateDataCollectionRuleWithEnrichment() {
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
		Location: to.Ptr("eastus"),
		Properties: &armmonitor.DataCollectionRuleResourceProperties{
			Description:              to.Ptr("A rule showcasing ingestion time enrichment"),
			DataCollectionEndpointID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionEndpoints/myDataCollectionEndpoint"),
			DataFlows: []*armmonitor.DataFlow{
				{
					Destinations: []*string{
						to.Ptr("centralWorkspace"),
					},
					OutputStream: to.Ptr("Custom-LOGS1_CL"),
					Streams: []*armmonitor.KnownDataFlowStreams{
						to.Ptr(armmonitor.KnownDataFlowStreams("Custom-TabularDataABC")),
					},
					TransformKql: to.Ptr("source | extend LookupData = lookup_string_am('mytextdatastore', Message) | project TimeGenerated, Message, AdditionalContext = LookupData.Message"),
				},
			},
			DataSources: &armmonitor.DataCollectionRuleDataSources{
				LogFiles: []*armmonitor.LogFilesDataSource{
					{
						Name:   to.Ptr("myTabularLogDataSource"),
						Format: to.Ptr(armmonitor.KnownLogFilesDataSourceFormatText),
						FilePatterns: []*string{
							to.Ptr("C:\\JavaLogs\\*\\*.log"),
						},
						Settings: &armmonitor.LogFilesDataSourceSettings{
							Text: &armmonitor.LogFileSettingsText{
								RecordStartTimestampFormat: to.Ptr(armmonitor.KnownLogFileTextSettingsRecordStartTimestampFormatISO8601),
							},
						},
						Streams: []*string{
							to.Ptr("Custom-TabularDataABC"),
						},
					},
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
			References: &armmonitor.DataCollectionRuleReferences{
				EnrichmentData: &armmonitor.ReferencesSpecEnrichmentData{
					StorageBlobs: []*armmonitor.StorageBlob{
						{
							Name:       to.Ptr("mytextdatastore"),
							BlobURL:    to.Ptr("https://myenrichmentstorage.blob.core.windows.net/enrichment"),
							LookupType: to.Ptr(armmonitor.KnownStorageBlobLookupTypeString),
							ResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourcegroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myenrichmentstorage"),
						},
					},
				},
			},
			StreamDeclarations: map[string]*armmonitor.StreamDeclaration{
				"Custom-TabularDataABC": {
					Columns: []*armmonitor.ColumnDefinition{
						{
							Name: to.Ptr("TimeGenerated"),
							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeDatetime),
						},
						{
							Name: to.Ptr("Message"),
							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeString),
						},
						{
							Name: to.Ptr("AdditionalContext"),
							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeString),
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
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armmonitor.DataCollectionRuleResourceProperties{
	// 			Description: to.Ptr("A rule showcasing ingestion time enrichment"),
	// 			DataCollectionEndpointID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.Insights/dataCollectionEndpoints/myDataCollectionEndpoint"),
	// 			DataFlows: []*armmonitor.DataFlow{
	// 				{
	// 					Destinations: []*string{
	// 						to.Ptr("centralWorkspace"),
	// 					},
	// 					OutputStream: to.Ptr("Custom-LOGS1_CL"),
	// 					Streams: []*armmonitor.KnownDataFlowStreams{
	// 						to.Ptr(armmonitor.KnownDataFlowStreams("Custom-TabularDataABC")),
	// 					},
	// 					TransformKql: to.Ptr("source | extend LookupData = lookup_string_am('mytextdatastore', Message) | project TimeGenerated, Message, AdditionalContext = LookupData.Message"),
	// 				},
	// 			},
	// 			DataSources: &armmonitor.DataCollectionRuleDataSources{
	// 				LogFiles: []*armmonitor.LogFilesDataSource{
	// 					{
	// 						Name: to.Ptr("myTabularLogDataSource"),
	// 						Format: to.Ptr(armmonitor.KnownLogFilesDataSourceFormatText),
	// 						FilePatterns: []*string{
	// 							to.Ptr("C:\\JavaLogs\\*\\*.log"),
	// 						},
	// 						Settings: &armmonitor.LogFilesDataSourceSettings{
	// 							Text: &armmonitor.LogFileSettingsText{
	// 								RecordStartTimestampFormat: to.Ptr(armmonitor.KnownLogFileTextSettingsRecordStartTimestampFormatISO8601),
	// 							},
	// 						},
	// 						Streams: []*string{
	// 							to.Ptr("Custom-TabularDataABC"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Destinations: &armmonitor.DataCollectionRuleDestinations{
	// 				LogAnalytics: []*armmonitor.LogAnalyticsDestination{
	// 					{
	// 						Name: to.Ptr("centralWorkspace"),
	// 						WorkspaceID: to.Ptr("9ba8bc53-bd36-4156-8667-e983e7ae0e4f"),
	// 						WorkspaceResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourceGroups/myResourceGroup/providers/Microsoft.OperationalInsights/workspaces/centralTeamWorkspace"),
	// 					},
	// 				},
	// 			},
	// 			ImmutableID: to.Ptr("dcr-ad96300ff0734d08a6a7195eb2be0000"),
	// 			ProvisioningState: to.Ptr(armmonitor.KnownDataCollectionRuleProvisioningStateSucceeded),
	// 			References: &armmonitor.DataCollectionRuleReferences{
	// 				EnrichmentData: &armmonitor.ReferencesSpecEnrichmentData{
	// 					StorageBlobs: []*armmonitor.StorageBlob{
	// 						{
	// 							Name: to.Ptr("mytextdatastore"),
	// 							BlobURL: to.Ptr("https://myenrichmentstorage.blob.core.windows.net/enrichment"),
	// 							LookupType: to.Ptr(armmonitor.KnownStorageBlobLookupTypeString),
	// 							ResourceID: to.Ptr("/subscriptions/703362b3-f278-4e4b-9179-c76eaf41ffc2/resourcegroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myenrichmentstorage"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			StreamDeclarations: map[string]*armmonitor.StreamDeclaration{
	// 				"Custom-TabularDataABC": &armmonitor.StreamDeclaration{
	// 					Columns: []*armmonitor.ColumnDefinition{
	// 						{
	// 							Name: to.Ptr("TimeGenerated"),
	// 							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeDatetime),
	// 						},
	// 						{
	// 							Name: to.Ptr("Message"),
	// 							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeString),
	// 						},
	// 						{
	// 							Name: to.Ptr("AdditionalContext"),
	// 							Type: to.Ptr(armmonitor.KnownColumnDefinitionTypeString),
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

package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ConnectorMappingsListByConnector.json
func ExampleConnectorMappingsClient_NewListByConnectorPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectorMappingsClient().NewListByConnectorPager("TestHubRG", "sdkTestHub", "testConnector8858", nil)
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
		// page.ConnectorMappingListResult = armcustomerinsights.ConnectorMappingListResult{
		// 	Value: []*armcustomerinsights.ConnectorMappingResourceFormat{
		// 		{
		// 			Name: to.Ptr("sdkTestHub/testConnector8858/testMapping12491"),
		// 			Type: to.Ptr("Microsoft.CustomerInsights/hubs/connectors/mappings"),
		// 			ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/sdkTestHub/connectors/testConnector8858/mappings/testMapping12491"),
		// 			Properties: &armcustomerinsights.ConnectorMapping{
		// 				Description: to.Ptr("Test mapping"),
		// 				ConnectorMappingName: to.Ptr("testMapping12491"),
		// 				ConnectorName: to.Ptr("testConnector8858"),
		// 				ConnectorType: to.Ptr(armcustomerinsights.ConnectorTypesAzureBlob),
		// 				DataFormatID: to.Ptr("4619d4e7e8a64c1e9bc34907f9f46505"),
		// 				DisplayName: to.Ptr("testMapping12491"),
		// 				EntityType: to.Ptr(armcustomerinsights.EntityTypesInteraction),
		// 				EntityTypeName: to.Ptr("TestInteractionType2967"),
		// 				MappingProperties: &armcustomerinsights.ConnectorMappingProperties{
		// 					Format: &armcustomerinsights.ConnectorMappingFormat{
		// 						AcceptLanguage: to.Ptr(""),
		// 						ArraySeparator: to.Ptr(";"),
		// 						ColumnDelimiter: to.Ptr("|"),
		// 						FormatType: to.Ptr("TextFormat"),
		// 						QuoteCharacter: to.Ptr("\\\""),
		// 						QuoteEscapeCharacter: to.Ptr("\\\""),
		// 					},
		// 					Availability: &armcustomerinsights.ConnectorMappingAvailability{
		// 						Frequency: to.Ptr(armcustomerinsights.FrequencyTypesHour),
		// 						Interval: to.Ptr[int32](5),
		// 					},
		// 					CompleteOperation: &armcustomerinsights.ConnectorMappingCompleteOperation{
		// 						CompletionOperationType: to.Ptr(armcustomerinsights.CompletionOperationTypesDeleteFile),
		// 						DestinationFolder: to.Ptr("fakePath"),
		// 					},
		// 					ErrorManagement: &armcustomerinsights.ConnectorMappingErrorManagement{
		// 						ErrorLimit: to.Ptr[int32](10),
		// 						ErrorManagementType: to.Ptr(armcustomerinsights.ErrorManagementTypesStopImport),
		// 					},
		// 					FileFilter: to.Ptr("unknown"),
		// 					FolderPath: to.Ptr("http://sample.dne/file"),
		// 					HasHeader: to.Ptr(false),
		// 					Structure: []*armcustomerinsights.ConnectorMappingStructure{
		// 						{
		// 							ColumnName: to.Ptr("unknown1"),
		// 							IsEncrypted: to.Ptr(false),
		// 							PropertyName: to.Ptr("unknwon1"),
		// 						},
		// 						{
		// 							ColumnName: to.Ptr("unknown2"),
		// 							IsEncrypted: to.Ptr(true),
		// 							PropertyName: to.Ptr("unknwon2"),
		// 					}},
		// 				},
		// 				State: to.Ptr(armcustomerinsights.ConnectorMappingStatesCreated),
		// 				TenantID: to.Ptr("sdktesthub"),
		// 			},
		// 	}},
		// }
	}
}

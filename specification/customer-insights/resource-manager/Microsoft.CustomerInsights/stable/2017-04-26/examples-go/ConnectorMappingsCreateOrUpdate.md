```go
package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ConnectorMappingsCreateOrUpdate.json
func ExampleConnectorMappingsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcustomerinsights.NewConnectorMappingsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"TestHubRG",
		"sdkTestHub",
		"testConnector8858",
		"testMapping12491",
		armcustomerinsights.ConnectorMappingResourceFormat{
			Properties: &armcustomerinsights.ConnectorMapping{
				Description:    to.Ptr("Test mapping"),
				DisplayName:    to.Ptr("testMapping12491"),
				EntityType:     to.Ptr(armcustomerinsights.EntityTypesInteraction),
				EntityTypeName: to.Ptr("TestInteractionType2967"),
				MappingProperties: &armcustomerinsights.ConnectorMappingProperties{
					Format: &armcustomerinsights.ConnectorMappingFormat{
						ColumnDelimiter: to.Ptr("|"),
						FormatType:      to.Ptr("TextFormat"),
					},
					Availability: &armcustomerinsights.ConnectorMappingAvailability{
						Frequency: to.Ptr(armcustomerinsights.FrequencyTypesHour),
						Interval:  to.Ptr[int32](5),
					},
					CompleteOperation: &armcustomerinsights.ConnectorMappingCompleteOperation{
						CompletionOperationType: to.Ptr(armcustomerinsights.CompletionOperationTypesDeleteFile),
						DestinationFolder:       to.Ptr("fakePath"),
					},
					ErrorManagement: &armcustomerinsights.ConnectorMappingErrorManagement{
						ErrorLimit:          to.Ptr[int32](10),
						ErrorManagementType: to.Ptr(armcustomerinsights.ErrorManagementTypesStopImport),
					},
					FileFilter: to.Ptr("unknown"),
					FolderPath: to.Ptr("http://sample.dne/file"),
					HasHeader:  to.Ptr(false),
					Structure: []*armcustomerinsights.ConnectorMappingStructure{
						{
							ColumnName:   to.Ptr("unknown1"),
							IsEncrypted:  to.Ptr(false),
							PropertyName: to.Ptr("unknwon1"),
						},
						{
							ColumnName:   to.Ptr("unknown2"),
							IsEncrypted:  to.Ptr(true),
							PropertyName: to.Ptr("unknwon2"),
						}},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcustomerinsights%2Farmcustomerinsights%2Fv1.0.0/sdk/resourcemanager/customerinsights/armcustomerinsights/README.md) on how to add the SDK to your project and authenticate.

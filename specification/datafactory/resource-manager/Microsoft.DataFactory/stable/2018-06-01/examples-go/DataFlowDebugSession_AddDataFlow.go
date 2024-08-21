package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/edb7904bfead536c7aa9716d44dba15bdabd0b00/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_AddDataFlow.json
func ExampleDataFlowDebugSessionClient_AddDataFlow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataFlowDebugSessionClient().AddDataFlow(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.DataFlowDebugPackage{
		DataFlow: &armdatafactory.DataFlowDebugResource{
			Name: to.Ptr("dataflow1"),
			Properties: &armdatafactory.MappingDataFlow{
				Type: to.Ptr("MappingDataFlow"),
				TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
					Script: to.Ptr("\n\nsource(output(\n		Column_1 as string\n	),\n	allowSchemaDrift: true,\n	validateSchema: false) ~> source1"),
					Sinks: []*armdatafactory.DataFlowSink{},
					Sources: []*armdatafactory.DataFlowSource{
						{
							Name: to.Ptr("source1"),
							Dataset: &armdatafactory.DatasetReference{
								Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
								ReferenceName: to.Ptr("DelimitedText2"),
							},
						}},
					Transformations: []*armdatafactory.Transformation{},
				},
			},
		},
		Datasets: []*armdatafactory.DatasetDebugResource{
			{
				Name: to.Ptr("dataset1"),
				Properties: &armdatafactory.DelimitedTextDataset{
					Type: to.Ptr("DelimitedText"),
					Schema: []any{
						map[string]any{
							"type": "String",
						},
					},
					Annotations: []any{},
					LinkedServiceName: &armdatafactory.LinkedServiceReference{
						Type:          to.Ptr(armdatafactory.LinkedServiceReferenceTypeLinkedServiceReference),
						ReferenceName: to.Ptr("linkedService5"),
					},
					TypeProperties: &armdatafactory.DelimitedTextDatasetTypeProperties{
						ColumnDelimiter:  ",",
						EscapeChar:       "\\",
						FirstRowAsHeader: true,
						Location: &armdatafactory.AzureBlobStorageLocation{
							Type:      to.Ptr("AzureBlobStorageLocation"),
							FileName:  "Ansiencoding.csv",
							Container: "dataflow-sample-data",
						},
						QuoteChar: "\"",
					},
				},
			}},
		DebugSettings: &armdatafactory.DataFlowDebugPackageDebugSettings{
			DatasetParameters: map[string]any{
				"Movies": map[string]any{
					"path": "abc",
				},
				"Output": map[string]any{
					"time": "def",
				},
			},
			Parameters: map[string]any{
				"sourcePath": "Toy",
			},
			SourceSettings: []*armdatafactory.DataFlowSourceSetting{
				{
					RowLimit:   to.Ptr[int32](1000),
					SourceName: to.Ptr("source1"),
				},
				{
					RowLimit:   to.Ptr[int32](222),
					SourceName: to.Ptr("source2"),
				}},
		},
		LinkedServices: []*armdatafactory.LinkedServiceDebugResource{
			{
				Name: to.Ptr("linkedService1"),
				Properties: &armdatafactory.AzureBlobStorageLinkedService{
					Type:        to.Ptr("AzureBlobStorage"),
					Annotations: []any{},
					TypeProperties: &armdatafactory.AzureBlobStorageLinkedServiceTypeProperties{
						ConnectionString:    "DefaultEndpointsProtocol=https;AccountName=<storageName>;EndpointSuffix=core.windows.net;",
						EncryptedCredential: to.Ptr("<credential>"),
					},
				},
			}},
		SessionID: to.Ptr("f06ed247-9d07-49b2-b05e-2cb4a2fc871e"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AddDataFlowToDebugSessionResponse = armdatafactory.AddDataFlowToDebugSessionResponse{
	// 	JobVersion: to.Ptr("e5328ee7-c524-4207-8ba4-b709010db33d"),
	// }
}

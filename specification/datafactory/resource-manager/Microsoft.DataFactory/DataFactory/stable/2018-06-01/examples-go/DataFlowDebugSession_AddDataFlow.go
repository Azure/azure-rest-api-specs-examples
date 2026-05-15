package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v11"
)

// Generated from example definition: 2018-06-01/DataFlowDebugSession_AddDataFlow.json
func ExampleDataFlowDebugSessionClient_AddDataFlow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataFlowDebugSessionClient().AddDataFlow(ctx, "exampleResourceGroup", "exampleFactoryName", armdatafactory.DataFlowDebugPackage{
		AdditionalProperties: map[string]any{
			"properties": map[string]any{
				"dataFlow": map[string]any{
					"name": "dataflow1",
					"properties": map[string]any{
						"type": "MappingDataFlow",
						"typeProperties": map[string]any{
							"script": "\n\nsource(output(\n		Column_1 as string\n	),\n	allowSchemaDrift: true,\n	validateSchema: false) ~> source1",
							"sinks":  []any{},
							"sources": []any{
								map[string]any{
									"name": "source1",
									"dataset": map[string]any{
										"type":          "DatasetReference",
										"referenceName": "DelimitedText2",
									},
								},
							},
							"transformations": []any{},
						},
					},
				},
				"datasets": []any{
					map[string]any{
						"name": "dataset1",
						"properties": map[string]any{
							"type": "DelimitedText",
							"schema": []any{
								map[string]any{
									"type": "String",
								},
							},
							"annotations": []any{},
							"linkedServiceName": map[string]any{
								"type":          "LinkedServiceReference",
								"referenceName": "linkedService5",
							},
							"typeProperties": map[string]any{
								"columnDelimiter":  ",",
								"escapeChar":       "\\",
								"firstRowAsHeader": true,
								"location": map[string]any{
									"type":      "AzureBlobStorageLocation",
									"container": "dataflow-sample-data",
									"fileName":  "Ansiencoding.csv",
								},
								"quoteChar": "\"",
							},
						},
					},
				},
				"debugSettings": map[string]any{
					"datasetParameters": map[string]any{
						"Movies": map[string]any{
							"path": "abc",
						},
						"Output": map[string]any{
							"time": "def",
						},
					},
					"parameters": map[string]any{
						"sourcePath": "Toy",
					},
					"sourceSettings": []any{
						map[string]any{
							"rowLimit":   1000,
							"sourceName": "source1",
						},
						map[string]any{
							"rowLimit":   222,
							"sourceName": "source2",
						},
					},
				},
				"linkedServices": []any{
					map[string]any{
						"name": "linkedService1",
						"properties": map[string]any{
							"type":        "AzureBlobStorage",
							"annotations": []any{},
							"typeProperties": map[string]any{
								"connectionString":    "DefaultEndpointsProtocol=https;AccountName=<storageName>;EndpointSuffix=core.windows.net;",
								"encryptedCredential": "<credential>",
							},
						},
					},
				},
				"sessionId": "f06ed247-9d07-49b2-b05e-2cb4a2fc871e",
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdatafactory.DataFlowDebugSessionClientAddDataFlowResponse{
	// 	AddDataFlowToDebugSessionResponse: &armdatafactory.AddDataFlowToDebugSessionResponse{
	// 		JobVersion: to.Ptr("e5328ee7-c524-4207-8ba4-b709010db33d"),
	// 	},
	// }
}

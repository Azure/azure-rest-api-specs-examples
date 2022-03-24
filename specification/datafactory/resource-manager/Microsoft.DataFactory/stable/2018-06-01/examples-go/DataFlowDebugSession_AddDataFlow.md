Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdatafactory%2Farmdatafactory%2Fv0.3.0/sdk/resourcemanager/datafactory/armdatafactory/README.md) on how to add the SDK to your project and authenticate.

```go
package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlowDebugSession_AddDataFlow.json
func ExampleDataFlowDebugSessionClient_AddDataFlow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewDataFlowDebugSessionClient("<subscription-id>", cred, nil)
	res, err := client.AddDataFlow(ctx,
		"<resource-group-name>",
		"<factory-name>",
		armdatafactory.DataFlowDebugPackage{
			DataFlow: &armdatafactory.DataFlowDebugResource{
				Name: to.StringPtr("<name>"),
				Properties: &armdatafactory.MappingDataFlow{
					Type: to.StringPtr("<type>"),
					TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
						Script: to.StringPtr("<script>"),
						Sinks:  []*armdatafactory.DataFlowSink{},
						Sources: []*armdatafactory.DataFlowSource{
							{
								Name: to.StringPtr("<name>"),
								Dataset: &armdatafactory.DatasetReference{
									Type:          armdatafactory.DatasetReferenceType("DatasetReference").ToPtr(),
									ReferenceName: to.StringPtr("<reference-name>"),
								},
							}},
						Transformations: []*armdatafactory.Transformation{},
					},
				},
			},
			Datasets: []*armdatafactory.DatasetDebugResource{
				{
					Name: to.StringPtr("<name>"),
					Properties: &armdatafactory.DelimitedTextDataset{
						Type: to.StringPtr("<type>"),
						Schema: []interface{}{
							map[string]interface{}{
								"type": "String",
							},
						},
						Annotations: []interface{}{},
						LinkedServiceName: &armdatafactory.LinkedServiceReference{
							Type:          armdatafactory.LinkedServiceReferenceType("LinkedServiceReference").ToPtr(),
							ReferenceName: to.StringPtr("<reference-name>"),
						},
						TypeProperties: &armdatafactory.DelimitedTextDatasetTypeProperties{
							ColumnDelimiter:  ",",
							EscapeChar:       "\\",
							FirstRowAsHeader: true,
							Location: &armdatafactory.AzureBlobStorageLocation{
								Type:      to.StringPtr("<type>"),
								FileName:  "Ansiencoding.csv",
								Container: "dataflow-sample-data",
							},
							QuoteChar: "\"",
						},
					},
				}},
			DebugSettings: &armdatafactory.DataFlowDebugPackageDebugSettings{
				DatasetParameters: map[string]interface{}{
					"Movies": map[string]interface{}{
						"path": "abc",
					},
					"Output": map[string]interface{}{
						"time": "def",
					},
				},
				Parameters: map[string]interface{}{
					"sourcePath": "Toy",
				},
				SourceSettings: []*armdatafactory.DataFlowSourceSetting{
					{
						RowLimit:   to.Int32Ptr(1000),
						SourceName: to.StringPtr("<source-name>"),
					},
					{
						RowLimit:   to.Int32Ptr(222),
						SourceName: to.StringPtr("<source-name>"),
					}},
			},
			LinkedServices: []*armdatafactory.LinkedServiceDebugResource{
				{
					Name: to.StringPtr("<name>"),
					Properties: &armdatafactory.AzureBlobStorageLinkedService{
						Type:        to.StringPtr("<type>"),
						Annotations: []interface{}{},
						TypeProperties: &armdatafactory.AzureBlobStorageLinkedServiceTypeProperties{
							ConnectionString:    "DefaultEndpointsProtocol=https;AccountName=<storageName>;EndpointSuffix=core.windows.net;",
							EncryptedCredential: to.StringPtr("<encrypted-credential>"),
						},
					},
				}},
			SessionID: to.StringPtr("<session-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DataFlowDebugSessionClientAddDataFlowResult)
}
```

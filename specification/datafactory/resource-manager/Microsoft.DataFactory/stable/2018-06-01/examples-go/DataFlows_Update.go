package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6051d2b126f5b1e4b623cde8edfa3e25cf730685/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Update.json
func ExampleDataFlowsClient_CreateOrUpdate_dataFlowsUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataFlowsClient().CreateOrUpdate(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataFlow", armdatafactory.DataFlowResource{
		Properties: &armdatafactory.MappingDataFlow{
			Type:        to.Ptr("MappingDataFlow"),
			Description: to.Ptr("Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation."),
			TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
				ScriptLines: []*string{
					to.Ptr("source(output("),
					to.Ptr("PreviousConversionRate as double,"),
					to.Ptr("Country as string,"),
					to.Ptr("DateTime1 as string,"),
					to.Ptr("CurrentConversionRate as double"),
					to.Ptr("),"),
					to.Ptr("allowSchemaDrift: false,"),
					to.Ptr("validateSchema: false) ~> USDCurrency"),
					to.Ptr("source(output("),
					to.Ptr("PreviousConversionRate as double,"),
					to.Ptr("Country as string,"),
					to.Ptr("DateTime1 as string,"),
					to.Ptr("CurrentConversionRate as double"),
					to.Ptr("),"),
					to.Ptr("allowSchemaDrift: true,"),
					to.Ptr("validateSchema: false) ~> CADSource"),
					to.Ptr("USDCurrency, CADSource union(byName: true)~> Union"),
					to.Ptr("Union derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn"),
					to.Ptr("NewCurrencyColumn split(Country == 'USD',"),
					to.Ptr("Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)"),
					to.Ptr("ConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink"),
					to.Ptr("ConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink")},
				Sinks: []*armdatafactory.DataFlowSink{
					{
						Name: to.Ptr("USDSink"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("USDOutput"),
						},
					},
					{
						Name: to.Ptr("CADSink"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("CADOutput"),
						},
					}},
				Sources: []*armdatafactory.DataFlowSource{
					{
						Name: to.Ptr("USDCurrency"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("CurrencyDatasetUSD"),
						},
					},
					{
						Name: to.Ptr("CADSource"),
						Dataset: &armdatafactory.DatasetReference{
							Type:          to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
							ReferenceName: to.Ptr("CurrencyDatasetCAD"),
						},
					}},
			},
		},
	}, &armdatafactory.DataFlowsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataFlowResource = armdatafactory.DataFlowResource{
	// 	Name: to.Ptr("exampleDataFlow"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/dataflows"),
	// 	Etag: to.Ptr("0a0068d4-0000-0000-0000-5b245bd30002"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/datasets/exampleDataset"),
	// 	Properties: &armdatafactory.MappingDataFlow{
	// 		Type: to.Ptr("MappingDataFlow"),
	// 		Description: to.Ptr("Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation."),
	// 		TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
	// 			ScriptLines: []*string{
	// 				to.Ptr("source(output("),
	// 				to.Ptr("PreviousConversionRate as double,"),
	// 				to.Ptr("Country as string,"),
	// 				to.Ptr("DateTime1 as string,"),
	// 				to.Ptr("CurrentConversionRate as double"),
	// 				to.Ptr("),"),
	// 				to.Ptr("allowSchemaDrift: false,"),
	// 				to.Ptr("validateSchema: false) ~> USDCurrency"),
	// 				to.Ptr("source(output("),
	// 				to.Ptr("PreviousConversionRate as double,"),
	// 				to.Ptr("Country as string,"),
	// 				to.Ptr("DateTime1 as string,"),
	// 				to.Ptr("CurrentConversionRate as double"),
	// 				to.Ptr("),"),
	// 				to.Ptr("allowSchemaDrift: true,"),
	// 				to.Ptr("validateSchema: false) ~> CADSource"),
	// 				to.Ptr("USDCurrency, CADSource union(byName: true)~> Union"),
	// 				to.Ptr("Union derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn"),
	// 				to.Ptr("NewCurrencyColumn split(Country == 'USD',"),
	// 				to.Ptr("Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)"),
	// 				to.Ptr("ConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink"),
	// 				to.Ptr("ConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink")},
	// 				Sinks: []*armdatafactory.DataFlowSink{
	// 					{
	// 						Name: to.Ptr("USDSink"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("USDOutput"),
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("CADSink"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("CADOutput"),
	// 						},
	// 				}},
	// 				Sources: []*armdatafactory.DataFlowSource{
	// 					{
	// 						Name: to.Ptr("USDCurrency"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("CurrencyDatasetUSD"),
	// 						},
	// 					},
	// 					{
	// 						Name: to.Ptr("CADSource"),
	// 						Dataset: &armdatafactory.DatasetReference{
	// 							Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 							ReferenceName: to.Ptr("CurrencyDatasetCAD"),
	// 						},
	// 				}},
	// 			},
	// 		},
	// 	}
}

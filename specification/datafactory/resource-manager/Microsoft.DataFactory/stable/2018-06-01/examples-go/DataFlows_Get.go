package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8eb3f7a4f66d408152c32b9d647e59147172d533/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Get.json
func ExampleDataFlowsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataFlowsClient().Get(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleDataFlow", &armdatafactory.DataFlowsClientGetOptions{IfNoneMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataFlowResource = armdatafactory.DataFlowResource{
	// 	Name: to.Ptr("exampleDataFlow"),
	// 	Type: to.Ptr("Microsoft.DataFactory/factories/dataflows"),
	// 	Etag: to.Ptr("15004c4f-0000-0200-0000-5cbe090e0000"),
	// 	ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/dataflows/exampleDataFlow"),
	// 	Properties: &armdatafactory.MappingDataFlow{
	// 		Type: to.Ptr("MappingDataFlow"),
	// 		Description: to.Ptr("Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation."),
	// 		TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
	// 			Script: to.Ptr("source(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: false,validateSchema: false) ~> USDCurrency\nsource(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: true,validateSchema: false) ~> CADSource\nUSDCurrency, CADSource union(byName: true)~> Union\nUnion derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn\nNewCurrencyColumn split(Country == 'USD',Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)\nConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink\nConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink"),
	// 			Sinks: []*armdatafactory.DataFlowSink{
	// 				{
	// 					Name: to.Ptr("USDSink"),
	// 					Dataset: &armdatafactory.DatasetReference{
	// 						Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 						ReferenceName: to.Ptr("USDOutput"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("CADSink"),
	// 					Dataset: &armdatafactory.DatasetReference{
	// 						Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 						ReferenceName: to.Ptr("CADOutput"),
	// 					},
	// 			}},
	// 			Sources: []*armdatafactory.DataFlowSource{
	// 				{
	// 					Name: to.Ptr("USDCurrency"),
	// 					Dataset: &armdatafactory.DatasetReference{
	// 						Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 						ReferenceName: to.Ptr("CurrencyDatasetUSD"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("CADSource"),
	// 					Dataset: &armdatafactory.DatasetReference{
	// 						Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
	// 						ReferenceName: to.Ptr("CurrencyDatasetCAD"),
	// 					},
	// 			}},
	// 		},
	// 	},
	// }
}

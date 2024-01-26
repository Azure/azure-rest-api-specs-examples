package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/be0d1c383d683a8eb22ab5fd5b75e380ac3c2eea/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_ListByFactory.json
func ExampleDataFlowsClient_NewListByFactoryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataFlowsClient().NewListByFactoryPager("exampleResourceGroup", "exampleFactoryName", nil)
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
		// page.DataFlowListResponse = armdatafactory.DataFlowListResponse{
		// 	Value: []*armdatafactory.DataFlowResource{
		// 		{
		// 			Name: to.Ptr("exampleDataFlow"),
		// 			Type: to.Ptr("Microsoft.DataFactory/factories/dataflows"),
		// 			Etag: to.Ptr("0a0068d4-0000-0000-0000-5b245bd30000"),
		// 			ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.DataFactory/factories/exampleFactoryName/dataflows/exampleDataFlow"),
		// 			Properties: &armdatafactory.MappingDataFlow{
		// 				Type: to.Ptr("MappingDataFlow"),
		// 				Description: to.Ptr("Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation."),
		// 				TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
		// 					Script: to.Ptr("source(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: false,validateSchema: false) ~> USDCurrency\nsource(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: true,validateSchema: false) ~> CADSource\nUSDCurrency, CADSource union(byName: true)~> Union\nUnion derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn\nNewCurrencyColumn split(Country == 'USD',Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)\nConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink\nConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink"),
		// 					Sinks: []*armdatafactory.DataFlowSink{
		// 						{
		// 							Name: to.Ptr("USDSink"),
		// 							Dataset: &armdatafactory.DatasetReference{
		// 								Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
		// 								ReferenceName: to.Ptr("USDOutput"),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("CADSink"),
		// 							Dataset: &armdatafactory.DatasetReference{
		// 								Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
		// 								ReferenceName: to.Ptr("CADOutput"),
		// 							},
		// 					}},
		// 					Sources: []*armdatafactory.DataFlowSource{
		// 						{
		// 							Name: to.Ptr("USDCurrency"),
		// 							Dataset: &armdatafactory.DatasetReference{
		// 								Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
		// 								ReferenceName: to.Ptr("CurrencyDatasetUSD"),
		// 							},
		// 						},
		// 						{
		// 							Name: to.Ptr("CADSource"),
		// 							Dataset: &armdatafactory.DatasetReference{
		// 								Type: to.Ptr(armdatafactory.DatasetReferenceTypeDatasetReference),
		// 								ReferenceName: to.Ptr("CurrencyDatasetCAD"),
		// 							},
		// 					}},
		// 				},
		// 			},
		// 	}},
		// }
	}
}

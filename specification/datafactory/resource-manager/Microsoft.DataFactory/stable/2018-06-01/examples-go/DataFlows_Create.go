package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Create.json
func ExampleDataFlowsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdatafactory.NewDataFlowsClient("12345678-1234-1234-1234-12345678abc", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"exampleResourceGroup",
		"exampleFactoryName",
		"exampleDataFlow",
		armdatafactory.DataFlowResource{
			Properties: &armdatafactory.MappingDataFlow{
				Type:        to.Ptr("MappingDataFlow"),
				Description: to.Ptr("Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation."),
				TypeProperties: &armdatafactory.MappingDataFlowTypeProperties{
					Script: to.Ptr("source(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: false,validateSchema: false) ~> USDCurrency\nsource(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: true,validateSchema: false) ~> CADSource\nUSDCurrency, CADSource union(byName: true)~> Union\nUnion derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn\nNewCurrencyColumn split(Country == 'USD',Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)\nConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink\nConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink"),
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
		},
		&armdatafactory.DataFlowsClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

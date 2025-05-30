using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Update.json
// this example is just showing the usage of "DataFlows_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryResource created on azure
// for more information of creating DataFactoryResource, please refer to the document of DataFactoryResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
ResourceIdentifier dataFactoryResourceId = DataFactoryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName);
DataFactoryResource dataFactory = client.GetDataFactoryResource(dataFactoryResourceId);

// get the collection of this DataFactoryDataFlowResource
DataFactoryDataFlowCollection collection = dataFactory.GetDataFactoryDataFlows();

// invoke the operation
string dataFlowName = "exampleDataFlow";
DataFactoryDataFlowData data = new DataFactoryDataFlowData(new DataFactoryMappingDataFlowProperties
{
    Sources = {new DataFlowSource("USDCurrency")
    {
    Dataset = new DatasetReference(DatasetReferenceType.DatasetReference, "CurrencyDatasetUSD"),
    }, new DataFlowSource("CADSource")
    {
    Dataset = new DatasetReference(DatasetReferenceType.DatasetReference, "CurrencyDatasetCAD"),
    }},
    Sinks = {new DataFlowSink("USDSink")
    {
    Dataset = new DatasetReference(DatasetReferenceType.DatasetReference, "USDOutput"),
    }, new DataFlowSink("CADSink")
    {
    Dataset = new DatasetReference(DatasetReferenceType.DatasetReference, "CADOutput"),
    }},
    ScriptLines = { "source(output(", "PreviousConversionRate as double,", "Country as string,", "DateTime1 as string,", "CurrentConversionRate as double", "),", "allowSchemaDrift: false,", "validateSchema: false) ~> USDCurrency", "source(output(", "PreviousConversionRate as double,", "Country as string,", "DateTime1 as string,", "CurrentConversionRate as double", "),", "allowSchemaDrift: true,", "validateSchema: false) ~> CADSource", "USDCurrency, CADSource union(byName: true)~> Union", "Union derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn", "NewCurrencyColumn split(Country == 'USD',", "Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)", "ConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink", "ConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink" },
    Description = "Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation.",
});
ArmOperation<DataFactoryDataFlowResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataFlowName, data);
DataFactoryDataFlowResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryDataFlowData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

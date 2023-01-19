using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataFactory;
using Azure.ResourceManager.DataFactory.Models;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Update.json
// this example is just showing the usage of "DataFlows_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FactoryDataFlowResource created on azure
// for more information of creating FactoryDataFlowResource, please refer to the document of FactoryDataFlowResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string dataFlowName = "exampleDataFlow";
ResourceIdentifier factoryDataFlowResourceId = FactoryDataFlowResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, dataFlowName);
FactoryDataFlowResource factoryDataFlow = client.GetFactoryDataFlowResource(factoryDataFlowResourceId);

// invoke the operation
FactoryDataFlowData data = new FactoryDataFlowData(new FactoryMappingDataFlowDefinition()
{
    Sources =
    {
    new DataFlowSource("USDCurrency")
    {
    Dataset = new DatasetReference(DatasetReferenceType.DatasetReference,"CurrencyDatasetUSD"),
    },new DataFlowSource("CADSource")
    {
    Dataset = new DatasetReference(DatasetReferenceType.DatasetReference,"CurrencyDatasetCAD"),
    }
    },
    Sinks =
    {
    new DataFlowSink("USDSink")
    {
    Dataset = new DatasetReference(DatasetReferenceType.DatasetReference,"USDOutput"),
    },new DataFlowSink("CADSink")
    {
    Dataset = new DatasetReference(DatasetReferenceType.DatasetReference,"CADOutput"),
    }
    },
    Script = "source(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: false,validateSchema: false) ~> USDCurrency\nsource(output(PreviousConversionRate as double,Country as string,DateTime1 as string,CurrentConversionRate as double),allowSchemaDrift: true,validateSchema: false) ~> CADSource\nUSDCurrency, CADSource union(byName: true)~> Union\nUnion derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn\nNewCurrencyColumn split(Country == 'USD',Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)\nConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink\nConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink",
    Description = "Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation.",
});
ArmOperation<FactoryDataFlowResource> lro = await factoryDataFlow.UpdateAsync(WaitUntil.Completed, data);
FactoryDataFlowResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FactoryDataFlowData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

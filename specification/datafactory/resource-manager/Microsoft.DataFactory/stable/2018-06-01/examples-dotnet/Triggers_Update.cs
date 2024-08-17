using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_Update.json
// this example is just showing the usage of "Triggers_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this DataFactoryTriggerResource
DataFactoryTriggerCollection collection = dataFactory.GetDataFactoryTriggers();

// invoke the operation
string triggerName = "exampleTrigger";
DataFactoryTriggerData data = new DataFactoryTriggerData(new DataFactoryScheduleTrigger(new ScheduleTriggerRecurrence()
{
    Frequency = DataFactoryRecurrenceFrequency.Minute,
    Interval = 4,
    StartOn = DateTimeOffset.Parse("2018-06-16T00:39:14.905167Z"),
    EndOn = DateTimeOffset.Parse("2018-06-16T00:55:14.905167Z"),
    TimeZone = "UTC",
})
{
    Pipelines =
    {
    new TriggerPipelineReference()
    {
    PipelineReference = new DataFactoryPipelineReference(DataFactoryPipelineReferenceType.PipelineReference,"examplePipeline"),
    Parameters =
    {
    ["OutputBlobNameList"] = BinaryData.FromObjectAsJson(new object[] { "exampleoutput.csv" }),
    },
    }
    },
    Description = "Example description",
});
ArmOperation<DataFactoryTriggerResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, triggerName, data);
DataFactoryTriggerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryTriggerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

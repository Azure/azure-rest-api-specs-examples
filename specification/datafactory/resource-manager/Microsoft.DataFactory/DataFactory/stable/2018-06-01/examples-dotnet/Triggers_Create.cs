using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataFactory.Models;
using Azure.ResourceManager.DataFactory;

// Generated from example definition: specification/datafactory/resource-manager/Microsoft.DataFactory/DataFactory/stable/2018-06-01/examples/Triggers_Create.json
// this example is just showing the usage of "Triggers_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataFactoryTriggerResource created on azure
// for more information of creating DataFactoryTriggerResource, please refer to the document of DataFactoryTriggerResource
string subscriptionId = "12345678-1234-1234-1234-12345678abc";
string resourceGroupName = "exampleResourceGroup";
string factoryName = "exampleFactoryName";
string triggerName = "exampleTrigger";
ResourceIdentifier dataFactoryTriggerResourceId = DataFactoryTriggerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, factoryName, triggerName);
DataFactoryTriggerResource dataFactoryTrigger = client.GetDataFactoryTriggerResource(dataFactoryTriggerResourceId);

// invoke the operation
DataFactoryTriggerData data = new DataFactoryTriggerData(new DataFactoryScheduleTrigger(new ScheduleTriggerRecurrence
{
    Frequency = DataFactoryRecurrenceFrequency.Minute,
    Interval = 4,
    StartOn = DateTimeOffset.Parse("2018-06-16T00:39:13.8441801Z"),
    EndOn = DateTimeOffset.Parse("2018-06-16T00:55:13.8441801Z"),
    TimeZone = "UTC",
})
{
    Pipelines = {new TriggerPipelineReference
    {
    PipelineReference = new DataFactoryPipelineReference(DataFactoryPipelineReferenceType.PipelineReference, "examplePipeline"),
    Parameters =
    {
    ["OutputBlobNameList"] = BinaryData.FromObjectAsJson(new object[]
    {
    "exampleoutput.csv"
    })
    },
    }},
});
ArmOperation<DataFactoryTriggerResource> lro = await dataFactoryTrigger.UpdateAsync(WaitUntil.Completed, data);
DataFactoryTriggerResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataFactoryTriggerData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

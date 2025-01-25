using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2024-11-01/Dataflow_CreateOrUpdate_ComplexContextualization.json
// this example is just showing the usage of "DataflowResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsDataflowProfileResource created on azure
// for more information of creating IotOperationsDataflowProfileResource, please refer to the document of IotOperationsDataflowProfileResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-name123";
string dataflowProfileName = "resource-name123";
ResourceIdentifier iotOperationsDataflowProfileResourceId = IotOperationsDataflowProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName, dataflowProfileName);
IotOperationsDataflowProfileResource iotOperationsDataflowProfile = client.GetIotOperationsDataflowProfileResource(iotOperationsDataflowProfileResourceId);

// get the collection of this IotOperationsDataflowResource
IotOperationsDataflowCollection collection = iotOperationsDataflowProfile.GetIotOperationsDataflows();

// invoke the operation
string dataflowName = "aio-to-adx-contexualized";
IotOperationsDataflowData data = new IotOperationsDataflowData(new IotOperationsExtendedLocation("qmbrfwcpwwhggszhrdjv", IotOperationsExtendedLocationType.CustomLocation))
{
    Properties = new IotOperationsDataflowProperties(new DataflowOperationProperties[]
{
new DataflowOperationProperties(DataflowOperationType.Source)
{
Name = "source1",
SourceSettings = new DataflowSourceOperationSettings("aio-builtin-broker-endpoint", new string[]{"azure-iot-operations/data/thermostat"}),
},
new DataflowOperationProperties(DataflowOperationType.BuiltInTransformation)
{
Name = "transformation1",
BuiltInTransformationSettings = new DataflowBuiltInTransformationSettings
{
Datasets = {new DataflowBuiltInTransformationDataset("quality", new string[]{"$source.country", "$context.country"})
{
Expression = "$1 == $2",
}},
Map = {new DataflowBuiltInTransformationMap(new string[]{"*"}, "*"), new DataflowBuiltInTransformationMap(new string[]{"$context(quality).*"}, "enriched.*")},
},
},
new DataflowOperationProperties(DataflowOperationType.Destination)
{
Name = "destination1",
DestinationSettings = new DataflowDestinationOperationSettings("adx-endpoint", "mytable"),
}
})
    {
        Mode = IotOperationsOperationalMode.Enabled,
    },
};
ArmOperation<IotOperationsDataflowResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dataflowName, data);
IotOperationsDataflowResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotOperationsDataflowData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

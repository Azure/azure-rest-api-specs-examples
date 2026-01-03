using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/Dataflow_CreateOrUpdate_SimpleFabric.json
// this example is just showing the usage of "DataflowResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsDataflowResource created on azure
// for more information of creating IotOperationsDataflowResource, please refer to the document of IotOperationsDataflowResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-name123";
string dataflowProfileName = "resource-name123";
string dataflowName = "aio-to-fabric";
ResourceIdentifier iotOperationsDataflowResourceId = IotOperationsDataflowResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName, dataflowProfileName, dataflowName);
IotOperationsDataflowResource iotOperationsDataflow = client.GetIotOperationsDataflowResource(iotOperationsDataflowResourceId);

// invoke the operation
IotOperationsDataflowData data = new IotOperationsDataflowData
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
BuiltInTransformationSettings = new DataflowBuiltInTransformationSettings
{
SerializationFormat = TransformationSerializationFormat.Parquet,
SchemaRef = "aio-sr://exampleNamespace/exmapleParquetSchema:1.0.0",
},
},
new DataflowOperationProperties(DataflowOperationType.Destination)
{
Name = "destination1",
DestinationSettings = new DataflowDestinationOperationSettings("fabric-endpoint", "telemetryTable"),
}
})
    {
        Mode = IotOperationsOperationalMode.Enabled,
    },
    ExtendedLocation = new IotOperationsExtendedLocation("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123", IotOperationsExtendedLocationType.CustomLocation),
};
ArmOperation<IotOperationsDataflowResource> lro = await iotOperationsDataflow.UpdateAsync(WaitUntil.Completed, data);
IotOperationsDataflowResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotOperationsDataflowData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

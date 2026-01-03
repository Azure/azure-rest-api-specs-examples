using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/DataflowEndpoint_CreateOrUpdate_Fabric.json
// this example is just showing the usage of "DataflowEndpointResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsDataflowEndpointResource created on azure
// for more information of creating IotOperationsDataflowEndpointResource, please refer to the document of IotOperationsDataflowEndpointResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-name123";
string dataflowEndpointName = "fabric-endpoint";
ResourceIdentifier iotOperationsDataflowEndpointResourceId = IotOperationsDataflowEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName, dataflowEndpointName);
IotOperationsDataflowEndpointResource iotOperationsDataflowEndpoint = client.GetIotOperationsDataflowEndpointResource(iotOperationsDataflowEndpointResourceId);

// invoke the operation
IotOperationsDataflowEndpointData data = new IotOperationsDataflowEndpointData
{
    Properties = new IotOperationsDataflowEndpointProperties(DataflowEndpointType.FabricOneLake)
    {
        FabricOneLakeSettings = new DataflowEndpointFabricOneLake(new DataflowEndpointFabricOneLakeAuthentication(FabricOneLakeAuthMethod.SystemAssignedManagedIdentity), new DataflowEndpointFabricOneLakeNames("example-lakehouse", "example-workspace"), DataflowEndpointFabricPathType.Tables, "onelake.dfs.fabric.microsoft.com"),
    },
    ExtendedLocation = new IotOperationsExtendedLocation("/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123", IotOperationsExtendedLocationType.CustomLocation),
};
ArmOperation<IotOperationsDataflowEndpointResource> lro = await iotOperationsDataflowEndpoint.UpdateAsync(WaitUntil.Completed, data);
IotOperationsDataflowEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotOperationsDataflowEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

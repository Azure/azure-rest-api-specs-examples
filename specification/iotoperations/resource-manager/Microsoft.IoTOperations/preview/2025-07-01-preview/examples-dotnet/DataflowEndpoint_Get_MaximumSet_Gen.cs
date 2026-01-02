using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/DataflowEndpoint_Get_MaximumSet_Gen.json
// this example is just showing the usage of "DataflowEndpointResource_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsDataflowEndpointResource created on azure
// for more information of creating IotOperationsDataflowEndpointResource, please refer to the document of IotOperationsDataflowEndpointResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-name123";
string dataflowEndpointName = "resource-name123";
ResourceIdentifier iotOperationsDataflowEndpointResourceId = IotOperationsDataflowEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName, dataflowEndpointName);
IotOperationsDataflowEndpointResource iotOperationsDataflowEndpoint = client.GetIotOperationsDataflowEndpointResource(iotOperationsDataflowEndpointResourceId);

// invoke the operation
IotOperationsDataflowEndpointResource result = await iotOperationsDataflowEndpoint.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotOperationsDataflowEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

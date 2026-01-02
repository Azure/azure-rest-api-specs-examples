using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotOperations.Models;
using Azure.ResourceManager.IotOperations;

// Generated from example definition: 2025-07-01-preview/RegistryEndpoint_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "RegistryEndpointResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotOperationsRegistryEndpointResource created on azure
// for more information of creating IotOperationsRegistryEndpointResource, please refer to the document of IotOperationsRegistryEndpointResource
string subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
string resourceGroupName = "rgiotoperations";
string instanceName = "resource-123";
string registryEndpointName = "resource-123";
ResourceIdentifier iotOperationsRegistryEndpointResourceId = IotOperationsRegistryEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, instanceName, registryEndpointName);
IotOperationsRegistryEndpointResource iotOperationsRegistryEndpoint = client.GetIotOperationsRegistryEndpointResource(iotOperationsRegistryEndpointResourceId);

// invoke the operation
await iotOperationsRegistryEndpoint.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HealthDataAIServices.Models;
using Azure.ResourceManager.HealthDataAIServices;

// Generated from example definition: 2024-09-20/PrivateEndpointConnections_ListByDeidService_MaximumSet_Gen.json
// this example is just showing the usage of "PrivateEndpointConnectionResource_ListByDeidService" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeidServiceResource created on azure
// for more information of creating DeidServiceResource, please refer to the document of DeidServiceResource
string subscriptionId = "F21BB31B-C214-42C0-ACF0-DACCA05D3011";
string resourceGroupName = "rgopenapi";
string deidServiceName = "deidTest";
ResourceIdentifier deidServiceResourceId = DeidServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deidServiceName);
DeidServiceResource deidService = client.GetDeidServiceResource(deidServiceResourceId);

// get the collection of this HealthDataAIServicesPrivateEndpointConnectionResource
HealthDataAIServicesPrivateEndpointConnectionResourceCollection collection = deidService.GetHealthDataAIServicesPrivateEndpointConnectionResources();

// invoke the operation and iterate over the result
await foreach (HealthDataAIServicesPrivateEndpointConnectionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HealthDataAIServicesPrivateEndpointConnectionResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

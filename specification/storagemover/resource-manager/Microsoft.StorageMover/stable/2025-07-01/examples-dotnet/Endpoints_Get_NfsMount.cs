using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageMover.Models;
using Azure.ResourceManager.StorageMover;

// Generated from example definition: 2025-07-01/Endpoints_Get_NfsMount.json
// this example is just showing the usage of "Endpoint_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StorageMoverResource created on azure
// for more information of creating StorageMoverResource, please refer to the document of StorageMoverResource
string subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
string resourceGroupName = "examples-rg";
string storageMoverName = "examples-storageMoverName";
ResourceIdentifier storageMoverResourceId = StorageMoverResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageMoverName);
StorageMoverResource storageMover = client.GetStorageMoverResource(storageMoverResourceId);

// get the collection of this StorageMoverEndpointResource
StorageMoverEndpointCollection collection = storageMover.GetStorageMoverEndpoints();

// invoke the operation
string endpointName = "examples-endpointName";
NullableResponse<StorageMoverEndpointResource> response = await collection.GetIfExistsAsync(endpointName);
StorageMoverEndpointResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    StorageMoverEndpointData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

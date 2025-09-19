using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/EnclaveEndpoints_ListByEnclaveResource.json
// this example is just showing the usage of "EnclaveEndpointResource_ListByEnclaveResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEnclaveResource created on azure
// for more information of creating VirtualEnclaveResource, please refer to the document of VirtualEnclaveResource
string subscriptionId = "73CEECEF-2C30-488E-946F-D20F414D99BA";
string resourceGroupName = "rgopenapi";
string virtualEnclaveName = "TestMyEnclave";
ResourceIdentifier virtualEnclaveResourceId = VirtualEnclaveResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualEnclaveName);
VirtualEnclaveResource virtualEnclave = client.GetVirtualEnclaveResource(virtualEnclaveResourceId);

// get the collection of this VirtualEnclaveEndpointResource
VirtualEnclaveEndpointCollection collection = virtualEnclave.GetVirtualEnclaveEndpoints();

// invoke the operation and iterate over the result
await foreach (VirtualEnclaveEndpointResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    VirtualEnclaveEndpointData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

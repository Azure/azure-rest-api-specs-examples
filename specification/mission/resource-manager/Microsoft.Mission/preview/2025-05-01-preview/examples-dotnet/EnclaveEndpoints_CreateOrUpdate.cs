using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/EnclaveEndpoints_CreateOrUpdate.json
// this example is just showing the usage of "EnclaveEndpointResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
string enclaveEndpointName = "TestMyEnclaveEndpoint";
VirtualEnclaveEndpointData data = new VirtualEnclaveEndpointData(new AzureLocation("West US"))
{
    Properties = new VirtualEnclaveEndpointProperties(new EnclaveEndpointDestinationRule[]
{
new EnclaveEndpointDestinationRule
{
Protocols = {EnclaveEndpointProtocol.Tcp},
EndpointRuleName = "54CEECEF-2C30-488E-946F-D20F414D99BA",
Destination = "10.0.0.0/24",
Ports = "443",
}
}),
    Tags =
    {
    ["sampletag"] = "samplevalue"
    },
};
ArmOperation<VirtualEnclaveEndpointResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, enclaveEndpointName, data);
VirtualEnclaveEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualEnclaveEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

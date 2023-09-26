using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/SubnetGet.json
// this example is just showing the usage of "Subnets_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualNetworkResource created on azure
// for more information of creating VirtualNetworkResource, please refer to the document of VirtualNetworkResource
string subscriptionId = "subid";
string resourceGroupName = "subnet-test";
string virtualNetworkName = "vnetname";
ResourceIdentifier virtualNetworkResourceId = VirtualNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkName);
VirtualNetworkResource virtualNetwork = client.GetVirtualNetworkResource(virtualNetworkResourceId);

// get the collection of this SubnetResource
SubnetCollection collection = virtualNetwork.GetSubnets();

// invoke the operation
string subnetName = "subnet1";
NullableResponse<SubnetResource> response = await collection.GetIfExistsAsync(subnetName);
SubnetResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SubnetData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

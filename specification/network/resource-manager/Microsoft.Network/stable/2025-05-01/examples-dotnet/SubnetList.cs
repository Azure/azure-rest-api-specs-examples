using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/SubnetList.json
// this example is just showing the usage of "Subnets_List" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
await foreach (SubnetResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SubnetData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

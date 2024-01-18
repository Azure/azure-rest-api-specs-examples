using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/VirtualNetworkGetResourceNavigationLinks.json
// this example is just showing the usage of "ResourceNavigationLinks_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubnetResource created on azure
// for more information of creating SubnetResource, please refer to the document of SubnetResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualNetworkName = "vnet";
string subnetName = "subnet";
ResourceIdentifier subnetResourceId = SubnetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkName, subnetName);
SubnetResource subnet = client.GetSubnetResource(subnetResourceId);

// invoke the operation and iterate over the result
await foreach (ResourceNavigationLink item in subnet.GetResourceNavigationLinksAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");

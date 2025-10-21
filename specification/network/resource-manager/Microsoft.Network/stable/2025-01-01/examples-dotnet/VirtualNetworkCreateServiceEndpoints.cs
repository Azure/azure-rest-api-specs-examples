using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-01-01/examples/VirtualNetworkCreateServiceEndpoints.json
// this example is just showing the usage of "VirtualNetworks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "vnetTest";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this VirtualNetworkResource
VirtualNetworkCollection collection = resourceGroupResource.GetVirtualNetworks();

// invoke the operation
string virtualNetworkName = "vnet1";
VirtualNetworkData data = new VirtualNetworkData
{
    AddressSpace = new VirtualNetworkAddressSpace
    {
        AddressPrefixes = { "10.0.0.0/16" },
    },
    Subnets = {new SubnetData
    {
    AddressPrefix = "10.0.0.0/16",
    ServiceEndpoints = {new ServiceEndpointProperties
    {
    Service = "Microsoft.Storage",
    }},
    Name = "test-1",
    }},
    Location = new AzureLocation("eastus"),
};
ArmOperation<VirtualNetworkResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, virtualNetworkName, data);
VirtualNetworkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

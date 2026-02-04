using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/PrivateEndpointCreate.json
// this example is just showing the usage of "PrivateEndpoints_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this PrivateEndpointResource
PrivateEndpointCollection collection = resourceGroupResource.GetPrivateEndpoints();

// invoke the operation
string privateEndpointName = "testPe";
PrivateEndpointData data = new PrivateEndpointData
{
    Subnet = new SubnetData
    {
        Id = new ResourceIdentifier("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/myVnet/subnets/mySubnet"),
    },
    IPVersionType = PrivateEndpointIPVersionType.IPv4,
    PrivateLinkServiceConnections = {new NetworkPrivateLinkServiceConnection
    {
    PrivateLinkServiceId = new ResourceIdentifier("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls"),
    GroupIds = {"groupIdFromResource"},
    RequestMessage = "Please approve my connection.",
    }},
    IPConfigurations = {new PrivateEndpointIPConfiguration
    {
    Name = "pestaticconfig",
    GroupId = "file",
    MemberName = "file",
    PrivateIPAddress = IPAddress.Parse("192.168.0.6"),
    }},
    CustomNetworkInterfaceName = "testPeNic",
    Location = new AzureLocation("eastus2euap"),
};
ArmOperation<PrivateEndpointResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointName, data);
PrivateEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PrivateEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

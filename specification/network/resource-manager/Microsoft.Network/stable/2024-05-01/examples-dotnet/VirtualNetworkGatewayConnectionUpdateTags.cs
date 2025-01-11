using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/VirtualNetworkGatewayConnectionUpdateTags.json
// this example is just showing the usage of "VirtualNetworkGatewayConnections_UpdateTags" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualNetworkGatewayConnectionResource created on azure
// for more information of creating VirtualNetworkGatewayConnectionResource, please refer to the document of VirtualNetworkGatewayConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualNetworkGatewayConnectionName = "test";
ResourceIdentifier virtualNetworkGatewayConnectionResourceId = VirtualNetworkGatewayConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkGatewayConnectionName);
VirtualNetworkGatewayConnectionResource virtualNetworkGatewayConnection = client.GetVirtualNetworkGatewayConnectionResource(virtualNetworkGatewayConnectionResourceId);

// invoke the operation
NetworkTagsObject networkTagsObject = new NetworkTagsObject
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
};
ArmOperation<VirtualNetworkGatewayConnectionResource> lro = await virtualNetworkGatewayConnection.UpdateAsync(WaitUntil.Completed, networkTagsObject);
VirtualNetworkGatewayConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualNetworkGatewayConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

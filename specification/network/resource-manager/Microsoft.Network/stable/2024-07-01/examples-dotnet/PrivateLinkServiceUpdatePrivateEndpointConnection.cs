using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/PrivateLinkServiceUpdatePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateLinkServices_UpdatePrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkPrivateEndpointConnectionResource created on azure
// for more information of creating NetworkPrivateEndpointConnectionResource, please refer to the document of NetworkPrivateEndpointConnectionResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string serviceName = "testPls";
string peConnectionName = "testPlePeConnection";
ResourceIdentifier networkPrivateEndpointConnectionResourceId = NetworkPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, peConnectionName);
NetworkPrivateEndpointConnectionResource networkPrivateEndpointConnection = client.GetNetworkPrivateEndpointConnectionResource(networkPrivateEndpointConnectionResourceId);

// invoke the operation
NetworkPrivateEndpointConnectionData data = new NetworkPrivateEndpointConnectionData
{
    ConnectionState = new NetworkPrivateLinkServiceConnectionState
    {
        Status = "Approved",
        Description = "approved it for some reason.",
    },
    Name = "testPlePeConnection",
};
ArmOperation<NetworkPrivateEndpointConnectionResource> lro = await networkPrivateEndpointConnection.UpdateAsync(WaitUntil.Completed, data);
NetworkPrivateEndpointConnectionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkPrivateEndpointConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

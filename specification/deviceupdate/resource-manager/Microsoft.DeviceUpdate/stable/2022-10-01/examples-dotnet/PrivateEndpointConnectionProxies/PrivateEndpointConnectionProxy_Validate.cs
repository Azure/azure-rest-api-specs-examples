using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeviceUpdate;
using Azure.ResourceManager.DeviceUpdate.Models;

// Generated from example definition: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Validate.json
// this example is just showing the usage of "PrivateEndpointConnectionProxies_Validate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateEndpointConnectionProxyResource created on azure
// for more information of creating PrivateEndpointConnectionProxyResource, please refer to the document of PrivateEndpointConnectionProxyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "test-rg";
string accountName = "contoso";
string privateEndpointConnectionProxyId = "peexample01";
ResourceIdentifier privateEndpointConnectionProxyResourceId = PrivateEndpointConnectionProxyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, privateEndpointConnectionProxyId);
PrivateEndpointConnectionProxyResource privateEndpointConnectionProxy = client.GetPrivateEndpointConnectionProxyResource(privateEndpointConnectionProxyResourceId);

// invoke the operation
PrivateEndpointConnectionProxyData data = new PrivateEndpointConnectionProxyData()
{
    RemotePrivateEndpoint = new RemotePrivateEndpoint()
    {
        Id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}",
        Location = new AzureLocation("westus2"),
        ImmutableSubscriptionId = "00000000-0000-0000-0000-000000000000",
        ImmutableResourceId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}",
        ManualPrivateLinkServiceConnections =
        {
        new PrivateLinkServiceConnection()
        {
        Name = "{privateEndpointConnectionProxyId}",
        GroupIds =
        {
        "DeviceUpdate"
        },
        RequestMessage = "Please approve my connection, thanks.",
        }
        },
        PrivateLinkServiceProxies =
        {
        new PrivateLinkServiceProxy()
        {
        Id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}",
        GroupConnectivityInformation =
        {
        },
        }
        },
    },
};
await privateEndpointConnectionProxy.ValidateAsync(data);

Console.WriteLine($"Succeeded");

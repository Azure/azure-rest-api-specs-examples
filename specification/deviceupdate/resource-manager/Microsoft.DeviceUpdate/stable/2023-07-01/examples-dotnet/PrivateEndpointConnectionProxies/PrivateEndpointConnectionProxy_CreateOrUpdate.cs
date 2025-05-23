using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceUpdate.Models;
using Azure.ResourceManager.DeviceUpdate;

// Generated from example definition: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_CreateOrUpdate.json
// this example is just showing the usage of "PrivateEndpointConnectionProxies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceUpdatePrivateEndpointConnectionProxyResource created on azure
// for more information of creating DeviceUpdatePrivateEndpointConnectionProxyResource, please refer to the document of DeviceUpdatePrivateEndpointConnectionProxyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "test-rg";
string accountName = "contoso";
string privateEndpointConnectionProxyId = "peexample01";
ResourceIdentifier deviceUpdatePrivateEndpointConnectionProxyResourceId = DeviceUpdatePrivateEndpointConnectionProxyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, privateEndpointConnectionProxyId);
DeviceUpdatePrivateEndpointConnectionProxyResource deviceUpdatePrivateEndpointConnectionProxy = client.GetDeviceUpdatePrivateEndpointConnectionProxyResource(deviceUpdatePrivateEndpointConnectionProxyResourceId);

// invoke the operation
DeviceUpdatePrivateEndpointConnectionProxyData data = new DeviceUpdatePrivateEndpointConnectionProxyData
{
    RemotePrivateEndpoint = new DeviceUpdateRemotePrivateEndpoint
    {
        Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
        Location = new AzureLocation("westus2"),
        ImmutableSubscriptionId = "00000000-0000-0000-0000-000000000000",
        ImmutableResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
        ManualPrivateLinkServiceConnections = {new DeviceUpdatePrivateLinkServiceConnection
        {
        Name = "{privateEndpointConnectionProxyId}",
        GroupIds = {"DeviceUpdate"},
        RequestMessage = "Please approve my connection, thanks.",
        }},
        PrivateLinkServiceProxies = {new DeviceUpdatePrivateLinkServiceProxy
        {
        Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}"),
        GroupConnectivityInformation = {},
        }},
    },
};
ArmOperation<DeviceUpdatePrivateEndpointConnectionProxyResource> lro = await deviceUpdatePrivateEndpointConnectionProxy.UpdateAsync(WaitUntil.Completed, data);
DeviceUpdatePrivateEndpointConnectionProxyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DeviceUpdatePrivateEndpointConnectionProxyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

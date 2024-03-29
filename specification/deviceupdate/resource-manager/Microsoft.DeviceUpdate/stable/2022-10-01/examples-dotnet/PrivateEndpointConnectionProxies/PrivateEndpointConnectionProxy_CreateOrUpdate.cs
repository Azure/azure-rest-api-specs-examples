using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeviceUpdate;
using Azure.ResourceManager.DeviceUpdate.Models;

// Generated from example definition: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_CreateOrUpdate.json
// this example is just showing the usage of "PrivateEndpointConnectionProxies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceUpdateAccountResource created on azure
// for more information of creating DeviceUpdateAccountResource, please refer to the document of DeviceUpdateAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "test-rg";
string accountName = "contoso";
ResourceIdentifier deviceUpdateAccountResourceId = DeviceUpdateAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
DeviceUpdateAccountResource deviceUpdateAccount = client.GetDeviceUpdateAccountResource(deviceUpdateAccountResourceId);

// get the collection of this PrivateEndpointConnectionProxyResource
PrivateEndpointConnectionProxyCollection collection = deviceUpdateAccount.GetPrivateEndpointConnectionProxies();

// invoke the operation
string privateEndpointConnectionProxyId = "peexample01";
PrivateEndpointConnectionProxyData data = new PrivateEndpointConnectionProxyData()
{
    RemotePrivateEndpoint = new RemotePrivateEndpoint()
    {
        Id = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}",
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
ArmOperation<PrivateEndpointConnectionProxyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateEndpointConnectionProxyId, data);
PrivateEndpointConnectionProxyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PrivateEndpointConnectionProxyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

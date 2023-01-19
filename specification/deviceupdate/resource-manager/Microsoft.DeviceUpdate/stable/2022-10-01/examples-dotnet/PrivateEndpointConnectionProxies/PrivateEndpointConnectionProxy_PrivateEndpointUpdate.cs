using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeviceUpdate;
using Azure.ResourceManager.DeviceUpdate.Models;

// Generated from example definition: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_PrivateEndpointUpdate.json
// this example is just showing the usage of "PrivateEndpointConnectionProxies_UpdatePrivateEndpointProperties" operation, for the dependent resources, they will have to be created separately.

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
PrivateEndpointUpdate privateEndpointUpdate = new PrivateEndpointUpdate()
{
    Id = "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}",
    Location = new AzureLocation("westus2"),
    ImmutableSubscriptionId = "00000000-0000-0000-0000-000000000000",
    ImmutableResourceId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}",
    VnetTrafficTag = "12345678",
};
await privateEndpointConnectionProxy.UpdatePrivateEndpointPropertiesAsync(privateEndpointUpdate);

Console.WriteLine($"Succeeded");

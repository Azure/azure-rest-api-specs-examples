using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeviceUpdate;
using Azure.ResourceManager.DeviceUpdate.Models;

// Generated from example definition: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Get.json
// this example is just showing the usage of "PrivateEndpointConnectionProxies_Get" operation, for the dependent resources, they will have to be created separately.

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
PrivateEndpointConnectionProxyResource result = await privateEndpointConnectionProxy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PrivateEndpointConnectionProxyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

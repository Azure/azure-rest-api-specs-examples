using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/PrivateLinkServiceGetPrivateEndpointConnection.json
// this example is just showing the usage of "PrivateLinkServices_GetPrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PrivateLinkServiceResource created on azure
// for more information of creating PrivateLinkServiceResource, please refer to the document of PrivateLinkServiceResource
string subscriptionId = "subId";
string resourceGroupName = "rg1";
string serviceName = "testPls";
ResourceIdentifier privateLinkServiceResourceId = PrivateLinkServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
PrivateLinkServiceResource privateLinkService = client.GetPrivateLinkServiceResource(privateLinkServiceResourceId);

// get the collection of this NetworkPrivateEndpointConnectionResource
NetworkPrivateEndpointConnectionCollection collection = privateLinkService.GetNetworkPrivateEndpointConnections();

// invoke the operation
string peConnectionName = "testPlePeConnection";
bool result = await collection.ExistsAsync(peConnectionName);

Console.WriteLine($"Succeeded: {result}");

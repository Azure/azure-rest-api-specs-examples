using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/PrivateLinkServiceDeletePrivateEndpointConnection.json
// this example is just showing the usage of "PrivateLinkServices_DeletePrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.

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
await networkPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

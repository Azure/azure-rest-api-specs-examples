using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualRouterPeerListAdvertisedRoute.json
// this example is just showing the usage of "VirtualHubBgpConnections_ListAdvertisedRoutes" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BgpConnectionResource created on azure
// for more information of creating BgpConnectionResource, please refer to the document of BgpConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string hubName = "virtualRouter1";
string connectionName = "peer1";
ResourceIdentifier bgpConnectionResourceId = BgpConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hubName, connectionName);
BgpConnectionResource bgpConnection = client.GetBgpConnectionResource(bgpConnectionResourceId);

// invoke the operation
ArmOperation<IDictionary<string, IList<PeerRoute>>> lro = await bgpConnection.GetVirtualHubBgpConnectionAdvertisedRoutesAsync(WaitUntil.Completed);
IDictionary<string, IList<PeerRoute>> result = lro.Value;

Console.WriteLine($"Succeeded: {result}");

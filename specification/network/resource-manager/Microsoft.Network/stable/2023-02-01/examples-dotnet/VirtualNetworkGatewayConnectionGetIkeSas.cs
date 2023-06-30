using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;
using Azure.ResourceManager.Network.Models;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/VirtualNetworkGatewayConnectionGetIkeSas.json
// this example is just showing the usage of "VirtualNetworkGatewayConnections_GetIkeSas" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualNetworkGatewayConnectionResource created on azure
// for more information of creating VirtualNetworkGatewayConnectionResource, please refer to the document of VirtualNetworkGatewayConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualNetworkGatewayConnectionName = "vpngwcn1";
ResourceIdentifier virtualNetworkGatewayConnectionResourceId = VirtualNetworkGatewayConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualNetworkGatewayConnectionName);
VirtualNetworkGatewayConnectionResource virtualNetworkGatewayConnection = client.GetVirtualNetworkGatewayConnectionResource(virtualNetworkGatewayConnectionResourceId);

// invoke the operation
ArmOperation<string> lro = await virtualNetworkGatewayConnection.GetIkeSasAsync(WaitUntil.Completed);
string result = lro.Value;

Console.WriteLine($"Succeeded: {result}");

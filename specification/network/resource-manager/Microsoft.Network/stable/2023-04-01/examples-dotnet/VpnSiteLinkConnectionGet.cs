using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2023-04-01/examples/VpnSiteLinkConnectionGet.json
// this example is just showing the usage of "VpnSiteLinkConnections_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VpnConnectionResource created on azure
// for more information of creating VpnConnectionResource, please refer to the document of VpnConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string gatewayName = "gateway1";
string connectionName = "vpnConnection1";
ResourceIdentifier vpnConnectionResourceId = VpnConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gatewayName, connectionName);
VpnConnectionResource vpnConnection = client.GetVpnConnectionResource(vpnConnectionResourceId);

// get the collection of this VpnSiteLinkConnectionResource
VpnSiteLinkConnectionCollection collection = vpnConnection.GetVpnSiteLinkConnections();

// invoke the operation
string linkConnectionName = "Connection-Link1";
bool result = await collection.ExistsAsync(linkConnectionName);

Console.WriteLine($"Succeeded: {result}");

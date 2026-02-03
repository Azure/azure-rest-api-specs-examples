using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VpnConnectionStartPacketCapture.json
// this example is just showing the usage of "VpnConnections_StartPacketCapture" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VpnConnectionResource created on azure
// for more information of creating VpnConnectionResource, please refer to the document of VpnConnectionResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string gatewayName = "gateway1";
string vpnConnectionName = "vpnConnection1";
ResourceIdentifier vpnConnectionResourceId = VpnConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gatewayName, vpnConnectionName);
VpnConnectionResource vpnConnection = client.GetVpnConnectionResource(vpnConnectionResourceId);

// invoke the operation
VpnConnectionPacketCaptureStartContent content = new VpnConnectionPacketCaptureStartContent
{
    LinkConnectionNames = { "siteLink1", "siteLink2" },
};
ArmOperation<string> lro = await vpnConnection.StartPacketCaptureAsync(WaitUntil.Completed, content: content);
string result = lro.Value;

Console.WriteLine($"Succeeded: {result}");

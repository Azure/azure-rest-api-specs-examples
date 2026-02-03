using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VpnConnectionStopPacketCapture.json
// this example is just showing the usage of "VpnConnections_StopPacketCapture" operation, for the dependent resources, they will have to be created separately.

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
VpnConnectionPacketCaptureStopContent content = new VpnConnectionPacketCaptureStopContent
{
    SasUri = new Uri("https://teststorage.blob.core.windows.net/?sv=2018-03-28&ss=bfqt&srt=sco&sp=rwdlacup&se=2019-09-13T07:44:05Z&st=2019-09-06T23:44:05Z&spr=https&sig=V1h9D1riltvZMI69d6ihENnFo%2FrCvTqGgjO2lf%2FVBhE%3D"),
    LinkConnectionNames = { "vpnSiteLink1", "vpnSiteLink2" },
};
ArmOperation<string> lro = await vpnConnection.StopPacketCaptureAsync(WaitUntil.Completed, content: content);
string result = lro.Value;

Console.WriteLine($"Succeeded: {result}");

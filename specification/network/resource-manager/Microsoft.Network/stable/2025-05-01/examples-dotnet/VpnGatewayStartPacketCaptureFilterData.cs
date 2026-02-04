using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VpnGatewayStartPacketCaptureFilterData.json
// this example is just showing the usage of "VpnGateways_StartPacketCapture" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VpnGatewayResource created on azure
// for more information of creating VpnGatewayResource, please refer to the document of VpnGatewayResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string gatewayName = "vpngw";
ResourceIdentifier vpnGatewayResourceId = VpnGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gatewayName);
VpnGatewayResource vpnGateway = client.GetVpnGatewayResource(vpnGatewayResourceId);

// invoke the operation
VpnGatewayPacketCaptureStartContent content = new VpnGatewayPacketCaptureStartContent
{
    FilterData = "{'TracingFlags': 11,'MaxPacketBufferSize': 120,'MaxFileSize': 200,'Filters': [{'SourceSubnets': ['20.1.1.0/24'],'DestinationSubnets': ['10.1.1.0/24'],'SourcePort': [500],'DestinationPort': [4500],'Protocol': 6,'TcpFlags': 16,'CaptureSingleDirectionTrafficOnly': true}]}",
};
ArmOperation<string> lro = await vpnGateway.StartPacketCaptureAsync(WaitUntil.Completed, content: content);
string result = lro.Value;

Console.WriteLine($"Succeeded: {result}");

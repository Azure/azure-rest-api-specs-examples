using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkWatcherPacketCaptureQueryStatus.json
// this example is just showing the usage of "PacketCaptures_GetStatus" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PacketCaptureResource created on azure
// for more information of creating PacketCaptureResource, please refer to the document of PacketCaptureResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkWatcherName = "nw1";
string packetCaptureName = "pc1";
ResourceIdentifier packetCaptureResourceId = PacketCaptureResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkWatcherName, packetCaptureName);
PacketCaptureResource packetCapture = client.GetPacketCaptureResource(packetCaptureResourceId);

// invoke the operation
ArmOperation<PacketCaptureQueryStatusResult> lro = await packetCapture.GetStatusAsync(WaitUntil.Completed);
PacketCaptureQueryStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");

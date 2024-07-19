using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MobileNetwork.Models;
using Azure.ResourceManager.MobileNetwork;

// Generated from example definition: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/PacketCaptureStop.json
// this example is just showing the usage of "PacketCaptures_Stop" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MobileNetworkPacketCaptureResource created on azure
// for more information of creating MobileNetworkPacketCaptureResource, please refer to the document of MobileNetworkPacketCaptureResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string packetCoreControlPlaneName = "TestPacketCoreCP";
string packetCaptureName = "pc1";
ResourceIdentifier mobileNetworkPacketCaptureResourceId = MobileNetworkPacketCaptureResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, packetCoreControlPlaneName, packetCaptureName);
MobileNetworkPacketCaptureResource mobileNetworkPacketCapture = client.GetMobileNetworkPacketCaptureResource(mobileNetworkPacketCaptureResourceId);

// invoke the operation
ArmOperation<AsyncOperationStatus> lro = await mobileNetworkPacketCapture.StopAsync(WaitUntil.Completed);
AsyncOperationStatus result = lro.Value;

Console.WriteLine($"Succeeded: {result}");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/NetworkDevices_Upgrade_MaximumSet_Gen.json
// this example is just showing the usage of "NetworkDevices_Upgrade" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkDeviceResource created on azure
// for more information of creating NetworkDeviceResource, please refer to the document of NetworkDeviceResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string networkDeviceName = "example-device";
ResourceIdentifier networkDeviceResourceId = NetworkDeviceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkDeviceName);
NetworkDeviceResource networkDevice = client.GetNetworkDeviceResource(networkDeviceResourceId);

// invoke the operation
NetworkFabricUpdateVersionContent content = new NetworkFabricUpdateVersionContent
{
    Version = "1.0.0",
};
ArmOperation<StateUpdateCommonPostActionResult> lro = await networkDevice.UpgradeAsync(WaitUntil.Completed, content);
StateUpdateCommonPostActionResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");

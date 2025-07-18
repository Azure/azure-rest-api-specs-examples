using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkWatcherPacketCaptureCreate.json
// this example is just showing the usage of "PacketCaptures_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkWatcherResource created on azure
// for more information of creating NetworkWatcherResource, please refer to the document of NetworkWatcherResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkWatcherName = "nw1";
ResourceIdentifier networkWatcherResourceId = NetworkWatcherResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkWatcherName);
NetworkWatcherResource networkWatcher = client.GetNetworkWatcherResource(networkWatcherResourceId);

// get the collection of this PacketCaptureResource
PacketCaptureCollection collection = networkWatcher.GetPacketCaptures();

// invoke the operation
string packetCaptureName = "pc1";
PacketCaptureCreateOrUpdateContent content = new PacketCaptureCreateOrUpdateContent("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Compute/virtualMachines/vm1", new PacketCaptureStorageLocation
{
    StorageId = new ResourceIdentifier("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Storage/storageAccounts/pcstore"),
    StoragePath = "https://mytestaccountname.blob.core.windows.net/capture/pc1.cap",
    FilePath = "D:\\capture\\pc1.cap",
})
{
    BytesToCapturePerPacket = 10000L,
    TotalBytesPerSession = 100000L,
    TimeLimitInSeconds = 100,
    Filters = {new PacketCaptureFilter
    {
    Protocol = PcProtocol.Tcp,
    LocalIPAddress = "10.0.0.4",
    LocalPort = "80",
    }},
};
ArmOperation<PacketCaptureResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, packetCaptureName, content);
PacketCaptureResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PacketCaptureData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

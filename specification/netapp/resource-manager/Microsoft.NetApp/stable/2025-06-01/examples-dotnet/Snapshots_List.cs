using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-06-01/examples/Snapshots_List.json
// this example is just showing the usage of "Snapshots_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppVolumeResource created on azure
// for more information of creating NetAppVolumeResource, please refer to the document of NetAppVolumeResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
string volumeName = "volume1";
ResourceIdentifier netAppVolumeResourceId = NetAppVolumeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName);
NetAppVolumeResource netAppVolume = client.GetNetAppVolumeResource(netAppVolumeResourceId);

// get the collection of this NetAppVolumeSnapshotResource
NetAppVolumeSnapshotCollection collection = netAppVolume.GetNetAppVolumeSnapshots();

// invoke the operation and iterate over the result
await foreach (NetAppVolumeSnapshotResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetAppVolumeSnapshotData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

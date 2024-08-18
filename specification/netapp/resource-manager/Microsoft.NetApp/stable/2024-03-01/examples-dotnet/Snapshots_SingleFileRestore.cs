using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/Snapshots_SingleFileRestore.json
// this example is just showing the usage of "Snapshots_RestoreFiles" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppVolumeSnapshotResource created on azure
// for more information of creating NetAppVolumeSnapshotResource, please refer to the document of NetAppVolumeSnapshotResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
string volumeName = "volume1";
string snapshotName = "snapshot1";
ResourceIdentifier netAppVolumeSnapshotResourceId = NetAppVolumeSnapshotResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName, snapshotName);
NetAppVolumeSnapshotResource netAppVolumeSnapshot = client.GetNetAppVolumeSnapshotResource(netAppVolumeSnapshotResourceId);

// invoke the operation
NetAppVolumeSnapshotRestoreFilesContent content = new NetAppVolumeSnapshotRestoreFilesContent(new string[]
{
"/dir1/customer1.db","/dir1/customer2.db"
});
await netAppVolumeSnapshot.RestoreFilesAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");

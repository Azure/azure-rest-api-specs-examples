using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetApp;
using Azure.ResourceManager.NetApp.Models;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/Snapshots_Get.json
// this example is just showing the usage of "Snapshots_Get" operation, for the dependent resources, they will have to be created separately.

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
NetAppVolumeSnapshotResource result = await netAppVolumeSnapshot.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppVolumeSnapshotData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

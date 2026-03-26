using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-12-15-preview/examples/ElasticSnapshots_CreateOrUpdate.json
// this example is just showing the usage of "ElasticSnapshots_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppElasticVolumeResource created on azure
// for more information of creating NetAppElasticVolumeResource, please refer to the document of NetAppElasticVolumeResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
string volumeName = "volume1";
ResourceIdentifier netAppElasticVolumeResourceId = NetAppElasticVolumeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName);
NetAppElasticVolumeResource netAppElasticVolume = client.GetNetAppElasticVolumeResource(netAppElasticVolumeResourceId);

// get the collection of this NetAppElasticSnapshotResource
NetAppElasticSnapshotCollection collection = netAppElasticVolume.GetNetAppElasticSnapshots();

// invoke the operation
string snapshotName = "snapshot1";
NetAppElasticSnapshotData data = new NetAppElasticSnapshotData();
ArmOperation<NetAppElasticSnapshotResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, snapshotName, data);
NetAppElasticSnapshotResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppElasticSnapshotData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

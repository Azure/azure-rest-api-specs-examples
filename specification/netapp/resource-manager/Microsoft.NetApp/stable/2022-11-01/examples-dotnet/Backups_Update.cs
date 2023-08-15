using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetApp;
using Azure.ResourceManager.NetApp.Models;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/Backups_Update.json
// this example is just showing the usage of "Backups_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppVolumeBackupResource created on azure
// for more information of creating NetAppVolumeBackupResource, please refer to the document of NetAppVolumeBackupResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
string volumeName = "volume1";
string backupName = "backup1";
ResourceIdentifier netAppVolumeBackupResourceId = NetAppVolumeBackupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName, backupName);
NetAppVolumeBackupResource netAppVolumeBackup = client.GetNetAppVolumeBackupResource(netAppVolumeBackupResourceId);

// invoke the operation
NetAppVolumeBackupPatch patch = new NetAppVolumeBackupPatch();
ArmOperation<NetAppVolumeBackupResource> lro = await netAppVolumeBackup.UpdateAsync(WaitUntil.Completed, patch);
NetAppVolumeBackupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppBackupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

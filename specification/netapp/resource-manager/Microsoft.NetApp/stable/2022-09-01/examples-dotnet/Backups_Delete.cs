using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetApp;
using Azure.ResourceManager.NetApp.Models;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-09-01/examples/Backups_Delete.json
// this example is just showing the usage of "Backups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppVolumeBackupResource created on azure
// for more information of creating NetAppVolumeBackupResource, please refer to the document of NetAppVolumeBackupResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "resourceGroup";
string accountName = "accountName";
string poolName = "poolName";
string volumeName = "volumeName";
string backupName = "backupName";
ResourceIdentifier netAppVolumeBackupResourceId = NetAppVolumeBackupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName, backupName);
NetAppVolumeBackupResource netAppVolumeBackup = client.GetNetAppVolumeBackupResource(netAppVolumeBackupResourceId);

// invoke the operation
await netAppVolumeBackup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

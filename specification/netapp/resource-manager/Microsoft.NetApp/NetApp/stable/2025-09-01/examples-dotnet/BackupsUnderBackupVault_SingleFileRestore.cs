using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/stable/2025-09-01/examples/BackupsUnderBackupVault_SingleFileRestore.json
// this example is just showing the usage of "BackupsUnderBackupVault_RestoreFiles" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppBackupVaultBackupResource created on azure
// for more information of creating NetAppBackupVaultBackupResource, please refer to the document of NetAppBackupVaultBackupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string backupVaultName = "backupVault1";
string backupName = "backup1";
ResourceIdentifier netAppBackupVaultBackupResourceId = NetAppBackupVaultBackupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, backupVaultName, backupName);
NetAppBackupVaultBackupResource netAppBackupVaultBackup = client.GetNetAppBackupVaultBackupResource(netAppBackupVaultBackupResourceId);

// invoke the operation
NetAppVolumeBackupBackupRestoreFilesContent content = new NetAppVolumeBackupBackupRestoreFilesContent(new string[] { "/dir1/customer1.db", "/dir1/customer2.db" }, new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1/volumes/volume1"));
await netAppBackupVaultBackup.RestoreFilesBackupsUnderBackupVaultAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-09-01-preview/examples/BackupsUnderBackupVault_Update.json
// this example is just showing the usage of "Backups_Update" operation, for the dependent resources, they will have to be created separately.

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
NetAppBackupVaultBackupPatch patch = new NetAppBackupVaultBackupPatch();
ArmOperation<NetAppBackupVaultBackupResource> lro = await netAppBackupVaultBackup.UpdateAsync(WaitUntil.Completed, patch);
NetAppBackupVaultBackupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppBackupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

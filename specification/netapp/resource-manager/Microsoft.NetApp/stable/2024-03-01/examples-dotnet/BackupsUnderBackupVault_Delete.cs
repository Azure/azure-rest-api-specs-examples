using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/BackupsUnderBackupVault_Delete.json
// this example is just showing the usage of "Backups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppBackupVaultBackupResource created on azure
// for more information of creating NetAppBackupVaultBackupResource, please refer to the document of NetAppBackupVaultBackupResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "resourceGroup";
string accountName = "account1";
string backupVaultName = "backupVault1";
string backupName = "backup1";
ResourceIdentifier netAppBackupVaultBackupResourceId = NetAppBackupVaultBackupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, backupVaultName, backupName);
NetAppBackupVaultBackupResource netAppBackupVaultBackup = client.GetNetAppBackupVaultBackupResource(netAppBackupVaultBackupResourceId);

// invoke the operation
await netAppBackupVaultBackup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

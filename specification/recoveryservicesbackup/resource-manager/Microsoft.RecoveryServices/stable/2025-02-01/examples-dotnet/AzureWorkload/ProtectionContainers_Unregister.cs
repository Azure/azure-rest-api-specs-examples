using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.RecoveryServicesBackup;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureWorkload/ProtectionContainers_Unregister.json
// this example is just showing the usage of "ProtectionContainers_Unregister" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupProtectionContainerResource created on azure
// for more information of creating BackupProtectionContainerResource, please refer to the document of BackupProtectionContainerResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "testRg";
string vaultName = "testVault";
string fabricName = "Azure";
string containerName = "storagecontainer;Storage;test-rg;teststorage";
ResourceIdentifier backupProtectionContainerResourceId = BackupProtectionContainerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, fabricName, containerName);
BackupProtectionContainerResource backupProtectionContainer = client.GetBackupProtectionContainerResource(backupProtectionContainerResourceId);

// invoke the operation
await backupProtectionContainer.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

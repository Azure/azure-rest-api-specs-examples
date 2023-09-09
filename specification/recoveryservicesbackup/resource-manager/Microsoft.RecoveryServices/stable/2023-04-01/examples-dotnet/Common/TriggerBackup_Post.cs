using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/Common/TriggerBackup_Post.json
// this example is just showing the usage of "Backups_Trigger" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupProtectedItemResource created on azure
// for more information of creating BackupProtectedItemResource, please refer to the document of BackupProtectedItemResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "linuxRsVaultRG";
string vaultName = "linuxRsVault";
string fabricName = "Azure";
string containerName = "IaasVMContainer;iaasvmcontainerv2;testrg;v1win2012r";
string protectedItemName = "VM;iaasvmcontainerv2;testrg;v1win2012r";
ResourceIdentifier backupProtectedItemResourceId = BackupProtectedItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, fabricName, containerName, protectedItemName);
BackupProtectedItemResource backupProtectedItem = client.GetBackupProtectedItemResource(backupProtectedItemResourceId);

// invoke the operation
TriggerBackupContent content = new TriggerBackupContent(new AzureLocation("placeholder"))
{
    Properties = new IaasVmBackupContent(),
};
await backupProtectedItem.TriggerBackupAsync(content);

Console.WriteLine($"Succeeded");

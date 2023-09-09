using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureIaasVm/Revoke_Ilr.json
// this example is just showing the usage of "ItemLevelRecoveryConnections_Revoke" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupRecoveryPointResource created on azure
// for more information of creating BackupRecoveryPointResource, please refer to the document of BackupRecoveryPointResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "PythonSDKBackupTestRg";
string vaultName = "PySDKBackupTestRsVault";
string fabricName = "Azure";
string containerName = "iaasvmcontainer;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1";
string protectedItemName = "vm;iaasvmcontainerv2;pysdktestrg;pysdktestv2vm1";
string recoveryPointId = "1";
ResourceIdentifier backupRecoveryPointResourceId = BackupRecoveryPointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, fabricName, containerName, protectedItemName, recoveryPointId);
BackupRecoveryPointResource backupRecoveryPoint = client.GetBackupRecoveryPointResource(backupRecoveryPointResourceId);

// invoke the operation
await backupRecoveryPoint.RevokeItemLevelRecoveryConnectionAsync();

Console.WriteLine($"Succeeded");

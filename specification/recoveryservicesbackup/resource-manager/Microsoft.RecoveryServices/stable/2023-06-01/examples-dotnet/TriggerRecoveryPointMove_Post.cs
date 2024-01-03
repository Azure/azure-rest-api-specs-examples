using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/TriggerRecoveryPointMove_Post.json
// this example is just showing the usage of "MoveRecoveryPoint" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupRecoveryPointResource created on azure
// for more information of creating BackupRecoveryPointResource, please refer to the document of BackupRecoveryPointResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "netsdktestrg";
string vaultName = "testVault";
string fabricName = "Azure";
string containerName = "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
string protectedItemName = "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
string recoveryPointId = "348916168024334";
ResourceIdentifier backupRecoveryPointResourceId = BackupRecoveryPointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, fabricName, containerName, protectedItemName, recoveryPointId);
BackupRecoveryPointResource backupRecoveryPoint = client.GetBackupRecoveryPointResource(backupRecoveryPointResourceId);

// invoke the operation
MoveRPAcrossTiersContent content = new MoveRPAcrossTiersContent()
{
    ObjectType = "MoveRPAcrossTiersRequest",
    SourceTierType = RecoveryPointTierType.HardenedRP,
    TargetTierType = RecoveryPointTierType.ArchivedRP,
};
await backupRecoveryPoint.MoveRecoveryPointAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");

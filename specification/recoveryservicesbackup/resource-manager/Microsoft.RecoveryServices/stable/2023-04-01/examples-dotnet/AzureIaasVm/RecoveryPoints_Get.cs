using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureIaasVm/RecoveryPoints_Get.json
// this example is just showing the usage of "RecoveryPoints_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupProtectedItemResource created on azure
// for more information of creating BackupProtectedItemResource, please refer to the document of BackupProtectedItemResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rshhtestmdvmrg";
string vaultName = "rshvault";
string fabricName = "Azure";
string containerName = "IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall";
string protectedItemName = "VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall";
ResourceIdentifier backupProtectedItemResourceId = BackupProtectedItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, fabricName, containerName, protectedItemName);
BackupProtectedItemResource backupProtectedItem = client.GetBackupProtectedItemResource(backupProtectedItemResourceId);

// get the collection of this BackupRecoveryPointResource
BackupRecoveryPointCollection collection = backupProtectedItem.GetBackupRecoveryPoints();

// invoke the operation
string recoveryPointId = "26083826328862";
bool result = await collection.ExistsAsync(recoveryPointId);

Console.WriteLine($"Succeeded: {result}");

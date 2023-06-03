using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureIaasVm/RecoveryPoints_Get.json
// this example is just showing the usage of "RecoveryPoints_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupRecoveryPointResource created on azure
// for more information of creating BackupRecoveryPointResource, please refer to the document of BackupRecoveryPointResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rshhtestmdvmrg";
string vaultName = "rshvault";
string fabricName = "Azure";
string containerName = "IaasVMContainer;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall";
string protectedItemName = "VM;iaasvmcontainerv2;rshhtestmdvmrg;rshmdvmsmall";
string recoveryPointId = "26083826328862";
ResourceIdentifier backupRecoveryPointResourceId = BackupRecoveryPointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, fabricName, containerName, protectedItemName, recoveryPointId);
BackupRecoveryPointResource backupRecoveryPoint = client.GetBackupRecoveryPointResource(backupRecoveryPointResourceId);

// invoke the operation
BackupRecoveryPointResource result = await backupRecoveryPoint.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupRecoveryPointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

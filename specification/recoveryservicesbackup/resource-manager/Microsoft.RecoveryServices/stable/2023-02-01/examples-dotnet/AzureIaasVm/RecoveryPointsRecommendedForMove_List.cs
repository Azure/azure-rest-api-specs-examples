using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureIaasVm/RecoveryPointsRecommendedForMove_List.json
// this example is just showing the usage of "RecoveryPointsRecommendedForMove_List" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
RecoveryPointsRecommendedForMoveContent content = new RecoveryPointsRecommendedForMoveContent()
{
    ObjectType = "ListRecoveryPointsRecommendedForMoveRequest",
    ExcludedRPList =
    {
    "348916168024334","348916168024335"
    },
};
await foreach (BackupRecoveryPointResource item in backupProtectedItem.GetRecoveryPointsRecommendedForMoveAsync(content))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    BackupRecoveryPointData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");

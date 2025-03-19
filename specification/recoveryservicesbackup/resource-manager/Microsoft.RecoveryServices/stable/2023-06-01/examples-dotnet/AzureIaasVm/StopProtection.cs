using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.RecoveryServicesBackup;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/AzureIaasVm/StopProtection.json
// this example is just showing the usage of "ProtectedItems_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupProtectedItemResource created on azure
// for more information of creating BackupProtectedItemResource, please refer to the document of BackupProtectedItemResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "SwaggerTestRg";
string vaultName = "NetSDKTestRsVault";
string fabricName = "Azure";
string containerName = "IaasVMContainer;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
string protectedItemName = "VM;iaasvmcontainerv2;netsdktestrg;netvmtestv2vm1";
ResourceIdentifier backupProtectedItemResourceId = BackupProtectedItemResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, fabricName, containerName, protectedItemName);
BackupProtectedItemResource backupProtectedItem = client.GetBackupProtectedItemResource(backupProtectedItemResourceId);

// invoke the operation
BackupProtectedItemData data = new BackupProtectedItemData(default)
{
    Properties = new IaasComputeVmProtectedItem
    {
        ProtectionState = BackupProtectionState.ProtectionStopped,
        SourceResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
    },
};
ArmOperation<BackupProtectedItemResource> lro = await backupProtectedItem.UpdateAsync(WaitUntil.Completed, data);
BackupProtectedItemResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupProtectedItemData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

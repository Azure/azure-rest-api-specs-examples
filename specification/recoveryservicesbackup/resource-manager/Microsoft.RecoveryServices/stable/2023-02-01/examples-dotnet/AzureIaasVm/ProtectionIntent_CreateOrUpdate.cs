using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureIaasVm/ProtectionIntent_CreateOrUpdate.json
// this example is just showing the usage of "ProtectionIntent_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupProtectionIntentResource created on azure
// for more information of creating BackupProtectionIntentResource, please refer to the document of BackupProtectionIntentResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string vaultName = "myVault";
string fabricName = "Azure";
string intentObjectName = "vm;iaasvmcontainerv2;chamsrgtest;chamscandel";
ResourceIdentifier backupProtectionIntentResourceId = BackupProtectionIntentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, fabricName, intentObjectName);
BackupProtectionIntentResource backupProtectionIntent = client.GetBackupProtectionIntentResource(backupProtectionIntentResourceId);

// invoke the operation
BackupProtectionIntentData data = new BackupProtectionIntentData(new AzureLocation("placeholder"))
{
    Properties = new ResourceProtectionIntent()
    {
        SourceResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/chamsrgtest/providers/Microsoft.Compute/virtualMachines/chamscandel"),
        PolicyId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.RecoveryServices/vaults/myVault/backupPolicies/myPolicy"),
    },
};
ArmOperation<BackupProtectionIntentResource> lro = await backupProtectionIntent.UpdateAsync(WaitUntil.Completed, data);
BackupProtectionIntentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupProtectionIntentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

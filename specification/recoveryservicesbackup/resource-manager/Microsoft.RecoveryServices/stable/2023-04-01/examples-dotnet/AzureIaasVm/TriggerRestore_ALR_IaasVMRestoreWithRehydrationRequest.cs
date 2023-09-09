using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureIaasVm/TriggerRestore_ALR_IaasVMRestoreWithRehydrationRequest.json
// this example is just showing the usage of "Restores_Trigger" operation, for the dependent resources, they will have to be created separately.

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
TriggerRestoreContent content = new TriggerRestoreContent(new AzureLocation("placeholder"))
{
    Properties = new IaasVmRestoreWithRehydrationContent()
    {
        RecoveryPointRehydrationInfo = new RecoveryPointRehydrationInfo()
        {
            RehydrationRetentionDuration = XmlConvert.ToTimeSpan("P7D"),
            RehydrationPriority = RehydrationPriority.High,
        },
        RecoveryPointId = "348916168024334",
        RecoveryType = FileShareRecoveryType.AlternateLocation,
        SourceResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg/providers/Microsoft.Compute/virtualMachines/netvmtestv2vm1"),
        TargetVirtualMachineId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2/providers/Microsoft.Compute/virtualmachines/RSMDALRVM981435"),
        TargetResourceGroupId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/netsdktestrg2"),
        StorageAccountId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Storage/storageAccounts/testingAccount"),
        VirtualNetworkId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet"),
        SubnetId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRg/providers/Microsoft.Network/virtualNetworks/testNet/subnets/default"),
        Region = new AzureLocation("southeastasia"),
        DoesCreateNewCloudService = false,
        OriginalStorageAccountOption = false,
        EncryptionDetails = new VmEncryptionDetails()
        {
            IsEncryptionEnabled = false,
        },
    },
};
await backupRecoveryPoint.TriggerRestoreAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");

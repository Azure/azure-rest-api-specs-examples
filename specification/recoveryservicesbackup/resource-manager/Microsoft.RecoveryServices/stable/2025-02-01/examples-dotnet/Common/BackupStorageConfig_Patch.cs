using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.RecoveryServicesBackup;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/Common/BackupStorageConfig_Patch.json
// this example is just showing the usage of "BackupResourceStorageConfigsNonCRR_Patch" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupResourceConfigResource created on azure
// for more information of creating BackupResourceConfigResource, please refer to the document of BackupResourceConfigResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "PythonSDKBackupTestRg";
string vaultName = "PySDKBackupTestRsVault";
ResourceIdentifier backupResourceConfigResourceId = BackupResourceConfigResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName);
BackupResourceConfigResource backupResourceConfig = client.GetBackupResourceConfigResource(backupResourceConfigResourceId);

// invoke the operation
BackupResourceConfigData data = new BackupResourceConfigData(default)
{
    Properties = new BackupResourceConfigProperties
    {
        StorageType = BackupStorageType.LocallyRedundant,
        StorageTypeState = BackupStorageTypeState.Unlocked,
    },
};
BackupResourceConfigResource result = await backupResourceConfig.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupResourceConfigData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

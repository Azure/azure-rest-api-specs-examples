using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.RecoveryServicesBackup;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/Common/BackupStorageConfig_Put.json
// this example is just showing the usage of "BackupResourceStorageConfigsNonCRR_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "PythonSDKBackupTestRg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this BackupResourceConfigResource
BackupResourceConfigCollection collection = resourceGroupResource.GetBackupResourceConfigs();

// invoke the operation
string vaultName = "PySDKBackupTestRsVault";
BackupResourceConfigData data = new BackupResourceConfigData(default)
{
    Properties = new BackupResourceConfigProperties
    {
        StorageType = BackupStorageType.LocallyRedundant,
        StorageTypeState = BackupStorageTypeState.Unlocked,
    },
};
ArmOperation<BackupResourceConfigResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vaultName, data);
BackupResourceConfigResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupResourceConfigData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

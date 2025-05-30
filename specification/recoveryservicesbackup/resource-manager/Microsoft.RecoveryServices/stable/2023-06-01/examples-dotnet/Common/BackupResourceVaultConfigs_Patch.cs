using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.RecoveryServicesBackup;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/Common/BackupResourceVaultConfigs_Patch.json
// this example is just showing the usage of "BackupResourceVaultConfigs_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupResourceVaultConfigResource created on azure
// for more information of creating BackupResourceVaultConfigResource, please refer to the document of BackupResourceVaultConfigResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "SwaggerTestRg";
string vaultName = "SwaggerTest";
ResourceIdentifier backupResourceVaultConfigResourceId = BackupResourceVaultConfigResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName);
BackupResourceVaultConfigResource backupResourceVaultConfig = client.GetBackupResourceVaultConfigResource(backupResourceVaultConfigResourceId);

// invoke the operation
BackupResourceVaultConfigData data = new BackupResourceVaultConfigData(default)
{
    Properties = new BackupResourceVaultConfigProperties
    {
        EnhancedSecurityState = EnhancedSecurityState.Enabled,
    },
};
BackupResourceVaultConfigResource result = await backupResourceVaultConfig.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupResourceVaultConfigData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

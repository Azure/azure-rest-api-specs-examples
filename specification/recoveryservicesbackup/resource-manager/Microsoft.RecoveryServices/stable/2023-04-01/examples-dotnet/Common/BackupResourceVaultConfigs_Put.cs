using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/Common/BackupResourceVaultConfigs_Put.json
// this example is just showing the usage of "BackupResourceVaultConfigs_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "SwaggerTestRg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this BackupResourceVaultConfigResource
BackupResourceVaultConfigCollection collection = resourceGroupResource.GetBackupResourceVaultConfigs();

// invoke the operation
string vaultName = "SwaggerTest";
BackupResourceVaultConfigData data = new BackupResourceVaultConfigData(new AzureLocation("placeholder"))
{
    Properties = new BackupResourceVaultConfigProperties()
    {
        EnhancedSecurityState = EnhancedSecurityState.Enabled,
        SoftDeleteFeatureState = SoftDeleteFeatureState.Enabled,
    },
};
ArmOperation<BackupResourceVaultConfigResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, vaultName, data);
BackupResourceVaultConfigResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupResourceVaultConfigData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

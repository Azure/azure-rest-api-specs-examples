using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/BackupResourceEncryptionConfig_Put.json
// this example is just showing the usage of "BackupResourceEncryptionConfigs_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupResourceEncryptionConfigExtendedResource created on azure
// for more information of creating BackupResourceEncryptionConfigExtendedResource, please refer to the document of BackupResourceEncryptionConfigExtendedResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "test-rg";
string vaultName = "source-rsv";
ResourceIdentifier backupResourceEncryptionConfigExtendedResourceId = BackupResourceEncryptionConfigExtendedResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName);
BackupResourceEncryptionConfigExtendedResource backupResourceEncryptionConfigExtended = client.GetBackupResourceEncryptionConfigExtendedResource(backupResourceEncryptionConfigExtendedResourceId);

// invoke the operation
BackupResourceEncryptionConfigExtendedCreateOrUpdateContent content = new BackupResourceEncryptionConfigExtendedCreateOrUpdateContent(new AzureLocation("placeholder"))
{
    Properties = new BackupResourceEncryptionConfig()
    {
        EncryptionAtRestType = BackupEncryptionAtRestType.CustomerManaged,
        KeyUri = new Uri("https://gktestkv1.vault.azure.net/keys/Test1/ed2e8cdc7f86477ebf0c6462b504a9ed"),
        SubscriptionId = "1a2311d9-66f5-47d3-a9fb-7a37da63934b",
        InfrastructureEncryptionState = new InfrastructureEncryptionState("true"),
    },
};
await backupResourceEncryptionConfigExtended.UpdateAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");

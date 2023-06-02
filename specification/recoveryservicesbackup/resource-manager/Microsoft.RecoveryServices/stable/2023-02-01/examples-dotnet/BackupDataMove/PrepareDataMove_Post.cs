using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/BackupDataMove/PrepareDataMove_Post.json
// this example is just showing the usage of "BMSPrepareDataMove" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupResourceConfigResource created on azure
// for more information of creating BackupResourceConfigResource, please refer to the document of BackupResourceConfigResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "sourceRG";
string vaultName = "source-rsv";
ResourceIdentifier backupResourceConfigResourceId = BackupResourceConfigResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName);
BackupResourceConfigResource backupResourceConfig = client.GetBackupResourceConfigResource(backupResourceConfigResourceId);

// invoke the operation
PrepareDataMoveContent content = new PrepareDataMoveContent(new ResourceIdentifier("/subscriptions/04cf684a-d41f-4550-9f70-7708a3a2283b/resourceGroups/targetRG/providers/Microsoft.RecoveryServices/vaults/target-rsv"), new AzureLocation("USGov Virginia"), DataMoveLevel.Vault);
await backupResourceConfig.PrepareDataMoveAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");

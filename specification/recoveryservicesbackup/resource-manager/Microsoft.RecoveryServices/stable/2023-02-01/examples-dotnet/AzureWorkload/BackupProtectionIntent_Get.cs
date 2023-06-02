using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.RecoveryServicesBackup;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-02-01/examples/AzureWorkload/BackupProtectionIntent_Get.json
// this example is just showing the usage of "ProtectionIntent_Get" operation, for the dependent resources, they will have to be created separately.

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
string intentObjectName = "249D9B07-D2EF-4202-AA64-65F35418564E";
ResourceIdentifier backupProtectionIntentResourceId = BackupProtectionIntentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, fabricName, intentObjectName);
BackupProtectionIntentResource backupProtectionIntent = client.GetBackupProtectionIntentResource(backupProtectionIntentResourceId);

// invoke the operation
BackupProtectionIntentResource result = await backupProtectionIntent.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupProtectionIntentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

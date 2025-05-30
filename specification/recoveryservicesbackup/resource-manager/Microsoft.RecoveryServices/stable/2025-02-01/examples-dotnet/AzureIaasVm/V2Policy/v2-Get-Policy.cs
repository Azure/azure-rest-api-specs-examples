using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServicesBackup.Models;
using Azure.ResourceManager.RecoveryServicesBackup;

// Generated from example definition: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/AzureIaasVm/V2Policy/v2-Get-Policy.json
// this example is just showing the usage of "ProtectionPolicies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BackupProtectionPolicyResource created on azure
// for more information of creating BackupProtectionPolicyResource, please refer to the document of BackupProtectionPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "SwaggerTestRg";
string vaultName = "NetSDKTestRsVault";
string policyName = "v2-daily-sample";
ResourceIdentifier backupProtectionPolicyResourceId = BackupProtectionPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vaultName, policyName);
BackupProtectionPolicyResource backupProtectionPolicy = client.GetBackupProtectionPolicyResource(backupProtectionPolicyResourceId);

// invoke the operation
BackupProtectionPolicyResource result = await backupProtectionPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BackupProtectionPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

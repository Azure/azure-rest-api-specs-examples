using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/BackupPolicies_Delete.json
// this example is just showing the usage of "BackupPolicies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppBackupPolicyResource created on azure
// for more information of creating NetAppBackupPolicyResource, please refer to the document of NetAppBackupPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resourceGroup";
string accountName = "accountName";
string backupPolicyName = "backupPolicyName";
ResourceIdentifier netAppBackupPolicyResourceId = NetAppBackupPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, backupPolicyName);
NetAppBackupPolicyResource netAppBackupPolicy = client.GetNetAppBackupPolicyResource(netAppBackupPolicyResourceId);

// invoke the operation
await netAppBackupPolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

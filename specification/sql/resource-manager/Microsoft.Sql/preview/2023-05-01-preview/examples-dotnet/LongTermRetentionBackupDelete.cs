using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/LongTermRetentionBackupDelete.json
// this example is just showing the usage of "LongTermRetentionBackups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionLongTermRetentionBackupResource created on azure
// for more information of creating SubscriptionLongTermRetentionBackupResource, please refer to the document of SubscriptionLongTermRetentionBackupResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
AzureLocation locationName = new AzureLocation("japaneast");
string longTermRetentionServerName = "testserver";
string longTermRetentionDatabaseName = "testDatabase";
string backupName = "55555555-6666-7777-8888-999999999999;131637960820000000;Hot";
ResourceIdentifier subscriptionLongTermRetentionBackupResourceId = SubscriptionLongTermRetentionBackupResource.CreateResourceIdentifier(subscriptionId, locationName, longTermRetentionServerName, longTermRetentionDatabaseName, backupName);
SubscriptionLongTermRetentionBackupResource subscriptionLongTermRetentionBackup = client.GetSubscriptionLongTermRetentionBackupResource(subscriptionLongTermRetentionBackupResourceId);

// invoke the operation
await subscriptionLongTermRetentionBackup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

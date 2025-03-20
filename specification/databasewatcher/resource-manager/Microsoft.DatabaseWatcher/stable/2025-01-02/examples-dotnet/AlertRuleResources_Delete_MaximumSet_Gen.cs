using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DatabaseWatcher.Models;
using Azure.ResourceManager.DatabaseWatcher;

// Generated from example definition: 2025-01-02/AlertRuleResources_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "AlertRuleResource_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatabaseWatcherAlertRuleResource created on azure
// for more information of creating DatabaseWatcherAlertRuleResource, please refer to the document of DatabaseWatcherAlertRuleResource
string subscriptionId = "A76F9850-996B-40B3-94D4-C98110A0EEC9";
string resourceGroupName = "rgWatcher";
string watcherName = "testWatcher";
string alertRuleResourceName = "testAlert";
ResourceIdentifier databaseWatcherAlertRuleResourceId = DatabaseWatcherAlertRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, watcherName, alertRuleResourceName);
DatabaseWatcherAlertRuleResource databaseWatcherAlertRule = client.GetDatabaseWatcherAlertRuleResource(databaseWatcherAlertRuleResourceId);

// invoke the operation
await databaseWatcherAlertRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

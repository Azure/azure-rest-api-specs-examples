using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automation.Models;
using Azure.ResourceManager.Automation;

// Generated from example definition: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/deleteWatcher.json
// this example is just showing the usage of "Watcher_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomationWatcherResource created on azure
// for more information of creating AutomationWatcherResource, please refer to the document of AutomationWatcherResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string automationAccountName = "MyTestAutomationAccount";
string watcherName = "MyTestWatcher";
ResourceIdentifier automationWatcherResourceId = AutomationWatcherResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, automationAccountName, watcherName);
AutomationWatcherResource automationWatcher = client.GetAutomationWatcherResource(automationWatcherResourceId);

// invoke the operation
await automationWatcher.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

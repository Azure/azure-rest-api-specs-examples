using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Elastic;

// Generated from example definition: specification/elastic/resource-manager/Microsoft.Elastic/preview/2020-07-01-preview/examples/TagRules_Delete.json
// this example is just showing the usage of "TagRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MonitoringTagRuleResource created on azure
// for more information of creating MonitoringTagRuleResource, please refer to the document of MonitoringTagRuleResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
string ruleSetName = "default";
ResourceIdentifier monitoringTagRuleResourceId = MonitoringTagRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName, ruleSetName);
MonitoringTagRuleResource monitoringTagRule = client.GetMonitoringTagRuleResource(monitoringTagRuleResourceId);

// invoke the operation
await monitoringTagRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/patchScheduledQueryRule.json
// this example is just showing the usage of "ScheduledQueryRules_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScheduledQueryRuleResource created on azure
// for more information of creating ScheduledQueryRuleResource, please refer to the document of ScheduledQueryRuleResource
string subscriptionId = "dd4bfc94-a096-412b-9c43-4bd13e35afbc";
string resourceGroupName = "QueryResourceGroupName";
string ruleName = "heartbeat";
ResourceIdentifier scheduledQueryRuleResourceId = ScheduledQueryRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ruleName);
ScheduledQueryRuleResource scheduledQueryRule = client.GetScheduledQueryRuleResource(scheduledQueryRuleResourceId);

// invoke the operation
ScheduledQueryRulePatch patch = new ScheduledQueryRulePatch()
{
    IsEnabled = false,
};
ScheduledQueryRuleResource result = await scheduledQueryRule.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScheduledQueryRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Datadog.Models;
using Azure.ResourceManager.Datadog;

// Generated from example definition: specification/datadog/resource-manager/Microsoft.Datadog/stable/2025-06-11/examples/TagRules_CreateOrUpdate.json
// this example is just showing the usage of "TagRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataMonitoringTagRuleResource created on azure
// for more information of creating DataMonitoringTagRuleResource, please refer to the document of DataMonitoringTagRuleResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
string ruleSetName = "default";
ResourceIdentifier dataMonitoringTagRuleResourceId = DataMonitoringTagRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName, ruleSetName);
DataMonitoringTagRuleResource dataMonitoringTagRule = client.GetDataMonitoringTagRuleResource(dataMonitoringTagRuleResourceId);

// invoke the operation
DataMonitoringTagRuleData data = new DataMonitoringTagRuleData
{
    Properties = new MonitoringTagRuleProperties
    {
        LogRules = new DatadogMonitorLogRules
        {
            IsAadLogsSent = false,
            IsSubscriptionLogsSent = true,
            IsResourceLogsSent = true,
            FilteringTags = {new DatadogMonitorFilteringTag
            {
            Name = "Environment",
            Value = "Prod",
            Action = DatadogMonitorTagAction.Include,
            }, new DatadogMonitorFilteringTag
            {
            Name = "Environment",
            Value = "Dev",
            Action = DatadogMonitorTagAction.Exclude,
            }},
        },
        MetricRulesFilteringTags = { },
        IsAutomutingEnabled = true,
    },
};
ArmOperation<DataMonitoringTagRuleResource> lro = await dataMonitoringTagRule.UpdateAsync(WaitUntil.Completed, data);
DataMonitoringTagRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataMonitoringTagRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

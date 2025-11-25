using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Datadog.Models;
using Azure.ResourceManager.Datadog;

// Generated from example definition: specification/datadog/resource-manager/Microsoft.Datadog/stable/2025-06-11/examples/MonitoredSubscriptions_Update.json
// this example is just showing the usage of "MonitoredSubscriptions_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatadogMonitoredSubscriptionResource created on azure
// for more information of creating DatadogMonitoredSubscriptionResource, please refer to the document of DatadogMonitoredSubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
string configurationName = "default";
ResourceIdentifier datadogMonitoredSubscriptionResourceId = DatadogMonitoredSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName, configurationName);
DatadogMonitoredSubscriptionResource datadogMonitoredSubscription = client.GetDatadogMonitoredSubscriptionResource(datadogMonitoredSubscriptionResourceId);

// invoke the operation
DatadogMonitoredSubscriptionData data = new DatadogMonitoredSubscriptionData
{
    Properties = new DatadogSubscriptionProperties
    {
        Operation = DatadogOperation.AddComplete,
        MonitoredSubscriptionList = {new DatadogMonitoredSubscriptionItem
        {
        SubscriptionId = "/subscriptions/00000000-0000-0000-0000-000000000000",
        Status = DatadogMonitorStatus.Active,
        TagRules = new MonitoringTagRuleProperties
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
        MetricRulesFilteringTags = {},
        IsAutomutingEnabled = true,
        },
        }, new DatadogMonitoredSubscriptionItem
        {
        SubscriptionId = "/subscriptions/00000000-0000-0000-0000-000000000001",
        Status = DatadogMonitorStatus.Failed,
        TagRules = new MonitoringTagRuleProperties
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
        MetricRulesFilteringTags = {},
        IsAutomutingEnabled = true,
        },
        }},
    },
};
ArmOperation<DatadogMonitoredSubscriptionResource> lro = await datadogMonitoredSubscription.UpdateAsync(WaitUntil.Completed, data);
DatadogMonitoredSubscriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatadogMonitoredSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

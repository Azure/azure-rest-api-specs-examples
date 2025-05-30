using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/UpdateMetricAlert.json
// this example is just showing the usage of "MetricAlerts_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MetricAlertResource created on azure
// for more information of creating MetricAlertResource, please refer to the document of MetricAlertResource
string subscriptionId = "14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7";
string resourceGroupName = "gigtest";
string ruleName = "chiricutin";
ResourceIdentifier metricAlertResourceId = MetricAlertResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ruleName);
MetricAlertResource metricAlert = client.GetMetricAlertResource(metricAlertResourceId);

// invoke the operation
MetricAlertPatch patch = new MetricAlertPatch
{
    Tags = { },
    Description = "This is the description of the rule1",
    Severity = 3,
    IsEnabled = true,
    Scopes = { "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme" },
    EvaluationFrequency = XmlConvert.ToTimeSpan("PT1M"),
    WindowSize = XmlConvert.ToTimeSpan("PT15M"),
    Criteria = new MetricAlertSingleResourceMultipleMetricCriteria
    {
        AllOf = {new MetricCriteria("High_CPU_80", "\\Processor(_Total)\\% Processor Time", MetricCriteriaTimeAggregationType.Average, MetricCriteriaOperator.GreaterThan, 80.5)
        {
        Dimensions = {},
        }},
    },
    IsAutoMitigateEnabled = true,
    Actions = {new MetricAlertAction
    {
    ActionGroupId = new ResourceIdentifier("/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
    WebHookProperties =
    {
    ["key11"] = "value11",
    ["key12"] = "value12"
    },
    }},
};
MetricAlertResource result = await metricAlert.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MetricAlertData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

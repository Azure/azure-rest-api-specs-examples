using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Monitor;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/createOrUpdateDynamicMetricAlertSingleResource.json
// this example is just showing the usage of "MetricAlerts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "gigtest";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MetricAlertResource
MetricAlertCollection collection = resourceGroupResource.GetMetricAlerts();

// invoke the operation
string ruleName = "chiricutin";
MetricAlertData data = new MetricAlertData(new AzureLocation("global"), 3, true, new string[]
{
"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme"
}, XmlConvert.ToTimeSpan("PT1M"), XmlConvert.ToTimeSpan("PT15M"), new MetricAlertMultipleResourceMultipleMetricCriteria()
{
    AllOf =
{
new DynamicMetricCriteria("High_CPU_80","Percentage CPU",MetricCriteriaTimeAggregationType.Average,DynamicThresholdOperator.GreaterOrLessThan,DynamicThresholdSensitivity.Medium,new DynamicThresholdFailingPeriods(4,4))
{
IgnoreDataBefore = DateTimeOffset.Parse("2019-04-04T21:00:00.000Z"),
MetricNamespace = "microsoft.compute/virtualmachines",
Dimensions =
{
},
}
},
})
{
    Description = "This is the description of the rule1",
    IsAutoMitigateEnabled = true,
    Actions =
    {
    new MetricAlertAction()
    {
    ActionGroupId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/actiongroups/group2"),
    WebHookProperties =
    {
    ["key11"] = "value11",
    ["key12"] = "value12",
    },
    }
    },
    Tags =
    {
    },
};
ArmOperation<MetricAlertResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ruleName, data);
MetricAlertResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MetricAlertData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

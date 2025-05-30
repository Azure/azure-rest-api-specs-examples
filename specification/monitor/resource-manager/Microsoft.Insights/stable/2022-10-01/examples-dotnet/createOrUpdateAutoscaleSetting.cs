using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/createOrUpdateAutoscaleSetting.json
// this example is just showing the usage of "AutoscaleSettings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
string resourceGroupName = "TestingMetricsScaleSet";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AutoscaleSettingResource
AutoscaleSettingCollection collection = resourceGroupResource.GetAutoscaleSettings();

// invoke the operation
string autoscaleSettingName = "MySetting";
AutoscaleSettingData data = new AutoscaleSettingData(new AzureLocation("West US"), new AutoscaleProfile[]
{
new AutoscaleProfile("adios", new MonitorScaleCapacity(1, 10, 1), new AutoscaleRule[]
{
new AutoscaleRule(new MetricTrigger(
    "Percentage CPU",
    new ResourceIdentifier("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
    XmlConvert.ToTimeSpan("PT1M"),
    MetricStatisticType.Average,
    XmlConvert.ToTimeSpan("PT5M"),
    MetricTriggerTimeAggregationType.Average,
    MetricTriggerComparisonOperation.GreaterThan,
    10)
{
IsDividedPerInstance = false,
}, new MonitorScaleAction(MonitorScaleDirection.Increase, MonitorScaleType.ChangeCount, XmlConvert.ToTimeSpan("PT5M"))
{
Value = "1",
}),
new AutoscaleRule(new MetricTrigger(
    "Percentage CPU",
    new ResourceIdentifier("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
    XmlConvert.ToTimeSpan("PT2M"),
    MetricStatisticType.Average,
    XmlConvert.ToTimeSpan("PT5M"),
    MetricTriggerTimeAggregationType.Average,
    MetricTriggerComparisonOperation.GreaterThan,
    15)
{
IsDividedPerInstance = false,
}, new MonitorScaleAction(MonitorScaleDirection.Decrease, MonitorScaleType.ChangeCount, XmlConvert.ToTimeSpan("PT6M"))
{
Value = "2",
})
})
{
FixedDate = new MonitorTimeWindow(DateTimeOffset.Parse("2015-03-05T14:00:00Z"), DateTimeOffset.Parse("2015-03-05T14:30:00Z"))
{
TimeZone = "UTC",
},
},
new AutoscaleProfile("saludos", new MonitorScaleCapacity(1, 10, 1), new AutoscaleRule[]
{
new AutoscaleRule(new MetricTrigger(
    "Percentage CPU",
    new ResourceIdentifier("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
    XmlConvert.ToTimeSpan("PT1M"),
    MetricStatisticType.Average,
    XmlConvert.ToTimeSpan("PT5M"),
    MetricTriggerTimeAggregationType.Average,
    MetricTriggerComparisonOperation.GreaterThan,
    10)
{
IsDividedPerInstance = false,
}, new MonitorScaleAction(MonitorScaleDirection.Increase, MonitorScaleType.ChangeCount, XmlConvert.ToTimeSpan("PT5M"))
{
Value = "1",
}),
new AutoscaleRule(new MetricTrigger(
    "Percentage CPU",
    new ResourceIdentifier("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
    XmlConvert.ToTimeSpan("PT2M"),
    MetricStatisticType.Average,
    XmlConvert.ToTimeSpan("PT5M"),
    MetricTriggerTimeAggregationType.Average,
    MetricTriggerComparisonOperation.GreaterThan,
    15)
{
IsDividedPerInstance = false,
}, new MonitorScaleAction(MonitorScaleDirection.Decrease, MonitorScaleType.ChangeCount, XmlConvert.ToTimeSpan("PT6M"))
{
Value = "2",
})
})
{
Recurrence = new MonitorRecurrence(RecurrenceFrequency.Week, new RecurrentSchedule("UTC", new MonitorDayOfWeek[]{new MonitorDayOfWeek("1")}, new int[]{5}, new int[]{15})),
}
})
{
    Notifications = {new AutoscaleNotification
    {
    Email = new EmailNotification
    {
    SendToSubscriptionAdministrator = true,
    SendToSubscriptionCoAdministrators = true,
    CustomEmails = {"gu@ms.com", "ge@ns.net"},
    },
    Webhooks = {new WebhookNotification
    {
    ServiceUri = new Uri("http://myservice.com"),
    Properties = {},
    }},
    }},
    IsEnabled = true,
    PredictiveAutoscalePolicy = new PredictiveAutoscalePolicy(PredictiveAutoscalePolicyScaleMode.Enabled),
    TargetResourceId = new ResourceIdentifier("/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/TestingMetricsScaleSet/providers/Microsoft.Compute/virtualMachineScaleSets/testingsc"),
    Tags =
    {
    ["key1"] = "value1",
    ["key2"] = "value2"
    },
};
ArmOperation<AutoscaleSettingResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, autoscaleSettingName, data);
AutoscaleSettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AutoscaleSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

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

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/createOrUpdateScheduledQueryRuleSubscription.json
// this example is just showing the usage of "ScheduledQueryRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "dd4bfc94-a096-412b-9c43-4bd13e35afbc";
string resourceGroupName = "QueryResourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ScheduledQueryRuleResource
ScheduledQueryRuleCollection collection = resourceGroupResource.GetScheduledQueryRules();

// invoke the operation
string ruleName = "perf";
ScheduledQueryRuleData data = new ScheduledQueryRuleData(new AzureLocation("eastus"))
{
    Description = "Performance rule",
    Severity = AlertSeverity.Four,
    IsEnabled = true,
    Scopes =
    {
    "/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147"
    },
    EvaluationFrequency = XmlConvert.ToTimeSpan("PT5M"),
    WindowSize = XmlConvert.ToTimeSpan("PT10M"),
    TargetResourceTypes =
    {
    "Microsoft.Compute/virtualMachines"
    },
    CriteriaAllOf =
    {
    new ScheduledQueryRuleCondition()
    {
    Query = "Perf | where ObjectName == \"Processor\"",
    TimeAggregation = ScheduledQueryRuleTimeAggregationType.Average,
    MetricMeasureColumn = "% Processor Time",
    ResourceIdColumn = "resourceId",
    Dimensions =
    {
    new MonitorDimension("ComputerIp",MonitorDimensionOperator.Exclude,new string[]
    {
    "192.168.1.1"
    }),new MonitorDimension("OSType",MonitorDimensionOperator.Include,new string[]
    {
    "*"
    })
    },
    Operator = MonitorConditionOperator.GreaterThan,
    Threshold = 70,
    FailingPeriods = new ConditionFailingPeriods()
    {
    NumberOfEvaluationPeriods = 1,
    MinFailingPeriodsToAlert = 1,
    },
    }
    },
    MuteActionsDuration = XmlConvert.ToTimeSpan("PT30M"),
    Actions = new ScheduledQueryRuleActions()
    {
        ActionGroups =
        {
        "/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup"
        },
        CustomProperties =
        {
        ["key11"] = "value11",
        ["key12"] = "value12",
        },
    },
    CheckWorkspaceAlertsStorageConfigured = true,
    SkipQueryValidation = true,
    AutoMitigate = true,
};
ArmOperation<ScheduledQueryRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ruleName, data);
ScheduledQueryRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScheduledQueryRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

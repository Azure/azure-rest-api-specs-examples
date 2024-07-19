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

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2022-10-01/examples/GetPredictiveMetric.json
// this example is just showing the usage of "PredictiveMetric_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutoscaleSettingResource created on azure
// for more information of creating AutoscaleSettingResource, please refer to the document of AutoscaleSettingResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string autoscaleSettingName = "vmss1-Autoscale-775";
ResourceIdentifier autoscaleSettingResourceId = AutoscaleSettingResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, autoscaleSettingName);
AutoscaleSettingResource autoscaleSetting = client.GetAutoscaleSettingResource(autoscaleSettingResourceId);

// invoke the operation
string timespan = "2021-10-14T22:00:00.000Z/2021-10-16T22:00:00.000Z";
TimeSpan interval = XmlConvert.ToTimeSpan("PT1H");
string metricNamespace = "Microsoft.Compute/virtualMachineScaleSets";
string metricName = "PercentageCPU";
string aggregation = "Total";
AutoscaleSettingPredicativeResult result = await autoscaleSetting.GetPredictiveMetricAsync(timespan, interval, metricNamespace, metricName, aggregation);

Console.WriteLine($"Succeeded: {result}");

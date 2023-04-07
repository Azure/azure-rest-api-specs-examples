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

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getDynamicMetricAlertMultipleResource.json
// this example is just showing the usage of "MetricAlerts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MetricAlertResource created on azure
// for more information of creating MetricAlertResource, please refer to the document of MetricAlertResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "gigtest";
string ruleName = "MetricAlertOnMultipleResources";
ResourceIdentifier metricAlertResourceId = MetricAlertResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ruleName);
MetricAlertResource metricAlert = client.GetMetricAlertResource(metricAlertResourceId);

// invoke the operation
MetricAlertResource result = await metricAlert.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MetricAlertData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

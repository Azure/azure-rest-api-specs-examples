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

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/createOrUpdateWebTestMetricAlert.json
// this example is just showing the usage of "MetricAlerts_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "12345678-1234-1234-1234-123456789101";
string resourceGroupName = "rg-example";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this MetricAlertResource
MetricAlertCollection collection = resourceGroupResource.GetMetricAlerts();

// invoke the operation
string ruleName = "webtest-name-example";
MetricAlertData data = new MetricAlertData(new AzureLocation("global"), 4, true, new string[]
{
"/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example","/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example"
}, XmlConvert.ToTimeSpan("PT1M"), XmlConvert.ToTimeSpan("PT15M"), new WebtestLocationAvailabilityCriteria(new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example"), new ResourceIdentifier("/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example"), 2))
{
    Description = "Automatically created alert rule for availability test \"component-example\" a",
    Actions =
    {
    },
    Tags =
    {
    ["hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example"] = "Resource",
    ["hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example"] = "Resource",
    },
};
ArmOperation<MetricAlertResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ruleName, data);
MetricAlertResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MetricAlertData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

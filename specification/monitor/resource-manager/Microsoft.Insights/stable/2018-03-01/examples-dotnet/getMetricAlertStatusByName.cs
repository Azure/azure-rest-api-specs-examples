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

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getMetricAlertStatusByName.json
// this example is just showing the usage of "MetricAlertsStatus_ListByName" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MetricAlertResource created on azure
// for more information of creating MetricAlertResource, please refer to the document of MetricAlertResource
string subscriptionId = "009f6022-67ec-423e-9aa7-691182870588";
string resourceGroupName = "EastUs";
string ruleName = "custom1";
ResourceIdentifier metricAlertResourceId = MetricAlertResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ruleName);
MetricAlertResource metricAlert = client.GetMetricAlertResource(metricAlertResourceId);

// invoke the operation and iterate over the result
string statusName = "cmVzb3VyY2VJZD0vc3Vic2NyaXB0aW9ucy8xNGRkZjBjNS03N2M1LTRiNTMtODRmNi1lMWZhNDNhZDY4ZjcvcmVzb3VyY2VHcm91cHMvZ2lndGVzdC9wcm92aWRlcnMvTWljcm9zb2Z0LkNvbXB1dGUvdmlydHVhbE1hY2hpbmVzL2dpZ3dhZG1l";
await foreach (MetricAlertStatus item in metricAlert.GetAllMetricAlertsStatusByNameAsync(statusName))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Datadog;

// Generated from example definition: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/TagRules_Get.json
// this example is just showing the usage of "TagRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DatadogMonitorResource created on azure
// for more information of creating DatadogMonitorResource, please refer to the document of DatadogMonitorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
ResourceIdentifier datadogMonitorResourceId = DatadogMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
DatadogMonitorResource datadogMonitorResource = client.GetDatadogMonitorResource(datadogMonitorResourceId);

// get the collection of this MonitoringTagRuleResource
MonitoringTagRuleCollection collection = datadogMonitorResource.GetMonitoringTagRules();

// invoke the operation
string ruleSetName = "default";
NullableResponse<MonitoringTagRuleResource> response = await collection.GetIfExistsAsync(ruleSetName);
MonitoringTagRuleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MonitoringTagRuleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

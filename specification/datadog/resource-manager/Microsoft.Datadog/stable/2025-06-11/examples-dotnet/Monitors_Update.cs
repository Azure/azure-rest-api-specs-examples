using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Datadog.Models;
using Azure.ResourceManager.Datadog;

// Generated from example definition: specification/datadog/resource-manager/Microsoft.Datadog/stable/2025-06-11/examples/Monitors_Update.json
// this example is just showing the usage of "Monitors_Update" operation, for the dependent resources, they will have to be created separately.

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
DatadogMonitorResource datadogMonitor = client.GetDatadogMonitorResource(datadogMonitorResourceId);

// invoke the operation
DatadogMonitorPatch patch = new DatadogMonitorPatch
{
    Properties = new DatadogMonitorResourcePatchProperties
    {
        MonitoringStatus = DatadogMonitoringStatus.Enabled,
    },
    Tags =
    {
    ["Environment"] = "Dev"
    },
};
ArmOperation<DatadogMonitorResource> lro = await datadogMonitor.UpdateAsync(WaitUntil.Completed, patch);
DatadogMonitorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DatadogMonitorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

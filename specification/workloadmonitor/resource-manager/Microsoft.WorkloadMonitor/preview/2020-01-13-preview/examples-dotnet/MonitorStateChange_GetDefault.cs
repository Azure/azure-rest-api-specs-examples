using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadMonitor;

// Generated from example definition: specification/workloadmonitor/resource-manager/Microsoft.WorkloadMonitor/preview/2020-01-13-preview/examples/MonitorStateChange_GetDefault.json
// this example is just showing the usage of "HealthMonitors_GetStateChange" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthMonitorResource created on azure
// for more information of creating HealthMonitorResource, please refer to the document of HealthMonitorResource
string subscriptionId = "bc27da3b-3ba2-4e00-a6ec-1fde64aa1e21";
string resourceGroupName = "tugamidiAlerts";
string providerName = "Microsoft.Compute";
string resourceCollectionName = "virtualMachines";
string resourceName = "linuxEUS";
string monitorId = "logical-disks|C@3A";
ResourceIdentifier healthMonitorResourceId = HealthMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorId);
HealthMonitorResource healthMonitor = client.GetHealthMonitorResource(healthMonitorResourceId);

// get the collection of this HealthMonitorStateChangeResource
HealthMonitorStateChangeCollection collection = healthMonitor.GetHealthMonitorStateChanges();

// invoke the operation
string timestampUnix = "1584316800";
NullableResponse<HealthMonitorStateChangeResource> response = await collection.GetIfExistsAsync(timestampUnix);
HealthMonitorStateChangeResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HealthMonitorStateChangeData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

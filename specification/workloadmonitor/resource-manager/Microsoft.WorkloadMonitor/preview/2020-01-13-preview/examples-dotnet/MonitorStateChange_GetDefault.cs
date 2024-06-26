using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.WorkloadMonitor;

// Generated from example definition: specification/workloadmonitor/resource-manager/Microsoft.WorkloadMonitor/preview/2020-01-13-preview/examples/MonitorStateChange_GetDefault.json
// this example is just showing the usage of "HealthMonitors_GetStateChange" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthMonitorStateChangeResource created on azure
// for more information of creating HealthMonitorStateChangeResource, please refer to the document of HealthMonitorStateChangeResource
string subscriptionId = "bc27da3b-3ba2-4e00-a6ec-1fde64aa1e21";
string resourceGroupName = "tugamidiAlerts";
string providerName = "Microsoft.Compute";
string resourceCollectionName = "virtualMachines";
string resourceName = "linuxEUS";
string monitorId = "logical-disks|C@3A";
string timestampUnix = "1584316800";
ResourceIdentifier healthMonitorStateChangeResourceId = HealthMonitorStateChangeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorId, timestampUnix);
HealthMonitorStateChangeResource healthMonitorStateChange = client.GetHealthMonitorStateChangeResource(healthMonitorStateChangeResourceId);

// invoke the operation
HealthMonitorStateChangeResource result = await healthMonitorStateChange.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HealthMonitorStateChangeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

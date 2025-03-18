using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.WorkloadMonitor;

// Generated from example definition: specification/workloadmonitor/resource-manager/Microsoft.WorkloadMonitor/preview/2020-01-13-preview/examples/Monitor_GetDefault.json
// this example is just showing the usage of "HealthMonitors_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "bc27da3b-3ba2-4e00-a6ec-1fde64aa1e21";
string resourceGroupName = "tugamidiAlerts";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this HealthMonitorResource
string providerName = "Microsoft.Compute";
string resourceCollectionName = "virtualMachines";
string resourceName = "linuxEUS";
HealthMonitorCollection collection = resourceGroupResource.GetHealthMonitors(providerName, resourceCollectionName, resourceName);

// invoke the operation
string monitorId = "logical-disks|C@3A|free-space";
NullableResponse<HealthMonitorResource> response = await collection.GetIfExistsAsync(monitorId);
HealthMonitorResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    HealthMonitorData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

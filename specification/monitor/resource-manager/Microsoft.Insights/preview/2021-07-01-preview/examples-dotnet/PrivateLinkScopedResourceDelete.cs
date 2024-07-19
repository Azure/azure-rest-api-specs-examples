using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopedResourceDelete.json
// this example is just showing the usage of "PrivateLinkScopedResources_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MonitorPrivateLinkScopedResource created on azure
// for more information of creating MonitorPrivateLinkScopedResource, please refer to the document of MonitorPrivateLinkScopedResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "MyResourceGroup";
string scopeName = "MyPrivateLinkScope";
string name = "scoped-resource-name";
ResourceIdentifier monitorPrivateLinkScopedResourceId = MonitorPrivateLinkScopedResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scopeName, name);
MonitorPrivateLinkScopedResource monitorPrivateLinkScopedResource = client.GetMonitorPrivateLinkScopedResource(monitorPrivateLinkScopedResourceId);

// invoke the operation
await monitorPrivateLinkScopedResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

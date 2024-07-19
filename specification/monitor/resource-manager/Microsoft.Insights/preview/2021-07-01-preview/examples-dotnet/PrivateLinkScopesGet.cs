using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/preview/2021-07-01-preview/examples/PrivateLinkScopesGet.json
// this example is just showing the usage of "PrivateLinkScopes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MonitorPrivateLinkScopeResource created on azure
// for more information of creating MonitorPrivateLinkScopeResource, please refer to the document of MonitorPrivateLinkScopeResource
string subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
string resourceGroupName = "my-resource-group";
string scopeName = "my-privatelinkscope";
ResourceIdentifier monitorPrivateLinkScopeResourceId = MonitorPrivateLinkScopeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scopeName);
MonitorPrivateLinkScopeResource monitorPrivateLinkScope = client.GetMonitorPrivateLinkScopeResource(monitorPrivateLinkScopeResourceId);

// invoke the operation
MonitorPrivateLinkScopeResource result = await monitorPrivateLinkScope.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MonitorPrivateLinkScopeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

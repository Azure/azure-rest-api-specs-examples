using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Workloads;
using Azure.ResourceManager.Workloads.Models;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/workloadmonitor/SapLandscapeMonitor_Delete.json
// this example is just showing the usage of "SapLandscapeMonitor_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapLandscapeMonitorResource created on azure
// for more information of creating SapLandscapeMonitorResource, please refer to the document of SapLandscapeMonitorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "mySapMonitor";
ResourceIdentifier sapLandscapeMonitorResourceId = SapLandscapeMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
SapLandscapeMonitorResource sapLandscapeMonitor = client.GetSapLandscapeMonitorResource(sapLandscapeMonitorResourceId);

// invoke the operation
await sapLandscapeMonitor.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

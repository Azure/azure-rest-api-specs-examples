using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Monitor/preview/2023-10-01-preview/examples/AzureMonitorWorkspaceGet.json
// this example is just showing the usage of "AzureMonitorWorkspaces_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MonitorWorkspaceResource created on azure
// for more information of creating MonitorWorkspaceResource, please refer to the document of MonitorWorkspaceResource
string subscriptionId = "703362b3-f278-4e4b-9179-c76eaf41ffc2";
string resourceGroupName = "myResourceGroup";
string azureMonitorWorkspaceName = "myAzureMonitorWorkspace";
ResourceIdentifier monitorWorkspaceResourceId = MonitorWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureMonitorWorkspaceName);
MonitorWorkspaceResource monitorWorkspaceResource = client.GetMonitorWorkspaceResource(monitorWorkspaceResourceId);

// invoke the operation
MonitorWorkspaceResource result = await monitorWorkspaceResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MonitorWorkspaceResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

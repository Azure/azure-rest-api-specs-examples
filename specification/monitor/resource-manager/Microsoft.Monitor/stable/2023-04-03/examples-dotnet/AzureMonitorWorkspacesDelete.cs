using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Monitor/stable/2023-04-03/examples/AzureMonitorWorkspacesDelete.json
// this example is just showing the usage of "AzureMonitorWorkspaces_Delete" operation, for the dependent resources, they will have to be created separately.

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
await monitorWorkspaceResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

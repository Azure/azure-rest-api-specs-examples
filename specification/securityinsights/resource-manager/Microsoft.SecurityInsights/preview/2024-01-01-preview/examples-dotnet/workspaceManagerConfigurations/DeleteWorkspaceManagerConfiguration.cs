using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/workspaceManagerConfigurations/DeleteWorkspaceManagerConfiguration.json
// this example is just showing the usage of "WorkspaceManagerConfigurations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkspaceManagerConfigurationResource created on azure
// for more information of creating WorkspaceManagerConfigurationResource, please refer to the document of WorkspaceManagerConfigurationResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
string workspaceManagerConfigurationName = "default";
ResourceIdentifier workspaceManagerConfigurationResourceId = WorkspaceManagerConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, workspaceManagerConfigurationName);
WorkspaceManagerConfigurationResource workspaceManagerConfiguration = client.GetWorkspaceManagerConfigurationResource(workspaceManagerConfigurationResourceId);

// invoke the operation
await workspaceManagerConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

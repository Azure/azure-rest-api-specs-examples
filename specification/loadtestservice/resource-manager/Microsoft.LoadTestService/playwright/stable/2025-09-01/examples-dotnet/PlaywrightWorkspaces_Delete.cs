using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Playwright.Models;
using Azure.ResourceManager.Playwright;

// Generated from example definition: 2025-09-01/PlaywrightWorkspaces_Delete.json
// this example is just showing the usage of "PlaywrightWorkspace_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PlaywrightWorkspaceResource created on azure
// for more information of creating PlaywrightWorkspaceResource, please refer to the document of PlaywrightWorkspaceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "dummyrg";
string playwrightWorkspaceName = "myWorkspace";
ResourceIdentifier playwrightWorkspaceResourceId = PlaywrightWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, playwrightWorkspaceName);
PlaywrightWorkspaceResource playwrightWorkspace = client.GetPlaywrightWorkspaceResource(playwrightWorkspaceResourceId);

// invoke the operation
await playwrightWorkspace.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

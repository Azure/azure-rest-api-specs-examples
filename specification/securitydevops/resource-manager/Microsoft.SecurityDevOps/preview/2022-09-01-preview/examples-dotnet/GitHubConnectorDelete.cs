using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityDevOps;
using Azure.ResourceManager.SecurityDevOps.Models;

// Generated from example definition: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorDelete.json
// this example is just showing the usage of "GitHubConnector_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GitHubConnectorResource created on azure
// for more information of creating GitHubConnectorResource, please refer to the document of GitHubConnectorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "westusrg";
string gitHubConnectorName = "testconnector";
ResourceIdentifier gitHubConnectorResourceId = GitHubConnectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gitHubConnectorName);
GitHubConnectorResource gitHubConnector = client.GetGitHubConnectorResource(gitHubConnectorResourceId);

// invoke the operation
await gitHubConnector.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityDevOps;

// Generated from example definition: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoUpdate.json
// this example is just showing the usage of "GitHubRepo_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GitHubRepoResource created on azure
// for more information of creating GitHubRepoResource, please refer to the document of GitHubRepoResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "westusrg";
string gitHubConnectorName = "testconnector";
string gitHubOwnerName = "Azure";
string gitHubRepoName = "azure-rest-api-specs";
ResourceIdentifier gitHubRepoResourceId = GitHubRepoResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gitHubConnectorName, gitHubOwnerName, gitHubRepoName);
GitHubRepoResource gitHubRepo = client.GetGitHubRepoResource(gitHubRepoResourceId);

// invoke the operation
GitHubRepoData data = new GitHubRepoData();
ArmOperation<GitHubRepoResource> lro = await gitHubRepo.UpdateAsync(WaitUntil.Completed, data);
GitHubRepoResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GitHubRepoData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

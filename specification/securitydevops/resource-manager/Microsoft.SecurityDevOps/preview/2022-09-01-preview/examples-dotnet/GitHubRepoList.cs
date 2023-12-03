using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityDevOps;

// Generated from example definition: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubRepoList.json
// this example is just showing the usage of "GitHubRepo_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GitHubOwnerResource created on azure
// for more information of creating GitHubOwnerResource, please refer to the document of GitHubOwnerResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "westusrg";
string gitHubConnectorName = "testconnector";
string gitHubOwnerName = "Azure";
ResourceIdentifier gitHubOwnerResourceId = GitHubOwnerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, gitHubConnectorName, gitHubOwnerName);
GitHubOwnerResource gitHubOwner = client.GetGitHubOwnerResource(gitHubOwnerResourceId);

// get the collection of this GitHubRepoResource
GitHubRepoCollection collection = gitHubOwner.GetGitHubRepos();

// invoke the operation and iterate over the result
await foreach (GitHubRepoResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    GitHubRepoData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");

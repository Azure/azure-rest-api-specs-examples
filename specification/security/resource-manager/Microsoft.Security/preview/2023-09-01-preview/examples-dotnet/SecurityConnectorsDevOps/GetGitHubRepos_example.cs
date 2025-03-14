using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2023-09-01-preview/examples/SecurityConnectorsDevOps/GetGitHubRepos_example.json
// this example is just showing the usage of "GitHubRepos_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityConnectorGitHubOwnerResource created on azure
// for more information of creating SecurityConnectorGitHubOwnerResource, please refer to the document of SecurityConnectorGitHubOwnerResource
string subscriptionId = "0806e1cd-cfda-4ff8-b99c-2b0af42cffd3";
string resourceGroupName = "myRg";
string securityConnectorName = "mySecurityConnectorName";
string ownerName = "myGitHubOwner";
ResourceIdentifier securityConnectorGitHubOwnerResourceId = SecurityConnectorGitHubOwnerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, securityConnectorName, ownerName);
SecurityConnectorGitHubOwnerResource securityConnectorGitHubOwner = client.GetSecurityConnectorGitHubOwnerResource(securityConnectorGitHubOwnerResourceId);

// get the collection of this SecurityConnectorGitHubRepositoryResource
SecurityConnectorGitHubRepositoryCollection collection = securityConnectorGitHubOwner.GetSecurityConnectorGitHubRepositories();

// invoke the operation
string repoName = "myGitHubRepo";
NullableResponse<SecurityConnectorGitHubRepositoryResource> response = await collection.GetIfExistsAsync(repoName);
SecurityConnectorGitHubRepositoryResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SecurityConnectorGitHubRepositoryData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

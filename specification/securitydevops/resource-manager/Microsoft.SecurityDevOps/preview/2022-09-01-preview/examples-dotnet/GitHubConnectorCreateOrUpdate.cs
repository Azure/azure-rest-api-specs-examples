using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityDevOps;
using Azure.ResourceManager.SecurityDevOps.Models;

// Generated from example definition: specification/securitydevops/resource-manager/Microsoft.SecurityDevOps/preview/2022-09-01-preview/examples/GitHubConnectorCreateOrUpdate.json
// this example is just showing the usage of "GitHubConnector_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "westusrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this GitHubConnectorResource
GitHubConnectorCollection collection = resourceGroupResource.GetGitHubConnectors();

// invoke the operation
string gitHubConnectorName = "testconnector";
GitHubConnectorData data = new GitHubConnectorData(new AzureLocation("West US"))
{
    Properties = new GitHubConnectorProperties()
    {
        Code = "00000000000000000000",
    },
};
ArmOperation<GitHubConnectorResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, gitHubConnectorName, data);
GitHubConnectorResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GitHubConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

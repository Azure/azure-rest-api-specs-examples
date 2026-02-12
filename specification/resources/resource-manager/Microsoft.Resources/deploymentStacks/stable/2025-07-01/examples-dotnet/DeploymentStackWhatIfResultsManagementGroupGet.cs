using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups;
using Azure.ResourceManager.Resources.DeploymentStacks.Models;
using Azure.ResourceManager.Resources.DeploymentStacks;

// Generated from example definition: 2025-07-01/DeploymentStackWhatIfResultsManagementGroupGet.json
// this example is just showing the usage of "DeploymentStacksWhatIfResult_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupResource created on azure
// for more information of creating ManagementGroupResource, please refer to the document of ManagementGroupResource
string managementGroupId = "myMg";
ResourceIdentifier managementGroupResourceId = ManagementGroupResource.CreateResourceIdentifier(managementGroupId);

// get the collection of this DeploymentStackWhatIfResultResource
DeploymentStackWhatIfResultCollection collection = client.GetDeploymentStackWhatIfResults(managementGroupResourceId);

// invoke the operation
string deploymentStacksWhatIfResultName = "simpleDeploymentStackWhatIfResult";
NullableResponse<DeploymentStackWhatIfResultResource> response = await collection.GetIfExistsAsync(deploymentStacksWhatIfResultName);
DeploymentStackWhatIfResultResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DeploymentStackWhatIfResultData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

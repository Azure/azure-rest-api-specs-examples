using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/UpdateIoTSecuritySolution.json
// this example is just showing the usage of "IotSecuritySolution_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotSecuritySolutionResource created on azure
// for more information of creating IotSecuritySolutionResource, please refer to the document of IotSecuritySolutionResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "myRg";
string solutionName = "default";
ResourceIdentifier iotSecuritySolutionResourceId = IotSecuritySolutionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, solutionName);
IotSecuritySolutionResource iotSecuritySolution = client.GetIotSecuritySolutionResource(iotSecuritySolutionResourceId);

// invoke the operation
IotSecuritySolutionPatch patch = new IotSecuritySolutionPatch()
{
    UserDefinedResources = new UserDefinedResourcesProperties("where type != \"microsoft.devices/iothubs\" | where name contains \"v2\"", new string[]
{
"075423e9-7d33-4166-8bdf-3920b04e3735"
}),
    RecommendationsConfiguration =
    {
    new RecommendationConfigurationProperties(IotSecurityRecommendationType.IotOpenPorts,RecommendationConfigStatus.Disabled),new RecommendationConfigurationProperties(IotSecurityRecommendationType.IotSharedCredentials,RecommendationConfigStatus.Disabled)
    },
    Tags =
    {
    ["foo"] = "bar",
    },
};
IotSecuritySolutionResource result = await iotSecuritySolution.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotSecuritySolutionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

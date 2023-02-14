using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/CreateIoTSecuritySolution.json
// this example is just showing the usage of "IotSecuritySolution_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string resourceGroupName = "MyGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this IotSecuritySolutionResource
IotSecuritySolutionCollection collection = resourceGroupResource.GetIotSecuritySolutions();

// invoke the operation
string solutionName = "default";
IotSecuritySolutionData data = new IotSecuritySolutionData(new AzureLocation("East Us"))
{
    Workspace = "/subscriptions/c4930e90-cd72-4aa5-93e9-2d081d129569/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace1",
    DisplayName = "Solution Default",
    Status = SecuritySolutionStatus.Enabled,
    Export =
    {
    },
    DisabledDataSources =
    {
    },
    IotHubs =
    {
    "/subscriptions/075423e9-7d33-4166-8bdf-3920b04e3735/resourceGroups/myRg/providers/Microsoft.Devices/IotHubs/FirstIotHub"
    },
    UserDefinedResources = new UserDefinedResourcesProperties("where type != \"microsoft.devices/iothubs\" | where name contains \"iot\"", new string[]
{
"075423e9-7d33-4166-8bdf-3920b04e3735"
}),
    RecommendationsConfiguration =
    {
    new RecommendationConfigurationProperties(IotSecurityRecommendationType.IotOpenPorts,RecommendationConfigStatus.Disabled),new RecommendationConfigurationProperties(IotSecurityRecommendationType.IotSharedCredentials,RecommendationConfigStatus.Disabled)
    },
    UnmaskedIPLoggingStatus = UnmaskedIPLoggingStatus.Enabled,
    Tags =
    {
    },
};
ArmOperation<IotSecuritySolutionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, solutionName, data);
IotSecuritySolutionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotSecuritySolutionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

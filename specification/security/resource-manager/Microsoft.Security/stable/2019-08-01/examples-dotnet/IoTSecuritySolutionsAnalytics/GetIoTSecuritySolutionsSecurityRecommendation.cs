using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityRecommendation.json
// this example is just showing the usage of "IotSecuritySolutionsAnalyticsRecommendation_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotSecuritySolutionAnalyticsModelResource created on azure
// for more information of creating IotSecuritySolutionAnalyticsModelResource, please refer to the document of IotSecuritySolutionAnalyticsModelResource
string subscriptionId = "075423e9-7d33-4166-8bdf-3920b04e3735";
string resourceGroupName = "IoTEdgeResources";
string solutionName = "default";
ResourceIdentifier iotSecuritySolutionAnalyticsModelResourceId = IotSecuritySolutionAnalyticsModelResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, solutionName);
IotSecuritySolutionAnalyticsModelResource iotSecuritySolutionAnalyticsModel = client.GetIotSecuritySolutionAnalyticsModelResource(iotSecuritySolutionAnalyticsModelResourceId);

// get the collection of this IotSecurityAggregatedRecommendationResource
IotSecurityAggregatedRecommendationCollection collection = iotSecuritySolutionAnalyticsModel.GetIotSecurityAggregatedRecommendations();

// invoke the operation
string aggregatedRecommendationName = "OpenPortsOnDevice";
bool result = await collection.ExistsAsync(aggregatedRecommendationName);

Console.WriteLine($"Succeeded: {result}");

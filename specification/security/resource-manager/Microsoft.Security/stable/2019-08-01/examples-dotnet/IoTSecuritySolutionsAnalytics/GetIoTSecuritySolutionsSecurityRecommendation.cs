using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutionsAnalytics/GetIoTSecuritySolutionsSecurityRecommendation.json
// this example is just showing the usage of "IotSecuritySolutionsAnalyticsRecommendation_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotSecurityAggregatedRecommendationResource created on azure
// for more information of creating IotSecurityAggregatedRecommendationResource, please refer to the document of IotSecurityAggregatedRecommendationResource
string subscriptionId = "075423e9-7d33-4166-8bdf-3920b04e3735";
string resourceGroupName = "IoTEdgeResources";
string solutionName = "default";
string aggregatedRecommendationName = "OpenPortsOnDevice";
ResourceIdentifier iotSecurityAggregatedRecommendationResourceId = IotSecurityAggregatedRecommendationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, solutionName, aggregatedRecommendationName);
IotSecurityAggregatedRecommendationResource iotSecurityAggregatedRecommendation = client.GetIotSecurityAggregatedRecommendationResource(iotSecurityAggregatedRecommendationResourceId);

// invoke the operation
IotSecurityAggregatedRecommendationResource result = await iotSecurityAggregatedRecommendation.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotSecurityAggregatedRecommendationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

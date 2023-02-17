using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Advisor;

// Generated from example definition: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GetSuppressionDetail.json
// this example is just showing the usage of "Suppressions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceRecommendationBaseResource created on azure
// for more information of creating ResourceRecommendationBaseResource, please refer to the document of ResourceRecommendationBaseResource
string resourceUri = "resourceUri";
string recommendationId = "recommendationId";
ResourceIdentifier resourceRecommendationBaseResourceId = ResourceRecommendationBaseResource.CreateResourceIdentifier(resourceUri, recommendationId);
ResourceRecommendationBaseResource resourceRecommendationBase = client.GetResourceRecommendationBaseResource(resourceRecommendationBaseResourceId);

// get the collection of this SuppressionContractResource
SuppressionContractCollection collection = resourceRecommendationBase.GetSuppressionContracts();

// invoke the operation
string name = "suppressionName1";
bool result = await collection.ExistsAsync(name);

Console.WriteLine($"Succeeded: {result}");

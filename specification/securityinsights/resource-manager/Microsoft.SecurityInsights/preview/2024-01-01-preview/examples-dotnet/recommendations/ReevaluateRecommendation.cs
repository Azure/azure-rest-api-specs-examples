using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityInsights.Models;
using Azure.ResourceManager.SecurityInsights;

// Generated from example definition: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2024-01-01-preview/examples/recommendations/ReevaluateRecommendation.json
// this example is just showing the usage of "Reevaluate_Recommendation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityInsightsRecommendationResource created on azure
// for more information of creating SecurityInsightsRecommendationResource, please refer to the document of SecurityInsightsRecommendationResource
string subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
string resourceGroupName = "myRg";
string workspaceName = "myWorkspace";
Guid recommendationId = Guid.Parse("6d4b54eb-8684-4aa3-a156-3aa37b8014bc");
ResourceIdentifier securityInsightsRecommendationResourceId = SecurityInsightsRecommendationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, recommendationId);
SecurityInsightsRecommendationResource securityInsightsRecommendation = client.GetSecurityInsightsRecommendationResource(securityInsightsRecommendationResourceId);

// invoke the operation
ReevaluateResult result = await securityInsightsRecommendation.RecommendationReevaluateAsync();

Console.WriteLine($"Succeeded: {result}");

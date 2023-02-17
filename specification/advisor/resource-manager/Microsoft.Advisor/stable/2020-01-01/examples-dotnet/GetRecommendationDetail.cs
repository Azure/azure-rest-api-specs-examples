using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Advisor;

// Generated from example definition: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/GetRecommendationDetail.json
// this example is just showing the usage of "Recommendations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this ResourceRecommendationBaseResource
string resourceUri = "resourceUri";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", resourceUri));
ResourceRecommendationBaseCollection collection = client.GetResourceRecommendationBases(scopeId);

// invoke the operation
string recommendationId = "recommendationId";
bool result = await collection.ExistsAsync(recommendationId);

Console.WriteLine($"Succeeded: {result}");

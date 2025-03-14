using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Advisor;

// Generated from example definition: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/ListRecommendations.json
// this example is just showing the usage of "Recommendations_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this ResourceRecommendationBaseResource
string subscriptionId = "subscriptionId";
string scope = $"/subscriptions/{subscriptionId}";
ResourceRecommendationBaseCollection collection = client.GetResourceRecommendationBases(new ResourceIdentifier(scope));

// invoke the operation and iterate over the result
int? top = 10;
await foreach (ResourceRecommendationBaseResource item in collection.GetAllAsync(top: top))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ResourceRecommendationBaseData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

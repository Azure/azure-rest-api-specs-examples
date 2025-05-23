using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Chaos;

// Generated from example definition: specification/chaos/resource-manager/Microsoft.Chaos/stable/2024-01-01/examples/ListTargetTypes.json
// this example is just showing the usage of "TargetTypes_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "6b052e15-03d3-4f17-b2e1-be7f07588291";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this ChaosTargetTypeResource
string locationName = "westus2";
ChaosTargetTypeCollection collection = subscriptionResource.GetChaosTargetTypes(locationName);

// invoke the operation and iterate over the result
string continuationToken = null;
await foreach (ChaosTargetTypeResource item in collection.GetAllAsync(continuationToken: continuationToken))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ChaosTargetTypeData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

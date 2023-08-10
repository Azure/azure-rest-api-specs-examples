using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/GetLocationsWithExtendedLocations.json
// this example is just showing the usage of "Subscriptions_ListLocations" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "a1ffc958-d2c7-493e-9f1e-125a0477f536";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscription = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation and iterate over the result
bool? includeExtendedLocations = true;
await foreach (LocationExpanded item in subscription.GetLocationsAsync(includeExtendedLocations: includeExtendedLocations))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");

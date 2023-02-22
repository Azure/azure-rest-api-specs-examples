using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ResourceHealth;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2022-10-01-preview/examples/ImpactedResources_Get.json
// this example is just showing the usage of "ImpactedResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionEventResource created on azure
// for more information of creating SubscriptionEventResource, please refer to the document of SubscriptionEventResource
string subscriptionId = "subscriptionId";
string eventTrackingId = "BC_1-FXZ";
ResourceIdentifier subscriptionEventResourceId = SubscriptionEventResource.CreateResourceIdentifier(subscriptionId, eventTrackingId);
SubscriptionEventResource subscriptionEvent = client.GetSubscriptionEventResource(subscriptionEventResourceId);

// get the collection of this SubscriptionResourceHealthEventImpactedResource
SubscriptionResourceHealthEventImpactedResourceCollection collection = subscriptionEvent.GetSubscriptionResourceHealthEventImpactedResources();

// invoke the operation
string impactedResourceName = "abc-123-ghj-456";
bool result = await collection.ExistsAsync(impactedResourceName);

Console.WriteLine($"Succeeded: {result}");

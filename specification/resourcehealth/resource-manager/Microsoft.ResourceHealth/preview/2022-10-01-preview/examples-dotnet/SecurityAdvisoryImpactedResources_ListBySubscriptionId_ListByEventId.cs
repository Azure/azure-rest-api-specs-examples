using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ResourceHealth;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2022-10-01-preview/examples/SecurityAdvisoryImpactedResources_ListBySubscriptionId_ListByEventId.json
// this example is just showing the usage of "SecurityAdvisoryImpactedResources_ListBySubscriptionIdAndEventId" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
await foreach (EventImpactedResourceData item in subscriptionEvent.GetSecurityAdvisoryImpactedResourcesBySubscriptionIdAndEventIdAsync())
{
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {item.Id}");
}

Console.WriteLine($"Succeeded");

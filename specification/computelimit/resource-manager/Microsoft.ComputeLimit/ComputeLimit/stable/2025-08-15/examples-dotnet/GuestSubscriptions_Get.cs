using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ComputeLimit;

// Generated from example definition: 2025-08-15/GuestSubscriptions_Get.json
// this example is just showing the usage of "GuestSubscription_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "12345678-1234-1234-1234-123456789012";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this ComputeLimitGuestSubscriptionResource
AzureLocation location = new AzureLocation("eastus");
ComputeLimitGuestSubscriptionCollection collection = subscriptionResource.GetComputeLimitGuestSubscriptions(location);

// invoke the operation
string guestSubscriptionId = "11111111-1111-1111-1111-111111111111";
NullableResponse<ComputeLimitGuestSubscriptionResource> response = await collection.GetIfExistsAsync(guestSubscriptionId);
ComputeLimitGuestSubscriptionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ComputeLimitGuestSubscriptionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

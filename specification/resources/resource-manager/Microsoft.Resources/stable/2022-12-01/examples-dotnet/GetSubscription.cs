using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/GetSubscription.json
// this example is just showing the usage of "Subscriptions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenant = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this SubscriptionResource
SubscriptionCollection collection = tenant.GetSubscriptions();

// invoke the operation
string subscriptionId = "291bba3f-e0a5-47bc-a099-3bdcb2a50a05";
NullableResponse<SubscriptionResource> response = await collection.GetIfExistsAsync(subscriptionId);
SubscriptionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SubscriptionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

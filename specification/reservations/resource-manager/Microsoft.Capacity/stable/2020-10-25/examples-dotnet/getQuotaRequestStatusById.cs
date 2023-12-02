using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Reservations;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getQuotaRequestStatusById.json
// this example is just showing the usage of "QuotaRequestStatus_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this QuotaRequestDetailResource
string providerId = "Microsoft.Compute";
AzureLocation location = new AzureLocation("eastus");
QuotaRequestDetailCollection collection = subscriptionResource.GetQuotaRequestDetails(providerId, location);

// invoke the operation
Guid id = Guid.Parse("2B5C8515-37D8-4B6A-879B-CD641A2CF605");
NullableResponse<QuotaRequestDetailResource> response = await collection.GetIfExistsAsync(id);
QuotaRequestDetailResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    QuotaRequestDetailData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

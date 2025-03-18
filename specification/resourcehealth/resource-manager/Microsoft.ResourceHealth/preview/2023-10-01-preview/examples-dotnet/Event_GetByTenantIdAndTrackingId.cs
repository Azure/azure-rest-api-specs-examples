using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ResourceHealth;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/Event_GetByTenantIdAndTrackingId.json
// this example is just showing the usage of "Event_GetByTenantIdAndTrackingId" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this TenantResourceHealthEventResource
TenantResourceHealthEventCollection collection = tenantResource.GetTenantResourceHealthEvents();

// invoke the operation
string eventTrackingId = "eventTrackingId";
string filter = "properties/status eq 'Active'";
string queryStartTime = "7/10/2022";
NullableResponse<TenantResourceHealthEventResource> response = await collection.GetIfExistsAsync(eventTrackingId, filter: filter, queryStartTime: queryStartTime);
TenantResourceHealthEventResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ResourceHealthEventData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

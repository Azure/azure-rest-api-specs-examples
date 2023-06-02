using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ResourceHealth;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/Event_GetByTenantIdAndTrackingId.json
// this example is just showing the usage of "Event_GetByTenantIdAndTrackingId" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this TenantResourceHealthEventResource
TenantResourceHealthEventCollection collection = tenantResource.GetTenantResourceHealthEvents();

// invoke the operation
string eventTrackingId = "eventTrackingId";
string filter = "properties/status eq 'Active'";
string queryStartTime = "7/10/2022";
bool result = await collection.ExistsAsync(eventTrackingId, filter: filter, queryStartTime: queryStartTime);

Console.WriteLine($"Succeeded: {result}");

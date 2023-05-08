using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ResourceHealth;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/ImpactedResources_GetByTenantId.json
// this example is just showing the usage of "ImpactedResources_GetByTenantId" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResourceHealthEventResource created on azure
// for more information of creating TenantResourceHealthEventResource, please refer to the document of TenantResourceHealthEventResource
string eventTrackingId = "BC_1-FXZ";
ResourceIdentifier tenantResourceHealthEventResourceId = TenantResourceHealthEventResource.CreateResourceIdentifier(eventTrackingId);
TenantResourceHealthEventResource tenantResourceHealthEvent = client.GetTenantResourceHealthEventResource(tenantResourceHealthEventResourceId);

// get the collection of this TenantResourceHealthEventImpactedResource
TenantResourceHealthEventImpactedResourceCollection collection = tenantResourceHealthEvent.GetTenantResourceHealthEventImpactedResources();

// invoke the operation
string impactedResourceName = "abc-123-ghj-456";
bool result = await collection.ExistsAsync(impactedResourceName);

Console.WriteLine($"Succeeded: {result}");

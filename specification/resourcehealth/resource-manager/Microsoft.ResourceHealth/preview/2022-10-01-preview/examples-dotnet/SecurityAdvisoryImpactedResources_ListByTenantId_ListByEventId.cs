using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ResourceHealth;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2022-10-01-preview/examples/SecurityAdvisoryImpactedResources_ListByTenantId_ListByEventId.json
// this example is just showing the usage of "SecurityAdvisoryImpactedResources_ListByTenantIdAndEventId" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantEventResource created on azure
// for more information of creating TenantEventResource, please refer to the document of TenantEventResource
string eventTrackingId = "BC_1-FXZ";
ResourceIdentifier tenantEventResourceId = TenantEventResource.CreateResourceIdentifier(eventTrackingId);
TenantEventResource tenantEvent = client.GetTenantEventResource(tenantEventResourceId);

// invoke the operation and iterate over the result
await foreach (EventImpactedResourceData item in tenantEvent.GetSecurityAdvisoryImpactedResourcesByTenantIdAndEventIdAsync())
{
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {item.Id}");
}

Console.WriteLine($"Succeeded");

using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ResourceHealth;

// Generated from example definition: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/ImpactedResources_ListByTenantId_ListByEventId.json
// this example is just showing the usage of "ImpactedResources_ListByTenantIdAndEventId" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
string filter = "targetRegion eq 'westus'";
await foreach (TenantResourceHealthEventImpactedResource item in collection.GetAllAsync(filter: filter))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ResourceHealthEventImpactedResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");

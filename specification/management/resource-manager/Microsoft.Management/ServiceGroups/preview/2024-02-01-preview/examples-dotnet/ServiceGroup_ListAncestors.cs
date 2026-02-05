using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceGroups.Models;
using Azure.ResourceManager.ServiceGroups;

// Generated from example definition: specification/management/resource-manager/Microsoft.Management/ServiceGroups/preview/2024-02-01-preview/examples/ServiceGroup_ListAncestors.json
// this example is just showing the usage of "ServiceGroups_ListAncestors" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this ServiceGroupResource
ServiceGroupCollection collection = tenantResource.GetServiceGroups();

// invoke the operation and iterate over the result
string serviceGroupName = "20000000-0001-0000-0000-000000000000";
await foreach (ServiceGroupResource item in collection.GetAncestorsAsync(serviceGroupName))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ServiceGroupData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

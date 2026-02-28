using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagementGroups;

// Generated from example definition: specification/managementgroups/resource-manager/Microsoft.Management/ManagementGroups/stable/2023-04-01/examples/ListManagementGroups.json
// this example is just showing the usage of "ManagementGroups_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this ManagementGroupResource
ManagementGroupCollection collection = tenantResource.GetManagementGroups();

// invoke the operation and iterate over the result
string cacheControl = "no-cache";
await foreach (ManagementGroupResource item in collection.GetAllAsync(cacheControl: cacheControl))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ManagementGroupData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

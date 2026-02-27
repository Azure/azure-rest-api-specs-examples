using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagementGroups;

// Generated from example definition: specification/managementgroups/resource-manager/Microsoft.Management/ManagementGroups/stable/2023-04-01/examples/GetManagementGroupWithExpand.json
// this example is just showing the usage of "ManagementGroups_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this ManagementGroupResource
ManagementGroupCollection collection = tenantResource.GetManagementGroups();

// invoke the operation
string groupId = "20000000-0001-0000-0000-000000000000";
ManagementGroupExpandType? expand = ManagementGroupExpandType.Children;
string cacheControl = "no-cache";
NullableResponse<ManagementGroupResource> response = await collection.GetIfExistsAsync(groupId, expand: expand, cacheControl: cacheControl);
ManagementGroupResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ManagementGroupData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

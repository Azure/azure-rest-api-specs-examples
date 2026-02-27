using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagementGroups;

// Generated from example definition: specification/managementgroups/resource-manager/Microsoft.Management/ManagementGroups/stable/2023-04-01/examples/PutManagementGroup.json
// this example is just showing the usage of "ManagementGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this ManagementGroupResource
ManagementGroupCollection collection = tenantResource.GetManagementGroups();

// invoke the operation
string groupId = "ChildGroup";
ManagementGroupCreateOrUpdateContent content = new ManagementGroupCreateOrUpdateContent
{
    DisplayName = "ChildGroup",
    Details = new CreateManagementGroupDetails
    {
        Parent = new ManagementGroupParentCreateOptions
        {
            Id = "/providers/Microsoft.Management/managementGroups/RootGroup",
        },
    },
};
string cacheControl = "no-cache";
ArmOperation<ManagementGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, groupId, content, cacheControl: cacheControl);
ManagementGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagementGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

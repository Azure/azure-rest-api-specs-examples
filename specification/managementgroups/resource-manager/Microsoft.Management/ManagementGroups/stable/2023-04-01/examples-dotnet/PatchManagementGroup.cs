using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups.Models;
using Azure.ResourceManager.ManagementGroups;

// Generated from example definition: specification/managementgroups/resource-manager/Microsoft.Management/ManagementGroups/stable/2023-04-01/examples/PatchManagementGroup.json
// this example is just showing the usage of "ManagementGroups_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupResource created on azure
// for more information of creating ManagementGroupResource, please refer to the document of ManagementGroupResource
string groupId = "ChildGroup";
ResourceIdentifier managementGroupResourceId = ManagementGroupResource.CreateResourceIdentifier(groupId);
ManagementGroupResource managementGroup = client.GetManagementGroupResource(managementGroupResourceId);

// invoke the operation
ManagementGroupPatch patch = new ManagementGroupPatch
{
    DisplayName = "AlternateDisplayName",
    ParentGroupId = "/providers/Microsoft.Management/managementGroups/AlternateRootGroup",
};
string cacheControl = "no-cache";
ManagementGroupResource result = await managementGroup.UpdateAsync(patch, cacheControl: cacheControl);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagementGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

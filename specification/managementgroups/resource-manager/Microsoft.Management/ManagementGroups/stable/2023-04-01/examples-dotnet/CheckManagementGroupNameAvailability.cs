using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagementGroups;

// Generated from example definition: specification/managementgroups/resource-manager/Microsoft.Management/ManagementGroups/stable/2023-04-01/examples/CheckManagementGroupNameAvailability.json
// this example is just showing the usage of "ManagementGroups_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// get the collection of this ManagementGroupResource
ManagementGroupCollection collection = tenantResource.GetManagementGroups();

// invoke the operation
ManagementGroupNameAvailabilityContent content = new ManagementGroupNameAvailabilityContent
{
    Name = "nameTocheck",
    ResourceType = new ResourceType("Microsoft.Management/managementGroups"),
};
ManagementGroupNameAvailabilityResult result = await collection.CheckNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");

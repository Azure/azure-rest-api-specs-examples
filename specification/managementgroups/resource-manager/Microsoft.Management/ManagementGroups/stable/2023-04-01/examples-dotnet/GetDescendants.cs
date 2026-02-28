using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups.Models;
using Azure.ResourceManager.ManagementGroups;

// Generated from example definition: specification/managementgroups/resource-manager/Microsoft.Management/ManagementGroups/stable/2023-04-01/examples/GetDescendants.json
// this example is just showing the usage of "ManagementGroups_GetDescendants" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupResource created on azure
// for more information of creating ManagementGroupResource, please refer to the document of ManagementGroupResource
string groupId = "20000000-0000-0000-0000-000000000000";
ResourceIdentifier managementGroupResourceId = ManagementGroupResource.CreateResourceIdentifier(groupId);
ManagementGroupResource managementGroup = client.GetManagementGroupResource(managementGroupResourceId);

// invoke the operation and iterate over the result
await foreach (DescendantData item in managementGroup.GetDescendantsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");

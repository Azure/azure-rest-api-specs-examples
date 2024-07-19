using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleManagementPolicyAssignmentByScope.json
// this example is just showing the usage of "RoleManagementPolicyAssignments_ListForScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this RoleManagementPolicyAssignmentResource
string scope = "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
RoleManagementPolicyAssignmentCollection collection = client.GetRoleManagementPolicyAssignments(scopeId);

// invoke the operation and iterate over the result
await foreach (RoleManagementPolicyAssignmentResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RoleManagementPolicyAssignmentData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");

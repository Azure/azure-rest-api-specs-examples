using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/PutRoleManagementPolicyAssignment.json
// this example is just showing the usage of "RoleManagementPolicyAssignments_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this RoleManagementPolicyAssignmentResource
string scope = "providers/Microsoft.Subscription/subscriptions/129ff972-28f8-46b8-a726-e497be039368";
RoleManagementPolicyAssignmentCollection collection = client.GetRoleManagementPolicyAssignments(new ResourceIdentifier(scope));

// invoke the operation
string roleManagementPolicyAssignmentName = "b959d571-f0b5-4042-88a7-01be6cb22db9_a1705bd2-3a8f-45a5-8683-466fcfd5cc24";
RoleManagementPolicyAssignmentData data = new RoleManagementPolicyAssignmentData
{
    Scope = "/subscriptions/129ff972-28f8-46b8-a726-e497be039368",
    RoleDefinitionId = new ResourceIdentifier("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleDefinitions/a1705bd2-3a8f-45a5-8683-466fcfd5cc24"),
    PolicyId = new ResourceIdentifier("/subscriptions/129ff972-28f8-46b8-a726-e497be039368/providers/Microsoft.Authorization/roleManagementPolicies/b959d571-f0b5-4042-88a7-01be6cb22db9"),
};
ArmOperation<RoleManagementPolicyAssignmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, roleManagementPolicyAssignmentName, data);
RoleManagementPolicyAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RoleManagementPolicyAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Authorization;
using Azure.ResourceManager.Authorization.Models;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_CreateForResourceGroup.json
// this example is just showing the usage of "RoleAssignments_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this RoleAssignmentResource
string scope = "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/resourceGroups/testrg";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
RoleAssignmentCollection collection = client.GetRoleAssignments(scopeId);

// invoke the operation
string roleAssignmentName = "05c5a614-a7d6-4502-b150-c2fb455033ff";
RoleAssignmentCreateOrUpdateContent content = new RoleAssignmentCreateOrUpdateContent(new ResourceIdentifier("/subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2/providers/Microsoft.Authorization/roleDefinitions/0b5fe924-9a61-425c-96af-cfe6e287ca2d"), Guid.Parse("ce2ce14e-85d7-4629-bdbc-454d0519d987"))
{
    PrincipalType = RoleManagementPrincipalType.User,
};
ArmOperation<RoleAssignmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, roleAssignmentName, content);
RoleAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RoleAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

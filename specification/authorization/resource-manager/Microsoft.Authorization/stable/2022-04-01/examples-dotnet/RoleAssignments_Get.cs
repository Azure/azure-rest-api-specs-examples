using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Authorization.Models;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/RoleAssignments_Get.json
// this example is just showing the usage of "RoleAssignments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RoleAssignmentResource created on azure
// for more information of creating RoleAssignmentResource, please refer to the document of RoleAssignmentResource
string scope = "subscriptions/a925f2f7-5c63-4b7b-8799-25a5f97bc3b2";
string roleAssignmentName = "b0f43c54-e787-4862-89b1-a653fa9cf747";
ResourceIdentifier roleAssignmentResourceId = RoleAssignmentResource.CreateResourceIdentifier(scope, roleAssignmentName);
RoleAssignmentResource roleAssignment = client.GetRoleAssignmentResource(roleAssignmentResourceId);

// invoke the operation
RoleAssignmentResource result = await roleAssignment.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RoleAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

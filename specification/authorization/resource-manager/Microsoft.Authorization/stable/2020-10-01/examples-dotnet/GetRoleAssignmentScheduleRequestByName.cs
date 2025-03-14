using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Authorization.Models;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleAssignmentScheduleRequestByName.json
// this example is just showing the usage of "RoleAssignmentScheduleRequests_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this RoleAssignmentScheduleRequestResource
string scope = "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
RoleAssignmentScheduleRequestCollection collection = client.GetRoleAssignmentScheduleRequests(new ResourceIdentifier(scope));

// invoke the operation
string roleAssignmentScheduleRequestName = "fea7a502-9a96-4806-a26f-eee560e52045";
NullableResponse<RoleAssignmentScheduleRequestResource> response = await collection.GetIfExistsAsync(roleAssignmentScheduleRequestName);
RoleAssignmentScheduleRequestResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RoleAssignmentScheduleRequestData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

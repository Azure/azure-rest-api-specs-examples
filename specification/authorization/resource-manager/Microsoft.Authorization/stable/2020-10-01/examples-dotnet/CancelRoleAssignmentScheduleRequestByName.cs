using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Authorization.Models;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/CancelRoleAssignmentScheduleRequestByName.json
// this example is just showing the usage of "RoleAssignmentScheduleRequests_Cancel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RoleAssignmentScheduleRequestResource created on azure
// for more information of creating RoleAssignmentScheduleRequestResource, please refer to the document of RoleAssignmentScheduleRequestResource
string scope = "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
string roleAssignmentScheduleRequestName = "fea7a502-9a96-4806-a26f-eee560e52045";
ResourceIdentifier roleAssignmentScheduleRequestResourceId = RoleAssignmentScheduleRequestResource.CreateResourceIdentifier(scope, roleAssignmentScheduleRequestName);
RoleAssignmentScheduleRequestResource roleAssignmentScheduleRequest = client.GetRoleAssignmentScheduleRequestResource(roleAssignmentScheduleRequestResourceId);

// invoke the operation
await roleAssignmentScheduleRequest.CancelAsync();

Console.WriteLine("Succeeded");

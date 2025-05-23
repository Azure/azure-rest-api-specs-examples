using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Authorization.Models;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/ValidateRoleAssignmentScheduleRequestByName.json
// this example is just showing the usage of "RoleAssignmentScheduleRequests_Validate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RoleAssignmentScheduleRequestResource created on azure
// for more information of creating RoleAssignmentScheduleRequestResource, please refer to the document of RoleAssignmentScheduleRequestResource
string scope = "subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
string roleAssignmentScheduleRequestName = "fea7a502-9a96-4806-a26f-eee560e52045";
ResourceIdentifier roleAssignmentScheduleRequestResourceId = RoleAssignmentScheduleRequestResource.CreateResourceIdentifier(scope, roleAssignmentScheduleRequestName);
RoleAssignmentScheduleRequestResource roleAssignmentScheduleRequest = client.GetRoleAssignmentScheduleRequestResource(roleAssignmentScheduleRequestResourceId);

// invoke the operation
RoleAssignmentScheduleRequestData data = new RoleAssignmentScheduleRequestData
{
    RoleDefinitionId = new ResourceIdentifier("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
    PrincipalId = Guid.Parse("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
    RequestType = RoleManagementScheduleRequestType.SelfActivate,
    LinkedRoleEligibilityScheduleId = new ResourceIdentifier("b1477448-2cc6-4ceb-93b4-54a202a89413"),
    Condition = "@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'",
    ConditionVersion = "1.0",
    StartOn = DateTimeOffset.Parse("2020-09-09T21:35:27.91Z"),
    ExpirationType = RoleManagementScheduleExpirationType.AfterDuration,
    EndOn = default,
    Duration = XmlConvert.ToTimeSpan("PT8H"),
};
RoleAssignmentScheduleRequestResource result = await roleAssignmentScheduleRequest.ValidateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RoleAssignmentScheduleRequestData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

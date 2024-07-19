using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Authorization.Models;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/ValidateRoleEligibilityScheduleRequestByName.json
// this example is just showing the usage of "RoleEligibilityScheduleRequests_Validate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RoleEligibilityScheduleRequestResource created on azure
// for more information of creating RoleEligibilityScheduleRequestResource, please refer to the document of RoleEligibilityScheduleRequestResource
string scope = "subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
string roleEligibilityScheduleRequestName = "64caffb6-55c0-4deb-a585-68e948ea1ad6";
ResourceIdentifier roleEligibilityScheduleRequestResourceId = RoleEligibilityScheduleRequestResource.CreateResourceIdentifier(scope, roleEligibilityScheduleRequestName);
RoleEligibilityScheduleRequestResource roleEligibilityScheduleRequest = client.GetRoleEligibilityScheduleRequestResource(roleEligibilityScheduleRequestResourceId);

// invoke the operation
RoleEligibilityScheduleRequestData data = new RoleEligibilityScheduleRequestData()
{
    RoleDefinitionId = new ResourceIdentifier("/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608"),
    PrincipalId = Guid.Parse("a3bb8764-cb92-4276-9d2a-ca1e895e55ea"),
    RequestType = RoleManagementScheduleRequestType.AdminAssign,
    Condition = "@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase 'foo_storage_container'",
    ConditionVersion = "1.0",
    StartOn = DateTimeOffset.Parse("2020-09-09T21:31:27.91Z"),
    ExpirationType = RoleManagementScheduleExpirationType.AfterDuration,
    EndOn = null,
    Duration = XmlConvert.ToTimeSpan("P365D"),
};
RoleEligibilityScheduleRequestResource result = await roleEligibilityScheduleRequest.ValidateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RoleEligibilityScheduleRequestData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Authorization;
using Azure.ResourceManager.Authorization.Models;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleEligibilityScheduleRequestByName.json
// this example is just showing the usage of "RoleEligibilityScheduleRequests_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this RoleEligibilityScheduleRequestResource created on azure
// for more information of creating RoleEligibilityScheduleRequestResource, please refer to the document of RoleEligibilityScheduleRequestResource
string scope = "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
string roleEligibilityScheduleRequestName = "64caffb6-55c0-4deb-a585-68e948ea1ad6";
ResourceIdentifier roleEligibilityScheduleRequestResourceId = RoleEligibilityScheduleRequestResource.CreateResourceIdentifier(scope, roleEligibilityScheduleRequestName);
RoleEligibilityScheduleRequestResource roleEligibilityScheduleRequest = client.GetRoleEligibilityScheduleRequestResource(roleEligibilityScheduleRequestResourceId);

// invoke the operation
RoleEligibilityScheduleRequestResource result = await roleEligibilityScheduleRequest.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
RoleEligibilityScheduleRequestData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

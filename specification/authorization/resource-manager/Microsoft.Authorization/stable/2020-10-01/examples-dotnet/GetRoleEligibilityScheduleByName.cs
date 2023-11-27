using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleEligibilityScheduleByName.json
// this example is just showing the usage of "RoleEligibilitySchedules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this RoleEligibilityScheduleResource
string scope = "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
RoleEligibilityScheduleCollection collection = client.GetRoleEligibilitySchedules(scopeId);

// invoke the operation
string roleEligibilityScheduleName = "b1477448-2cc6-4ceb-93b4-54a202a89413";
NullableResponse<RoleEligibilityScheduleResource> response = await collection.GetIfExistsAsync(roleEligibilityScheduleName);
RoleEligibilityScheduleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RoleEligibilityScheduleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

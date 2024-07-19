using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Authorization.Models;
using Azure.ResourceManager.Authorization;

// Generated from example definition: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/GetRoleEligibilityScheduleRequestByScope.json
// this example is just showing the usage of "RoleEligibilityScheduleRequests_ListForScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this RoleEligibilityScheduleRequestResource
string scope = "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
RoleEligibilityScheduleRequestCollection collection = client.GetRoleEligibilityScheduleRequests(scopeId);

// invoke the operation and iterate over the result
string filter = "assignedTo('A3BB8764-CB92-4276-9D2A-CA1E895E55EA')";
await foreach (RoleEligibilityScheduleRequestResource item in collection.GetAllAsync(filter: filter))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    RoleEligibilityScheduleRequestData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");

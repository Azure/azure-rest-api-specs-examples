using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetwork;

// Generated from example definition: specification/managednetwork/resource-manager/Microsoft.ManagedNetwork/preview/2019-06-01-preview/examples/ScopeAssignment/ScopeAssignmentsList.json
// this example is just showing the usage of "ScopeAssignments_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this ScopeAssignmentResource
string scope = "subscriptions/subscriptionC";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
ScopeAssignmentCollection collection = client.GetScopeAssignments(scopeId);

// invoke the operation and iterate over the result
await foreach (ScopeAssignmentResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ScopeAssignmentData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");

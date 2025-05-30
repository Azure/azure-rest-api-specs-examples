using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/listPolicyAssignmentsForResource.json
// this example is just showing the usage of "PolicyAssignments_ListForResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this PolicyAssignmentResource
string subscriptionId = "ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
string resourceGroupName = "TestResourceGroup";
string resourceProviderNamespace = "Microsoft.Compute";
string parentResourcePath = "virtualMachines/MyTestVm";
string resourceType = "domainNames";
string resourceName = "MyTestComputer.cloudapp.net";
string scope = $"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}";
PolicyAssignmentCollection collection = client.GetGenericResource(new ResourceIdentifier(scope)).GetPolicyAssignments();

// invoke the operation and iterate over the result
await foreach (PolicyAssignmentResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PolicyAssignmentData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

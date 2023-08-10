using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/listPolicyAssignmentsForManagementGroup.json
// this example is just showing the usage of "PolicyAssignments_ListForManagementGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this PolicyAssignmentResource
string managementGroupId = "TestManagementGroup";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/providers/Microsoft.Management/managementGroups/{0}", managementGroupId));
PolicyAssignmentCollection collection = client.GetGenericResource(scopeId).GetPolicyAssignments();

// invoke the operation and iterate over the result
string filter = "atScope()";
await foreach (PolicyAssignmentResource item in collection.GetAllAsync(filter: filter))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PolicyAssignmentData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");

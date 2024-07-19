using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/updatePolicyAssignmentWithIdentity.json
// this example is just showing the usage of "PolicyAssignments_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PolicyAssignmentResource created on azure
// for more information of creating PolicyAssignmentResource, please refer to the document of PolicyAssignmentResource
string scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
string policyAssignmentName = "EnforceNaming";
ResourceIdentifier policyAssignmentResourceId = PolicyAssignmentResource.CreateResourceIdentifier(scope, policyAssignmentName);
PolicyAssignmentResource policyAssignment = client.GetPolicyAssignmentResource(policyAssignmentResourceId);

// invoke the operation
PolicyAssignmentPatch patch = new PolicyAssignmentPatch()
{
    Location = new AzureLocation("eastus"),
    Identity = new ManagedServiceIdentity("SystemAssigned"),
};
PolicyAssignmentResource result = await policyAssignment.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

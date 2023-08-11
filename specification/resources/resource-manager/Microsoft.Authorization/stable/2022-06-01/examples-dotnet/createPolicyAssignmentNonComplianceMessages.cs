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

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/createPolicyAssignmentNonComplianceMessages.json
// this example is just showing the usage of "PolicyAssignments_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this PolicyAssignmentResource
string scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
PolicyAssignmentCollection collection = client.GetGenericResource(scopeId).GetPolicyAssignments();

// invoke the operation
string policyAssignmentName = "securityInitAssignment";
PolicyAssignmentData data = new PolicyAssignmentData()
{
    DisplayName = "Enforce security policies",
    PolicyDefinitionId = "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/securityInitiative",
    NonComplianceMessages =
    {
    new NonComplianceMessage("Resources must comply with all internal security policies. See <internal site URL> for more info."),new NonComplianceMessage("Resource names must start with 'DeptA' and end with '-LC'.")
    {
    PolicyDefinitionReferenceId = "10420126870854049575",
    },new NonComplianceMessage("Storage accounts must have firewall rules configured.")
    {
    PolicyDefinitionReferenceId = "8572513655450389710",
    }
    },
};
ArmOperation<PolicyAssignmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, policyAssignmentName, data);
PolicyAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

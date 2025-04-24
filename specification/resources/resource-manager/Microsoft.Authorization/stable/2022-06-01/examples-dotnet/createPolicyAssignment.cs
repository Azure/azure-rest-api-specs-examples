using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Authorization/stable/2022-06-01/examples/createPolicyAssignment.json
// this example is just showing the usage of "PolicyAssignments_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this PolicyAssignmentResource
string scope = "subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2";
PolicyAssignmentCollection collection = client.GetGenericResource(new ResourceIdentifier(scope)).GetPolicyAssignments();

// invoke the operation
string policyAssignmentName = "EnforceNaming";
PolicyAssignmentData data = new PolicyAssignmentData
{
    DisplayName = "Enforce resource naming rules",
    PolicyDefinitionId = "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming",
    Parameters =
    {
    ["prefix"] = new ArmPolicyParameterValue
    {
    Value = BinaryData.FromObjectAsJson("DeptA"),
    },
    ["suffix"] = new ArmPolicyParameterValue
    {
    Value = BinaryData.FromObjectAsJson("-LC"),
    }
    },
    Description = "Force resource names to begin with given DeptA and end with -LC",
    Metadata = BinaryData.FromObjectAsJson(new
    {
        assignedBy = "Special Someone",
    }),
    NonComplianceMessages = { new NonComplianceMessage("Resource names must start with 'DeptA' and end with '-LC'.") },
};
ArmOperation<PolicyAssignmentResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, policyAssignmentName, data);
PolicyAssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicyAssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

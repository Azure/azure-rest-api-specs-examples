using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Blueprint;
using Azure.ResourceManager.Blueprint.Models;
using Azure.ResourceManager.Models;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/BlueprintAssignment_Create_UserAssignedMSI.json
// this example is just showing the usage of "Assignments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AssignmentResource created on azure
// for more information of creating AssignmentResource, please refer to the document of AssignmentResource
string resourceScope = "managementGroups/ContosoOnlineGroup";
string assignmentName = "assignSimpleBlueprint";
ResourceIdentifier assignmentResourceId = AssignmentResource.CreateResourceIdentifier(resourceScope, assignmentName);
AssignmentResource assignment = client.GetAssignmentResource(assignmentResourceId);

// invoke the operation
AssignmentData data = new AssignmentData(new Models.ManagedServiceIdentity(ManagedServiceIdentityType.UserAssigned)
{
    UserAssignedIdentities =
    {
    ["/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-resource-group/providers/Microsoft.ManagedIdentity/userAssignedIdentities/contoso-identity"] = new UserAssignedIdentity(),
    },
}, new Dictionary<string, ParameterValue>()
{
    ["costCenter"] = new ParameterValue()
    {
        Value = BinaryData.FromString("Contoso/Online/Shopping/Production"),
    },
    ["owners"] = new ParameterValue()
    {
        Value = BinaryData.FromObjectAsJson(new object[] { "johnDoe@contoso.com", "johnsteam@contoso.com" }),
    },
    ["storageAccountType"] = new ParameterValue()
    {
        Value = BinaryData.FromString("Standard_LRS"),
    },
}, new Dictionary<string, ResourceGroupValue>()
{
    ["storageRG"] = new ResourceGroupValue()
    {
        Name = "defaultRG",
        Location = new AzureLocation("eastus"),
    },
}, new AzureLocation("eastus"))
{
    Description = "enforce pre-defined simpleBlueprint to this XXXXXXXX subscription.",
    BlueprintId = "/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint",
    Scope = "subscriptions/00000000-0000-0000-0000-000000000000",
};
ArmOperation<AssignmentResource> lro = await assignment.UpdateAsync(WaitUntil.Completed, data);
AssignmentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AssignmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

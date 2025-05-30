using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Blueprint;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/BlueprintAssignmentOperation_List.json
// this example is just showing the usage of "AssignmentOperations_List" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this AssignmentOperationResource
AssignmentOperationCollection collection = assignment.GetAssignmentOperations();

// invoke the operation and iterate over the result
await foreach (AssignmentOperationResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AssignmentOperationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

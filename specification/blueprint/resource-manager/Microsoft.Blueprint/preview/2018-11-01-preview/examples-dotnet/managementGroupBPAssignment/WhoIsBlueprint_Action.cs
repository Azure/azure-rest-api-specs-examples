using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Blueprint;
using Azure.ResourceManager.Blueprint.Models;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/managementGroupBPAssignment/WhoIsBlueprint_Action.json
// this example is just showing the usage of "Assignments_WhoIsBlueprint" operation, for the dependent resources, they will have to be created separately.

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
WhoIsBlueprintContract result = await assignment.WhoIsBlueprintAsync();

Console.WriteLine($"Succeeded: {result}");

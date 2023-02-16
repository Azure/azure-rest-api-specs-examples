using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Blueprint;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/BlueprintAssignmentOperation_Get.json
// this example is just showing the usage of "AssignmentOperations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AssignmentResource created on azure
// for more information of creating AssignmentResource, please refer to the document of AssignmentResource
string resourceScope = "subscriptions/00000000-0000-0000-0000-000000000000";
string assignmentName = "assignSimpleBlueprint";
ResourceIdentifier assignmentResourceId = AssignmentResource.CreateResourceIdentifier(resourceScope, assignmentName);
AssignmentResource assignment = client.GetAssignmentResource(assignmentResourceId);

// get the collection of this AssignmentOperationResource
AssignmentOperationCollection collection = assignment.GetAssignmentOperations();

// invoke the operation
string assignmentOperationName = "fb5d4dcb-7ce2-4087-ba7a-459aa74e5e0f";
bool result = await collection.ExistsAsync(assignmentOperationName);

Console.WriteLine($"Succeeded: {result}");

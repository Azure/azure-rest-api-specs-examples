using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Blueprint;

// Generated from example definition: specification/blueprint/resource-manager/Microsoft.Blueprint/preview/2018-11-01-preview/examples/subscriptionBPAssignment/BlueprintAssignmentOperation_Get.json
// this example is just showing the usage of "AssignmentOperations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AssignmentOperationResource created on azure
// for more information of creating AssignmentOperationResource, please refer to the document of AssignmentOperationResource
string resourceScope = "subscriptions/00000000-0000-0000-0000-000000000000";
string assignmentName = "assignSimpleBlueprint";
string assignmentOperationName = "fb5d4dcb-7ce2-4087-ba7a-459aa74e5e0f";
ResourceIdentifier assignmentOperationResourceId = AssignmentOperationResource.CreateResourceIdentifier(resourceScope, assignmentName, assignmentOperationName);
AssignmentOperationResource assignmentOperation = client.GetAssignmentOperationResource(assignmentOperationResourceId);

// invoke the operation
AssignmentOperationResource result = await assignmentOperation.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AssignmentOperationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

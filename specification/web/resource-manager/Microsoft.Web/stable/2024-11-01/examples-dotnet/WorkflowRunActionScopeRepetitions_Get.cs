using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/WorkflowRunActionScopeRepetitions_Get.json
// this example is just showing the usage of "WorkflowRunActionScopeRepetitions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkflowRunActionScopeRepetitionResource created on azure
// for more information of creating WorkflowRunActionScopeRepetitionResource, please refer to the document of WorkflowRunActionScopeRepetitionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string name = "test-name";
string workflowName = "testFlow";
string runName = "08586776228332053161046300351";
string actionName = "for_each";
string repetitionName = "000000";
ResourceIdentifier workflowRunActionScopeRepetitionResourceId = WorkflowRunActionScopeRepetitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, workflowName, runName, actionName, repetitionName);
WorkflowRunActionScopeRepetitionResource workflowRunActionScopeRepetition = client.GetWorkflowRunActionScopeRepetitionResource(workflowRunActionScopeRepetitionResourceId);

// invoke the operation
WorkflowRunActionScopeRepetitionResource result = await workflowRunActionScopeRepetition.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WorkflowRunActionRepetitionDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

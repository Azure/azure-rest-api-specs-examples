using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActionScopeRepetitions_Get.json
// this example is just showing the usage of "WorkflowRunActionScopeRepetitions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowRunActionResource created on azure
// for more information of creating LogicWorkflowRunActionResource, please refer to the document of LogicWorkflowRunActionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string workflowName = "testFlow";
string runName = "08586776228332053161046300351";
string actionName = "for_each";
ResourceIdentifier logicWorkflowRunActionResourceId = LogicWorkflowRunActionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName, runName, actionName);
LogicWorkflowRunActionResource logicWorkflowRunAction = client.GetLogicWorkflowRunActionResource(logicWorkflowRunActionResourceId);

// get the collection of this LogicWorkflowRunActionScopeRepetitionResource
LogicWorkflowRunActionScopeRepetitionCollection collection = logicWorkflowRunAction.GetLogicWorkflowRunActionScopeRepetitions();

// invoke the operation
string repetitionName = "000000";
bool result = await collection.ExistsAsync(repetitionName);

Console.WriteLine($"Succeeded: {result}");

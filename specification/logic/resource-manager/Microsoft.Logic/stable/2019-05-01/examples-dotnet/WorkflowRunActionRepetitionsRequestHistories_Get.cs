using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActionRepetitionsRequestHistories_Get.json
// this example is just showing the usage of "WorkflowRunActionRepetitionsRequestHistories_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowRunActionRepetitionResource created on azure
// for more information of creating LogicWorkflowRunActionRepetitionResource, please refer to the document of LogicWorkflowRunActionRepetitionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "test-resource-group";
string workflowName = "test-workflow";
string runName = "08586776228332053161046300351";
string actionName = "HTTP_Webhook";
string repetitionName = "000001";
ResourceIdentifier logicWorkflowRunActionRepetitionResourceId = LogicWorkflowRunActionRepetitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName, runName, actionName, repetitionName);
LogicWorkflowRunActionRepetitionResource logicWorkflowRunActionRepetition = client.GetLogicWorkflowRunActionRepetitionResource(logicWorkflowRunActionRepetitionResourceId);

// get the collection of this LogicWorkflowRunActionRepetitionRequestHistoryResource
LogicWorkflowRunActionRepetitionRequestHistoryCollection collection = logicWorkflowRunActionRepetition.GetLogicWorkflowRunActionRepetitionRequestHistories();

// invoke the operation
string requestHistoryName = "08586611142732800686";
bool result = await collection.ExistsAsync(requestHistoryName);

Console.WriteLine($"Succeeded: {result}");

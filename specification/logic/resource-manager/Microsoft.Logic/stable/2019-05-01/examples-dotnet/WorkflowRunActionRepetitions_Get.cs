using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;
using Azure.ResourceManager.Logic.Models;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActionRepetitions_Get.json
// this example is just showing the usage of "WorkflowRunActionRepetitions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowRunActionRepetitionResource created on azure
// for more information of creating LogicWorkflowRunActionRepetitionResource, please refer to the document of LogicWorkflowRunActionRepetitionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string workflowName = "testFlow";
string runName = "08586776228332053161046300351";
string actionName = "testAction";
string repetitionName = "000001";
ResourceIdentifier logicWorkflowRunActionRepetitionResourceId = LogicWorkflowRunActionRepetitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName, runName, actionName, repetitionName);
LogicWorkflowRunActionRepetitionResource logicWorkflowRunActionRepetition = client.GetLogicWorkflowRunActionRepetitionResource(logicWorkflowRunActionRepetitionResourceId);

// invoke the operation
LogicWorkflowRunActionRepetitionResource result = await logicWorkflowRunActionRepetition.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LogicWorkflowRunActionRepetitionDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Logic.Models;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunActions_Get.json
// this example is just showing the usage of "WorkflowRunActions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowRunActionResource created on azure
// for more information of creating LogicWorkflowRunActionResource, please refer to the document of LogicWorkflowRunActionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "test-resource-group";
string workflowName = "test-workflow";
string runName = "08586676746934337772206998657CU22";
string actionName = "HTTP";
ResourceIdentifier logicWorkflowRunActionResourceId = LogicWorkflowRunActionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName, runName, actionName);
LogicWorkflowRunActionResource logicWorkflowRunAction = client.GetLogicWorkflowRunActionResource(logicWorkflowRunActionResourceId);

// invoke the operation
LogicWorkflowRunActionResource result = await logicWorkflowRunAction.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LogicWorkflowRunActionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

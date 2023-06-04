using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRuns_Get.json
// this example is just showing the usage of "WorkflowRuns_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowResource created on azure
// for more information of creating LogicWorkflowResource, please refer to the document of LogicWorkflowResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "test-resource-group";
string workflowName = "test-workflow";
ResourceIdentifier logicWorkflowResourceId = LogicWorkflowResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName);
LogicWorkflowResource logicWorkflow = client.GetLogicWorkflowResource(logicWorkflowResourceId);

// get the collection of this LogicWorkflowRunResource
LogicWorkflowRunCollection collection = logicWorkflow.GetLogicWorkflowRuns();

// invoke the operation
string runName = "08586676746934337772206998657CU22";
bool result = await collection.ExistsAsync(runName);

Console.WriteLine($"Succeeded: {result}");

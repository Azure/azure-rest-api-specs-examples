using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRunOperations_Get.json
// this example is just showing the usage of "WorkflowRunOperations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowRunOperationResource created on azure
// for more information of creating LogicWorkflowRunOperationResource, please refer to the document of LogicWorkflowRunOperationResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string workflowName = "testFlow";
string runName = "08586774142730039209110422528";
string operationId = "ebdcbbde-c4db-43ec-987c-fd0f7726f43b";
ResourceIdentifier logicWorkflowRunOperationResourceId = LogicWorkflowRunOperationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName, runName, operationId);
LogicWorkflowRunOperationResource logicWorkflowRunOperation = client.GetLogicWorkflowRunOperationResource(logicWorkflowRunOperationResourceId);

// invoke the operation
LogicWorkflowRunOperationResource result = await logicWorkflowRunOperation.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LogicWorkflowRunData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;
using Azure.ResourceManager.Logic.Models;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggers_GetSchemaJson.json
// this example is just showing the usage of "WorkflowTriggers_GetSchemaJson" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowTriggerResource created on azure
// for more information of creating LogicWorkflowTriggerResource, please refer to the document of LogicWorkflowTriggerResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string workflowName = "testWorkflow";
string triggerName = "testTrigger";
ResourceIdentifier logicWorkflowTriggerResourceId = LogicWorkflowTriggerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName, triggerName);
LogicWorkflowTriggerResource logicWorkflowTrigger = client.GetLogicWorkflowTriggerResource(logicWorkflowTriggerResourceId);

// invoke the operation
LogicJsonSchema result = await logicWorkflowTrigger.GetSchemaJsonAsync();

Console.WriteLine($"Succeeded: {result}");

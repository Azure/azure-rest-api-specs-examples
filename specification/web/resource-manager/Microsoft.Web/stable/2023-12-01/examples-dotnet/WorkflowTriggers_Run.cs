using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/WorkflowTriggers_Run.json
// this example is just showing the usage of "WorkflowTriggers_Run" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkflowTriggerResource created on azure
// for more information of creating WorkflowTriggerResource, please refer to the document of WorkflowTriggerResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "test-resource-group";
string name = "test-name";
string workflowName = "test-workflow";
string triggerName = "recurrence";
ResourceIdentifier workflowTriggerResourceId = WorkflowTriggerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, workflowName, triggerName);
WorkflowTriggerResource workflowTrigger = client.GetWorkflowTriggerResource(workflowTriggerResourceId);

// invoke the operation
await workflowTrigger.RunAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

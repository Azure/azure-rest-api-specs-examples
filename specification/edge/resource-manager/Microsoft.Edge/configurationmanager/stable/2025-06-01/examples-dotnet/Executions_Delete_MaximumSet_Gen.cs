using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/Executions_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "Execution_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeExecutionResource created on azure
// for more information of creating EdgeExecutionResource, please refer to the document of EdgeExecutionResource
string subscriptionId = "EE6D9590-0D52-4B1C-935C-FE49DBF838EB";
string resourceGroupName = "rgconfigurationmanager";
string contextName = "abcde";
string workflowName = "abcde";
string versionName = "abcde";
string executionName = "abcde";
ResourceIdentifier edgeExecutionResourceId = EdgeExecutionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, contextName, workflowName, versionName, executionName);
EdgeExecutionResource edgeExecution = client.GetEdgeExecutionResource(edgeExecutionResourceId);

// invoke the operation
await edgeExecution.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

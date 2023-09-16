using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DefenderEasm;
using Azure.ResourceManager.DefenderEasm.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/riskiq/resource-manager/Microsoft.Easm/preview/2023-04-01-preview/examples/Tasks_GetByWorkspace.json
// this example is just showing the usage of "Tasks_GetByWorkspace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EasmWorkspaceResource created on azure
// for more information of creating EasmWorkspaceResource, please refer to the document of EasmWorkspaceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "dummyrg";
string workspaceName = "ThisisaWorkspace";
ResourceIdentifier easmWorkspaceResourceId = EasmWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
EasmWorkspaceResource easmWorkspace = client.GetEasmWorkspaceResource(easmWorkspaceResourceId);

// invoke the operation
string taskId = "ThisisaTaskId";
EasmTask result = await easmWorkspace.GetTaskByWorkspaceAsync(taskId);

Console.WriteLine($"Succeeded: {result}");

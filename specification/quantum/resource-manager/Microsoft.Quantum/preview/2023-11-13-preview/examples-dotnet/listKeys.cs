using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Quantum.Models;
using Azure.ResourceManager.Quantum;

// Generated from example definition: specification/quantum/resource-manager/Microsoft.Quantum/preview/2023-11-13-preview/examples/listKeys.json
// this example is just showing the usage of "Workspace_ListKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this QuantumWorkspaceResource created on azure
// for more information of creating QuantumWorkspaceResource, please refer to the document of QuantumWorkspaceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "quantumResourcegroup";
string workspaceName = "quantumworkspace1";
ResourceIdentifier quantumWorkspaceResourceId = QuantumWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
QuantumWorkspaceResource quantumWorkspace = client.GetQuantumWorkspaceResource(quantumWorkspaceResourceId);

// invoke the operation
WorkspaceKeyListResult result = await quantumWorkspace.GetKeysWorkspaceAsync();

Console.WriteLine($"Succeeded: {result}");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DefenderEasm.Models;
using Azure.ResourceManager.DefenderEasm;

// Generated from example definition: specification/riskiq/resource-manager/Microsoft.Easm/preview/2023-04-01-preview/examples/Workspaces_Get.json
// this example is just showing the usage of "Workspaces_Get" operation, for the dependent resources, they will have to be created separately.

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
EasmWorkspaceResource result = await easmWorkspace.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EasmWorkspaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

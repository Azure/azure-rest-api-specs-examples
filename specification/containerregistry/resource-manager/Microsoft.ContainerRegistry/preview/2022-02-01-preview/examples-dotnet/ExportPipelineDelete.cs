using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerRegistry;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.Models;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ExportPipelineDelete.json
// this example is just showing the usage of "ExportPipelines_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExportPipelineResource created on azure
// for more information of creating ExportPipelineResource, please refer to the document of ExportPipelineResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string exportPipelineName = "myExportPipeline";
ResourceIdentifier exportPipelineResourceId = ExportPipelineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, exportPipelineName);
ExportPipelineResource exportPipeline = client.GetExportPipelineResource(exportPipelineResourceId);

// invoke the operation
await exportPipeline.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

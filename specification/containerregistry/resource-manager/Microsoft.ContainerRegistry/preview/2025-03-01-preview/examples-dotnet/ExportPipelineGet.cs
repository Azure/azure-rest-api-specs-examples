using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/ExportPipelineGet.json
// this example is just showing the usage of "ExportPipelines_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryExportPipelineResource created on azure
// for more information of creating ContainerRegistryExportPipelineResource, please refer to the document of ContainerRegistryExportPipelineResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string exportPipelineName = "myExportPipeline";
ResourceIdentifier containerRegistryExportPipelineResourceId = ContainerRegistryExportPipelineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, exportPipelineName);
ContainerRegistryExportPipelineResource containerRegistryExportPipeline = client.GetContainerRegistryExportPipelineResource(containerRegistryExportPipelineResourceId);

// invoke the operation
ContainerRegistryExportPipelineResource result = await containerRegistryExportPipeline.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryExportPipelineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

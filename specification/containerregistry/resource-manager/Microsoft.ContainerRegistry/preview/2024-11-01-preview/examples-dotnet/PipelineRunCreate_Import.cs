using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2024-11-01-preview/examples/PipelineRunCreate_Import.json
// this example is just showing the usage of "PipelineRuns_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryResource created on azure
// for more information of creating ContainerRegistryResource, please refer to the document of ContainerRegistryResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
ResourceIdentifier containerRegistryResourceId = ContainerRegistryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName);
ContainerRegistryResource containerRegistry = client.GetContainerRegistryResource(containerRegistryResourceId);

// get the collection of this ContainerRegistryPipelineRunResource
ContainerRegistryPipelineRunCollection collection = containerRegistry.GetContainerRegistryPipelineRuns();

// invoke the operation
string pipelineRunName = "myPipelineRun";
ContainerRegistryPipelineRunData data = new ContainerRegistryPipelineRunData
{
    Request = new ConnectedRegistryPipelineRunContent
    {
        PipelineResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerRegistry/registries/myRegistry/importPipelines/myImportPipeline"),
        Source = new ContainerRegistryPipelineRunSourceProperties
        {
            SourceType = ContainerRegistryPipelineRunSourceType.AzureStorageBlob,
            Name = "myblob.tar.gz",
        },
        CatalogDigest = "sha256@",
    },
    ForceUpdateTag = "2020-03-04T17:23:21.9261521+00:00",
};
ArmOperation<ContainerRegistryPipelineRunResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, pipelineRunName, data);
ContainerRegistryPipelineRunResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryPipelineRunData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/ImportPipelineCreate.json
// this example is just showing the usage of "ImportPipelines_Create" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ContainerRegistryImportPipelineResource
ContainerRegistryImportPipelineCollection collection = containerRegistry.GetContainerRegistryImportPipelines();

// invoke the operation
string importPipelineName = "myImportPipeline";
ContainerRegistryImportPipelineData data = new ContainerRegistryImportPipelineData
{
    Location = new AzureLocation("westus"),
    Identity = new ManagedServiceIdentity("UserAssigned")
    {
        UserAssignedIdentities =
        {
        [new ResourceIdentifier("/subscriptions/f9d7ebed-adbd-4cb4-b973-aaf82c136138/resourcegroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2")] = new UserAssignedIdentity()
        },
    },
    Source = new ImportPipelineSourceProperties(new Uri("https://myvault.vault.azure.net/secrets/acrimportsas"))
    {
        ContainerRegistryPipelineSourceType = ContainerRegistryPipelineSourceType.AzureStorageBlobContainer,
        Uri = new Uri("https://accountname.blob.core.windows.net/containername"),
    },
    Options = { ContainerRegistryPipelineOption.OverwriteTags, ContainerRegistryPipelineOption.DeleteSourceBlobOnSuccess, ContainerRegistryPipelineOption.ContinueOnErrors },
};
ArmOperation<ContainerRegistryImportPipelineResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, importPipelineName, data);
ContainerRegistryImportPipelineResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryImportPipelineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

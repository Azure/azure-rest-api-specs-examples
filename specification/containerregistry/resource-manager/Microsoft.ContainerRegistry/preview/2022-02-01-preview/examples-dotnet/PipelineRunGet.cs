using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerRegistry;
using Azure.ResourceManager.ContainerRegistry.Models;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/PipelineRunGet.json
// this example is just showing the usage of "PipelineRuns_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryPipelineRunResource created on azure
// for more information of creating ContainerRegistryPipelineRunResource, please refer to the document of ContainerRegistryPipelineRunResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string pipelineRunName = "myPipelineRun";
ResourceIdentifier containerRegistryPipelineRunResourceId = ContainerRegistryPipelineRunResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, pipelineRunName);
ContainerRegistryPipelineRunResource containerRegistryPipelineRun = client.GetContainerRegistryPipelineRunResource(containerRegistryPipelineRunResourceId);

// invoke the operation
ContainerRegistryPipelineRunResource result = await containerRegistryPipelineRun.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryPipelineRunData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

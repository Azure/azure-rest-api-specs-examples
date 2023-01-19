using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerRegistry;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.Models;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/ImportPipelineGet.json
// this example is just showing the usage of "ImportPipelines_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ImportPipelineResource created on azure
// for more information of creating ImportPipelineResource, please refer to the document of ImportPipelineResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string importPipelineName = "myImportPipeline";
ResourceIdentifier importPipelineResourceId = ImportPipelineResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, importPipelineName);
ImportPipelineResource importPipeline = client.GetImportPipelineResource(importPipelineResourceId);

// invoke the operation
ImportPipelineResource result = await importPipeline.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ImportPipelineData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

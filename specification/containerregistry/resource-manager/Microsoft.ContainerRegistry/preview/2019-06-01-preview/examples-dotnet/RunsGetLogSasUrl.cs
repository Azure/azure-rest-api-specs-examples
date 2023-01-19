using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ContainerRegistry;
using Azure.ResourceManager.ContainerRegistry.Models;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RunsGetLogSasUrl.json
// this example is just showing the usage of "Runs_GetLogSasUrl" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryRunResource created on azure
// for more information of creating ContainerRegistryRunResource, please refer to the document of ContainerRegistryRunResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string runId = "0accec26-d6de-4757-8e74-d080f38eaaab";
ResourceIdentifier containerRegistryRunResourceId = ContainerRegistryRunResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, runId);
ContainerRegistryRunResource containerRegistryRun = client.GetContainerRegistryRunResource(containerRegistryRunResourceId);

// invoke the operation
ContainerRegistryRunGetLogResult result = await containerRegistryRun.GetLogSasUrlAsync();

Console.WriteLine($"Succeeded: {result}");

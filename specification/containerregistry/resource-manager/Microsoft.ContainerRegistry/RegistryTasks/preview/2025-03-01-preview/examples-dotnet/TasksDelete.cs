using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/TasksDelete.json
// this example is just showing the usage of "Tasks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryTaskResource created on azure
// for more information of creating ContainerRegistryTaskResource, please refer to the document of ContainerRegistryTaskResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
string taskName = "myTask";
ResourceIdentifier containerRegistryTaskResourceId = ContainerRegistryTaskResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName, taskName);
ContainerRegistryTaskResource containerRegistryTask = client.GetContainerRegistryTaskResource(containerRegistryTaskResourceId);

// invoke the operation
await containerRegistryTask.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

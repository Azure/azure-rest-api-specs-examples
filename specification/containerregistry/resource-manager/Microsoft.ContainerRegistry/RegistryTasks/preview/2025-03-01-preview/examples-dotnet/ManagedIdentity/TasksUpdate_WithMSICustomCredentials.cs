using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/ManagedIdentity/TasksUpdate_WithMSICustomCredentials.json
// this example is just showing the usage of "Tasks_Update" operation, for the dependent resources, they will have to be created separately.

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
ContainerRegistryTaskPatch patch = new ContainerRegistryTaskPatch
{
    Tags =
    {
    ["testkey"] = "value"
    },
    Status = ContainerRegistryTaskStatus.Enabled,
    AgentCpu = 3,
    Step = new ContainerRegistryDockerBuildStepUpdateContent
    {
        ImageNames = { "azurerest:testtag1" },
        DockerFilePath = "src/DockerFile",
    },
    Trigger = new ContainerRegistryTriggerUpdateContent
    {
        SourceTriggers = {new ContainerRegistrySourceTriggerUpdateContent("mySourceTrigger")
        {
        SourceRepository = new SourceCodeRepoUpdateContent
        {
        SourceControlAuthProperties = new SourceCodeRepoAuthInfoUpdateContent
        {
        TokenType = SourceCodeRepoAuthTokenType.Pat,
        Token = "xxxxx",
        },
        },
        SourceTriggerEvents = {ContainerRegistrySourceTriggerEvent.Commit},
        }},
    },
    Credentials = new ContainerRegistryCredentials
    {
        CustomRegistries =
        {
        ["myregistry.azurecr.io"] = new CustomRegistryCredentials
        {
        Identity = "[system]",
        }
        },
    },
    LogTemplate = null,
};
ContainerRegistryTaskResource result = await containerRegistryTask.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryTaskData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

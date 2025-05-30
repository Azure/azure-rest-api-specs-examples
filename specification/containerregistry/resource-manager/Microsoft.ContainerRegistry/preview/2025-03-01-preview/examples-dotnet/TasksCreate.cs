using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2025-03-01-preview/examples/TasksCreate.json
// this example is just showing the usage of "Tasks_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerRegistryResource created on azure
// for more information of creating ContainerRegistryResource, please refer to the document of ContainerRegistryResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "myResourceGroup";
string registryName = "myRegistry";
ResourceIdentifier containerRegistryResourceId = ContainerRegistryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, registryName);
ContainerRegistryResource containerRegistry = client.GetContainerRegistryResource(containerRegistryResourceId);

// get the collection of this ContainerRegistryTaskResource
ContainerRegistryTaskCollection collection = containerRegistry.GetContainerRegistryTasks();

// invoke the operation
string taskName = "mytTask";
ContainerRegistryTaskData data = new ContainerRegistryTaskData(new AzureLocation("eastus"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Status = ContainerRegistryTaskStatus.Enabled,
    Platform = new ContainerRegistryPlatformProperties(ContainerRegistryOS.Linux)
    {
        Architecture = ContainerRegistryOSArchitecture.Amd64,
    },
    AgentCpu = 2,
    Step = new ContainerRegistryDockerBuildStep("src/DockerFile")
    {
        ImageNames = { "azurerest:testtag" },
        IsPushEnabled = true,
        NoCache = false,
        Arguments = {new ContainerRegistryRunArgument("mytestargument", "mytestvalue")
        {
        IsSecret = false,
        }, new ContainerRegistryRunArgument("mysecrettestargument", "mysecrettestvalue")
        {
        IsSecret = true,
        }},
        ContextPath = "src",
    },
    Trigger = new ContainerRegistryTriggerProperties
    {
        TimerTriggers = { new ContainerRegistryTimerTrigger("30 9 * * 1-5", "myTimerTrigger") },
        SourceTriggers = {new ContainerRegistrySourceTrigger(new SourceCodeRepoProperties(SourceControlType.Github, new Uri("https://github.com/Azure/azure-rest-api-specs"))
        {
        Branch = "master",
        SourceControlAuthProperties = new SourceCodeRepoAuthInfo(SourceCodeRepoAuthTokenType.Pat, "xxxxx"),
        }, new ContainerRegistrySourceTriggerEvent[]{ContainerRegistrySourceTriggerEvent.Commit}, "mySourceTrigger")},
        BaseImageTrigger = new ContainerRegistryBaseImageTrigger(ContainerRegistryBaseImageTriggerType.Runtime, "myBaseImageTrigger")
        {
            UpdateTriggerEndpoint = "https://user:pass@mycicd.webhook.com?token=foo",
            UpdateTriggerPayloadType = ContainerRegistryUpdateTriggerPayloadType.Token,
        },
    },
    LogTemplate = "acr/tasks:{{.Run.OS}}",
    IsSystemTask = false,
    Tags =
    {
    ["testkey"] = "value"
    },
};
ArmOperation<ContainerRegistryTaskResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, taskName, data);
ContainerRegistryTaskResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryTaskData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

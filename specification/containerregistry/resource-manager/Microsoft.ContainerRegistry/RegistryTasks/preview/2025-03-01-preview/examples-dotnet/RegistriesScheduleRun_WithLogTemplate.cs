using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/RegistriesScheduleRun_WithLogTemplate.json
// this example is just showing the usage of "Schedules_ScheduleRun" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
ContainerRegistryRunContent content = new ContainerRegistryDockerBuildContent("DockerFile", new ContainerRegistryPlatformProperties(ContainerRegistryOS.Linux)
{
    Architecture = ContainerRegistryOSArchitecture.Amd64,
})
{
    ImageNames = { "azurerest:testtag" },
    IsPushEnabled = true,
    NoCache = true,
    Arguments = {new ContainerRegistryRunArgument("mytestargument", "mytestvalue")
    {
    IsSecret = false,
    }, new ContainerRegistryRunArgument("mysecrettestargument", "mysecrettestvalue")
    {
    IsSecret = true,
    }},
    AgentCpu = 2,
    SourceLocation = "https://myaccount.blob.core.windows.net/sascontainer/source.zip?sv=2015-04-05&st=2015-04-29T22%3A18%3A26Z&se=2015-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=Z%2FRHIX5Xcg0Mq2rqI3OlWTjEg2tYkboXr1P9ZUXDtkk%3D",
    IsArchiveEnabled = true,
    LogTemplate = "acr/tasks:{{.Run.OS}}",
};
ContainerRegistryRunResource result = await containerRegistry.ScheduleRunAsync(content);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryRunData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

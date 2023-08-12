using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/Job_Start.json
// this example is just showing the usage of "Jobs_Start" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppJobResource created on azure
// for more information of creating ContainerAppJobResource, please refer to the document of ContainerAppJobResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string jobName = "testcontainerAppsJob0";
ResourceIdentifier containerAppJobResourceId = ContainerAppJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName);
ContainerAppJobResource containerAppJob = client.GetContainerAppJobResource(containerAppJobResourceId);

// invoke the operation
ContainerAppJobExecutionTemplate template = new ContainerAppJobExecutionTemplate()
{
    Containers =
    {
    new JobExecutionContainer()
    {
    Image = "repo/testcontainerAppsJob0:v4",
    Name = "testcontainerAppsJob0",
    Resources = new AppContainerResources()
    {
    Cpu = 0.2,
    Memory = "100Mi",
    },
    }
    },
    InitContainers =
    {
    new JobExecutionContainer()
    {
    Image = "repo/testcontainerAppsJob0:v4",
    Name = "testinitcontainerAppsJob0",
    Command =
    {
    "/bin/sh"
    },
    Args =
    {
    "-c","while true; do echo hello; sleep 10;done"
    },
    Resources = new AppContainerResources()
    {
    Cpu = 0.2,
    Memory = "100Mi",
    },
    }
    },
};
ArmOperation<ContainerAppJobExecutionBase> lro = await containerAppJob.StartAsync(WaitUntil.Completed, template: template);
ContainerAppJobExecutionBase result = lro.Value;

Console.WriteLine($"Succeeded: {result}");

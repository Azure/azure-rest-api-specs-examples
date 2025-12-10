using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2025-03-01-preview/examples/RegistriesScheduleRun_Task.json
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
ContainerRegistryRunContent content = new ContainerRegistryTaskRunContent(new ResourceIdentifier("myTask"))
{
    OverrideTaskStepProperties = new ContainerRegistryOverrideTaskStepProperties
    {
        File = "overriddenDockerfile",
        Arguments = {new ContainerRegistryRunArgument("mytestargument", "mytestvalue")
        {
        IsSecret = false,
        }, new ContainerRegistryRunArgument("mysecrettestargument", "mysecrettestvalue")
        {
        IsSecret = true,
        }},
        Target = "build",
        Values = {new ContainerRegistryTaskOverridableValue("mytestname", "mytestvalue")
        {
        IsSecret = false,
        }, new ContainerRegistryTaskOverridableValue("mysecrettestname", "mysecrettestvalue")
        {
        IsSecret = true,
        }},
        UpdateTriggerToken = "aGVsbG8gd29ybGQ=",
    },
};
ContainerRegistryRunResource result = await containerRegistry.ScheduleRunAsync(content);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryRunData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

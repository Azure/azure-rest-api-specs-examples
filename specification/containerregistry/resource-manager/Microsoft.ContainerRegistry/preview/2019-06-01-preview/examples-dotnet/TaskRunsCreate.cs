using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerRegistry.Models;
using Azure.ResourceManager.ContainerRegistry;

// Generated from example definition: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/TaskRunsCreate.json
// this example is just showing the usage of "TaskRuns_Create" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this ContainerRegistryTaskRunResource
ContainerRegistryTaskRunCollection collection = containerRegistry.GetContainerRegistryTaskRuns();

// invoke the operation
string taskRunName = "myRun";
ContainerRegistryTaskRunData data = new ContainerRegistryTaskRunData
{
    RunRequest = new ContainerRegistryEncodedTaskRunContent("c3RlcHM6IAogIC0gY21kOiB7eyAuVmFsdWVzLmNvbW1hbmQgfX0K", new ContainerRegistryPlatformProperties(ContainerRegistryOS.Linux)
    {
        Architecture = ContainerRegistryOSArchitecture.Amd64,
    })
    {
        EncodedValuesContent = "Y29tbWFuZDogYmFzaCBlY2hvIHt7LlJ1bi5SZWdpc3RyeX19Cg==",
        Values = { },
        Credentials = new ContainerRegistryCredentials(),
    },
    ForceUpdateTag = "test",
};
ArmOperation<ContainerRegistryTaskRunResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, taskRunName, data);
ContainerRegistryTaskRunResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerRegistryTaskRunData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

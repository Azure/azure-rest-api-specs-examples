using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Job_Executions_Get.json
// this example is just showing the usage of "JobsExecutions_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppJobResource created on azure
// for more information of creating ContainerAppJobResource, please refer to the document of ContainerAppJobResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string jobName = "testcontainerappsjob0";
ResourceIdentifier containerAppJobResourceId = ContainerAppJobResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, jobName);
ContainerAppJobResource containerAppJob = client.GetContainerAppJobResource(containerAppJobResourceId);

// get the collection of this ContainerAppJobExecutionResource
ContainerAppJobExecutionCollection collection = containerAppJob.GetContainerAppJobExecutions();

// invoke the operation and iterate over the result
await foreach (ContainerAppJobExecutionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ContainerAppJobExecutionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

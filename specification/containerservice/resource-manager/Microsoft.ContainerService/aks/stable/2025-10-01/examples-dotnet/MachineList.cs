using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/MachineList.json
// this example is just showing the usage of "Machines_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceAgentPoolResource created on azure
// for more information of creating ContainerServiceAgentPoolResource, please refer to the document of ContainerServiceAgentPoolResource
string subscriptionId = "26fe00f8-9173-4872-9134-bb1d2e00343a";
string resourceGroupName = "rg1";
string resourceName = "clustername1";
string agentPoolName = "agentpool1";
ResourceIdentifier containerServiceAgentPoolResourceId = ContainerServiceAgentPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, agentPoolName);
ContainerServiceAgentPoolResource containerServiceAgentPool = client.GetContainerServiceAgentPoolResource(containerServiceAgentPoolResourceId);

// get the collection of this ContainerServiceMachineResource
ContainerServiceMachineCollection collection = containerServiceAgentPool.GetContainerServiceMachines();

// invoke the operation and iterate over the result
await foreach (ContainerServiceMachineResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ContainerServiceMachineData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

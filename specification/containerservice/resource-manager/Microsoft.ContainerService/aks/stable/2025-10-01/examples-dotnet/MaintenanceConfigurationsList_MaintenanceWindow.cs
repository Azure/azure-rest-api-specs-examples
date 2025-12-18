using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerService.Models;
using Azure.ResourceManager.ContainerService;

// Generated from example definition: specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/MaintenanceConfigurationsList_MaintenanceWindow.json
// this example is just showing the usage of "MaintenanceConfigurations_ListByManagedCluster" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerServiceManagedClusterResource created on azure
// for more information of creating ContainerServiceManagedClusterResource, please refer to the document of ContainerServiceManagedClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string resourceName = "clustername1";
ResourceIdentifier containerServiceManagedClusterResourceId = ContainerServiceManagedClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ContainerServiceManagedClusterResource containerServiceManagedCluster = client.GetContainerServiceManagedClusterResource(containerServiceManagedClusterResourceId);

// get the collection of this ContainerServiceMaintenanceConfigurationResource
ContainerServiceMaintenanceConfigurationCollection collection = containerServiceManagedCluster.GetContainerServiceMaintenanceConfigurations();

// invoke the operation and iterate over the result
await foreach (ContainerServiceMaintenanceConfigurationResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ContainerServiceMaintenanceConfigurationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

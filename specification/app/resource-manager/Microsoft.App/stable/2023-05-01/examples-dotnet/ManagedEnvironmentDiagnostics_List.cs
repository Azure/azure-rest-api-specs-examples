using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ManagedEnvironmentDiagnostics_List.json
// this example is just showing the usage of "ManagedEnvironmentDiagnostics_ListDetectors" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppManagedEnvironmentResource created on azure
// for more information of creating ContainerAppManagedEnvironmentResource, please refer to the document of ContainerAppManagedEnvironmentResource
string subscriptionId = "f07f3711-b45e-40fe-a941-4e6d93f851e6";
string resourceGroupName = "mikono-workerapp-test-rg";
string environmentName = "mikonokubeenv";
ResourceIdentifier containerAppManagedEnvironmentResourceId = ContainerAppManagedEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, environmentName);
ContainerAppManagedEnvironmentResource containerAppManagedEnvironment = client.GetContainerAppManagedEnvironmentResource(containerAppManagedEnvironmentResourceId);

// get the collection of this ContainerAppManagedEnvironmentDetectorResource
ContainerAppManagedEnvironmentDetectorCollection collection = containerAppManagedEnvironment.GetContainerAppManagedEnvironmentDetectors();

// invoke the operation and iterate over the result
await foreach (ContainerAppManagedEnvironmentDetectorResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ContainerAppDiagnosticData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");

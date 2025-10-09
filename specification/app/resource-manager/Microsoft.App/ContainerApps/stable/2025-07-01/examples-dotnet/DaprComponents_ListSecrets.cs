using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/DaprComponents_ListSecrets.json
// this example is just showing the usage of "DaprComponents_ListSecrets" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppManagedEnvironmentDaprComponentResource created on azure
// for more information of creating ContainerAppManagedEnvironmentDaprComponentResource, please refer to the document of ContainerAppManagedEnvironmentDaprComponentResource
string subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
string resourceGroupName = "examplerg";
string environmentName = "myenvironment";
string componentName = "reddog";
ResourceIdentifier containerAppManagedEnvironmentDaprComponentResourceId = ContainerAppManagedEnvironmentDaprComponentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, environmentName, componentName);
ContainerAppManagedEnvironmentDaprComponentResource containerAppManagedEnvironmentDaprComponent = client.GetContainerAppManagedEnvironmentDaprComponentResource(containerAppManagedEnvironmentDaprComponentResourceId);

// invoke the operation and iterate over the result
await foreach (ContainerAppDaprSecret item in containerAppManagedEnvironmentDaprComponent.GetSecretsAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");

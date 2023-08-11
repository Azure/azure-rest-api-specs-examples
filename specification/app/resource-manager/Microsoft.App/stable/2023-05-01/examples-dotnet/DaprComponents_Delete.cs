using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;
using Azure.ResourceManager.AppContainers.Models;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/DaprComponents_Delete.json
// this example is just showing the usage of "DaprComponents_Delete" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
await containerAppManagedEnvironmentDaprComponent.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

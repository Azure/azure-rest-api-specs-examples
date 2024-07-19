using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ManagedEnvironmentsStorages_Delete.json
// this example is just showing the usage of "ManagedEnvironmentsStorages_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppManagedEnvironmentStorageResource created on azure
// for more information of creating ContainerAppManagedEnvironmentStorageResource, please refer to the document of ContainerAppManagedEnvironmentStorageResource
string subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
string resourceGroupName = "examplerg";
string environmentName = "managedEnv";
string storageName = "jlaw-demo1";
ResourceIdentifier containerAppManagedEnvironmentStorageResourceId = ContainerAppManagedEnvironmentStorageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, environmentName, storageName);
ContainerAppManagedEnvironmentStorageResource containerAppManagedEnvironmentStorage = client.GetContainerAppManagedEnvironmentStorageResource(containerAppManagedEnvironmentStorageResourceId);

// invoke the operation
await containerAppManagedEnvironmentStorage.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

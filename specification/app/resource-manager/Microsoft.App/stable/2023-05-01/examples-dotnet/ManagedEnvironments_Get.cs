using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ManagedEnvironments_Get.json
// this example is just showing the usage of "ManagedEnvironments_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppManagedEnvironmentResource created on azure
// for more information of creating ContainerAppManagedEnvironmentResource, please refer to the document of ContainerAppManagedEnvironmentResource
string subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
string resourceGroupName = "examplerg";
string environmentName = "jlaw-demo1";
ResourceIdentifier containerAppManagedEnvironmentResourceId = ContainerAppManagedEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, environmentName);
ContainerAppManagedEnvironmentResource containerAppManagedEnvironment = client.GetContainerAppManagedEnvironmentResource(containerAppManagedEnvironmentResourceId);

// invoke the operation
ContainerAppManagedEnvironmentResource result = await containerAppManagedEnvironment.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppManagedEnvironmentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

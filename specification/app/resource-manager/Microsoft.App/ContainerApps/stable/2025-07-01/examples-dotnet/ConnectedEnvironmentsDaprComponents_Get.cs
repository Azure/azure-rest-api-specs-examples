using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/ConnectedEnvironmentsDaprComponents_Get.json
// this example is just showing the usage of "ConnectedEnvironmentsDaprComponents_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppConnectedEnvironmentDaprComponentResource created on azure
// for more information of creating ContainerAppConnectedEnvironmentDaprComponentResource, please refer to the document of ContainerAppConnectedEnvironmentDaprComponentResource
string subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
string resourceGroupName = "examplerg";
string connectedEnvironmentName = "myenvironment";
string componentName = "reddog";
ResourceIdentifier containerAppConnectedEnvironmentDaprComponentResourceId = ContainerAppConnectedEnvironmentDaprComponentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, connectedEnvironmentName, componentName);
ContainerAppConnectedEnvironmentDaprComponentResource containerAppConnectedEnvironmentDaprComponent = client.GetContainerAppConnectedEnvironmentDaprComponentResource(containerAppConnectedEnvironmentDaprComponentResourceId);

// invoke the operation
ContainerAppConnectedEnvironmentDaprComponentResource result = await containerAppConnectedEnvironmentDaprComponent.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ContainerAppDaprComponentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

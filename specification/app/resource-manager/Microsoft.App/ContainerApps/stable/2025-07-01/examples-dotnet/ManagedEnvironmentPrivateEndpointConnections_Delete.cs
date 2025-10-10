using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/ManagedEnvironmentPrivateEndpointConnections_Delete.json
// this example is just showing the usage of "ManagedEnvironmentPrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppPrivateEndpointConnectionResource created on azure
// for more information of creating ContainerAppPrivateEndpointConnectionResource, please refer to the document of ContainerAppPrivateEndpointConnectionResource
string subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
string resourceGroupName = "examplerg";
string environmentName = "managedEnv";
string privateEndpointConnectionName = "jlaw-demo1";
ResourceIdentifier containerAppPrivateEndpointConnectionResourceId = ContainerAppPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, environmentName, privateEndpointConnectionName);
ContainerAppPrivateEndpointConnectionResource containerAppPrivateEndpointConnection = client.GetContainerAppPrivateEndpointConnectionResource(containerAppPrivateEndpointConnectionResourceId);

// invoke the operation
await containerAppPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

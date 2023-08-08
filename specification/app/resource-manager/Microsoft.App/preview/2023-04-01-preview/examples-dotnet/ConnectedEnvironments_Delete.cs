using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/preview/2023-04-01-preview/examples/ConnectedEnvironments_Delete.json
// this example is just showing the usage of "ConnectedEnvironments_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppConnectedEnvironmentResource created on azure
// for more information of creating ContainerAppConnectedEnvironmentResource, please refer to the document of ContainerAppConnectedEnvironmentResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "examplerg";
string connectedEnvironmentName = "examplekenv";
ResourceIdentifier containerAppConnectedEnvironmentResourceId = ContainerAppConnectedEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, connectedEnvironmentName);
ContainerAppConnectedEnvironmentResource containerAppConnectedEnvironment = client.GetContainerAppConnectedEnvironmentResource(containerAppConnectedEnvironmentResourceId);

// invoke the operation
await containerAppConnectedEnvironment.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerInstance.Models;
using Azure.ResourceManager.ContainerInstance;

// Generated from example definition: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/stable/2025-09-01/examples/ContainerGroupsStart.json
// this example is just showing the usage of "ContainerGroups_Start" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerGroupResource created on azure
// for more information of creating ContainerGroupResource, please refer to the document of ContainerGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "demo";
string containerGroupName = "demo1";
ResourceIdentifier containerGroupResourceId = ContainerGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, containerGroupName);
ContainerGroupResource containerGroup = client.GetContainerGroupResource(containerGroupResourceId);

// invoke the operation
await containerGroup.StartAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

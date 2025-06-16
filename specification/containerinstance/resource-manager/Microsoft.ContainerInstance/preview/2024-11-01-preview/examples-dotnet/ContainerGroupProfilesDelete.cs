using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ContainerInstance.Models;
using Azure.ResourceManager.ContainerInstance;

// Generated from example definition: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-11-01-preview/examples/ContainerGroupProfilesDelete.json
// this example is just showing the usage of "CGProfile_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerGroupProfileResource created on azure
// for more information of creating ContainerGroupProfileResource, please refer to the document of ContainerGroupProfileResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "demo";
string containerGroupProfileName = "demo1";
ResourceIdentifier containerGroupProfileResourceId = ContainerGroupProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, containerGroupProfileName);
ContainerGroupProfileResource containerGroupProfile = client.GetContainerGroupProfileResource(containerGroupProfileResourceId);

// invoke the operation
await containerGroupProfile.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

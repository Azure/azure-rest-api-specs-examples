using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppContainers;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ContainerApps_GetAuthToken.json
// this example is just showing the usage of "ContainerApps_GetAuthToken" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppResource created on azure
// for more information of creating ContainerAppResource, please refer to the document of ContainerAppResource
string subscriptionId = "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
string resourceGroupName = "rg";
string containerAppName = "testcontainerApp0";
ResourceIdentifier containerAppResourceId = ContainerAppResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, containerAppName);
ContainerAppResource containerApp = client.GetContainerAppResource(containerAppResourceId);

// invoke the operation
ContainerAppAuthToken result = await containerApp.GetAuthTokenAsync();

Console.WriteLine($"Succeeded: {result}");

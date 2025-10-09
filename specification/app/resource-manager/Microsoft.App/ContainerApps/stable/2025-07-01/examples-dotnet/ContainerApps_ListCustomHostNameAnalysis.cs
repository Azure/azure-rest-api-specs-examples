using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppContainers.Models;
using Azure.ResourceManager.AppContainers;

// Generated from example definition: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/ContainerApps_ListCustomHostNameAnalysis.json
// this example is just showing the usage of "ContainerApps_ListCustomHostNameAnalysis" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ContainerAppResource created on azure
// for more information of creating ContainerAppResource, please refer to the document of ContainerAppResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "rg";
string containerAppName = "testcontainerapp0";
ResourceIdentifier containerAppResourceId = ContainerAppResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, containerAppName);
ContainerAppResource containerApp = client.GetContainerAppResource(containerAppResourceId);

// invoke the operation
string customHostname = "my.name.corp";
ContainerAppCustomHostnameAnalysisResult result = await containerApp.GetCustomHostNameAnalysisAsync(customHostname: customHostname);

Console.WriteLine($"Succeeded: {result}");

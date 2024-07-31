using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/AppServiceEnvironments_ListCapacities.json
// this example is just showing the usage of "AppServiceEnvironments_ListCapacities" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppServiceEnvironmentResource created on azure
// for more information of creating AppServiceEnvironmentResource, please refer to the document of AppServiceEnvironmentResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "test-rg";
string name = "test-ase";
ResourceIdentifier appServiceEnvironmentResourceId = AppServiceEnvironmentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
AppServiceEnvironmentResource appServiceEnvironment = client.GetAppServiceEnvironmentResource(appServiceEnvironmentResourceId);

// invoke the operation and iterate over the result
await foreach (StampCapacity item in appServiceEnvironment.GetCapacitiesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");

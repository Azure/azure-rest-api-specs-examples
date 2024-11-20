using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-04-01/examples/AppServiceEnvironments_ListWorkerPools.json
// this example is just showing the usage of "AppServiceEnvironments_ListWorkerPools" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this HostingEnvironmentWorkerPoolResource
HostingEnvironmentWorkerPoolCollection collection = appServiceEnvironment.GetHostingEnvironmentWorkerPools();

// invoke the operation and iterate over the result
await foreach (HostingEnvironmentWorkerPoolResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AppServiceWorkerPoolData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

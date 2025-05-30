using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/ServiceRegistries_Get.json
// this example is just showing the usage of "ServiceRegistries_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformServiceRegistryResource created on azure
// for more information of creating AppPlatformServiceRegistryResource, please refer to the document of AppPlatformServiceRegistryResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string serviceRegistryName = "default";
ResourceIdentifier appPlatformServiceRegistryResourceId = AppPlatformServiceRegistryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, serviceRegistryName);
AppPlatformServiceRegistryResource appPlatformServiceRegistry = client.GetAppPlatformServiceRegistryResource(appPlatformServiceRegistryResourceId);

// invoke the operation
AppPlatformServiceRegistryResource result = await appPlatformServiceRegistry.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformServiceRegistryData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

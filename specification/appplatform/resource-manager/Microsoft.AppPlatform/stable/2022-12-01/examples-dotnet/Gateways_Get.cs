using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppPlatform;
using Azure.ResourceManager.AppPlatform.Models;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Gateways_Get.json
// this example is just showing the usage of "Gateways_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformServiceResource created on azure
// for more information of creating AppPlatformServiceResource, please refer to the document of AppPlatformServiceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
ResourceIdentifier appPlatformServiceResourceId = AppPlatformServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName);
AppPlatformServiceResource appPlatformService = client.GetAppPlatformServiceResource(appPlatformServiceResourceId);

// get the collection of this AppPlatformGatewayResource
AppPlatformGatewayCollection collection = appPlatformService.GetAppPlatformGateways();

// invoke the operation
string gatewayName = "default";
bool result = await collection.ExistsAsync(gatewayName);

Console.WriteLine($"Succeeded: {result}");

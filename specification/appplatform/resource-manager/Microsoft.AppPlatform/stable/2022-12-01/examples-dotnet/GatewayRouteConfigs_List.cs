using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppPlatform.Models;
using Azure.ResourceManager.AppPlatform;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/GatewayRouteConfigs_List.json
// this example is just showing the usage of "GatewayRouteConfigs_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppPlatformGatewayResource created on azure
// for more information of creating AppPlatformGatewayResource, please refer to the document of AppPlatformGatewayResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string serviceName = "myservice";
string gatewayName = "default";
ResourceIdentifier appPlatformGatewayResourceId = AppPlatformGatewayResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceName, gatewayName);
AppPlatformGatewayResource appPlatformGateway = client.GetAppPlatformGatewayResource(appPlatformGatewayResourceId);

// get the collection of this AppPlatformGatewayRouteConfigResource
AppPlatformGatewayRouteConfigCollection collection = appPlatformGateway.GetAppPlatformGatewayRouteConfigs();

// invoke the operation and iterate over the result
await foreach (AppPlatformGatewayRouteConfigResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AppPlatformGatewayRouteConfigData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

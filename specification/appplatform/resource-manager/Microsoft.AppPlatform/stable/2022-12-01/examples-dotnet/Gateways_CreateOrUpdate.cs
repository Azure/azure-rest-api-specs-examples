using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppPlatform;
using Azure.ResourceManager.AppPlatform.Models;

// Generated from example definition: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Gateways_CreateOrUpdate.json
// this example is just showing the usage of "Gateways_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
AppPlatformGatewayData data = new AppPlatformGatewayData()
{
    Properties = new AppPlatformGatewayProperties()
    {
        IsPublic = true,
        ResourceRequests = new AppPlatformGatewayResourceRequirements()
        {
            Cpu = "1",
            Memory = "1G",
        },
    },
    Sku = new AppPlatformSku()
    {
        Name = "E0",
        Tier = "Enterprise",
        Capacity = 2,
    },
};
ArmOperation<AppPlatformGatewayResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, gatewayName, data);
AppPlatformGatewayResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppPlatformGatewayData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

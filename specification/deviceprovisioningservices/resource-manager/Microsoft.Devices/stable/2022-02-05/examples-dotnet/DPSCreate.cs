using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceProvisioningServices.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DeviceProvisioningServices;

// Generated from example definition: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSCreate.json
// this example is just showing the usage of "IotDpsResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DeviceProvisioningServiceResource
DeviceProvisioningServiceCollection collection = resourceGroupResource.GetDeviceProvisioningServices();

// invoke the operation
string provisioningServiceName = "myFirstProvisioningService";
DeviceProvisioningServiceData data = new DeviceProvisioningServiceData(new AzureLocation("East US"), new DeviceProvisioningServiceProperties
{
    IsDataResidencyEnabled = false,
}, new DeviceProvisioningServicesSkuInfo
{
    Name = DeviceProvisioningServicesSku.S1,
    Capacity = 1L,
})
{
    Tags = { },
};
ArmOperation<DeviceProvisioningServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, provisioningServiceName, data);
DeviceProvisioningServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DeviceProvisioningServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

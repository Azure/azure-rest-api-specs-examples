using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceProvisioningServices.Models;
using Azure.ResourceManager.DeviceProvisioningServices;

// Generated from example definition: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSGetCertificates.json
// this example is just showing the usage of "DpsCertificate_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceProvisioningServiceResource created on azure
// for more information of creating DeviceProvisioningServiceResource, please refer to the document of DeviceProvisioningServiceResource
string subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
string resourceGroupName = "myResourceGroup";
string provisioningServiceName = "myFirstProvisioningService";
ResourceIdentifier deviceProvisioningServiceResourceId = DeviceProvisioningServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, provisioningServiceName);
DeviceProvisioningServiceResource deviceProvisioningService = client.GetDeviceProvisioningServiceResource(deviceProvisioningServiceResourceId);

// get the collection of this DeviceProvisioningServicesCertificateResource
DeviceProvisioningServicesCertificateCollection collection = deviceProvisioningService.GetDeviceProvisioningServicesCertificates();

// invoke the operation and iterate over the result
await foreach (DeviceProvisioningServicesCertificateResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DeviceProvisioningServicesCertificateData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeviceProvisioningServices;
using Azure.ResourceManager.DeviceProvisioningServices.Models;

// Generated from example definition: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSGetCertificate.json
// this example is just showing the usage of "DpsCertificate_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceProvisioningServicesCertificateResource created on azure
// for more information of creating DeviceProvisioningServicesCertificateResource, please refer to the document of DeviceProvisioningServicesCertificateResource
string subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
string resourceGroupName = "myResourceGroup";
string provisioningServiceName = "myFirstProvisioningService";
string certificateName = "cert";
ResourceIdentifier deviceProvisioningServicesCertificateResourceId = DeviceProvisioningServicesCertificateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, provisioningServiceName, certificateName);
DeviceProvisioningServicesCertificateResource deviceProvisioningServicesCertificate = client.GetDeviceProvisioningServicesCertificateResource(deviceProvisioningServicesCertificateResourceId);

// invoke the operation
DeviceProvisioningServicesCertificateResource result = await deviceProvisioningServicesCertificate.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DeviceProvisioningServicesCertificateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

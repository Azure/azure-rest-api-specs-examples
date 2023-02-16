using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeviceProvisioningServices;
using Azure.ResourceManager.DeviceProvisioningServices.Models;

// Generated from example definition: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSGetPrivateEndpointConnection.json
// this example is just showing the usage of "IotDpsResource_GetPrivateEndpointConnection" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceProvisioningServiceResource created on azure
// for more information of creating DeviceProvisioningServiceResource, please refer to the document of DeviceProvisioningServiceResource
string subscriptionId = "91d12660-3dec-467a-be2a-213b5544ddc0";
string resourceGroupName = "myResourceGroup";
string resourceName = "myFirstProvisioningService";
ResourceIdentifier deviceProvisioningServiceResourceId = DeviceProvisioningServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
DeviceProvisioningServiceResource deviceProvisioningService = client.GetDeviceProvisioningServiceResource(deviceProvisioningServiceResourceId);

// get the collection of this DeviceProvisioningServicesPrivateEndpointConnectionResource
DeviceProvisioningServicesPrivateEndpointConnectionCollection collection = deviceProvisioningService.GetDeviceProvisioningServicesPrivateEndpointConnections();

// invoke the operation
string privateEndpointConnectionName = "myPrivateEndpointConnection";
bool result = await collection.ExistsAsync(privateEndpointConnectionName);

Console.WriteLine($"Succeeded: {result}");

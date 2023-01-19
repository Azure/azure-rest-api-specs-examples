using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DataBoxEdge;
using Azure.ResourceManager.DataBoxEdge.Models;

// Generated from example definition: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/RolePut.json
// this example is just showing the usage of "Roles_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DataBoxEdgeDeviceResource created on azure
// for more information of creating DataBoxEdgeDeviceResource, please refer to the document of DataBoxEdgeDeviceResource
string subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
string resourceGroupName = "GroupForEdgeAutomation";
string deviceName = "testedgedevice";
ResourceIdentifier dataBoxEdgeDeviceResourceId = DataBoxEdgeDeviceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, deviceName);
DataBoxEdgeDeviceResource dataBoxEdgeDevice = client.GetDataBoxEdgeDeviceResource(dataBoxEdgeDeviceResourceId);

// get the collection of this DataBoxEdgeRoleResource
DataBoxEdgeRoleCollection collection = dataBoxEdgeDevice.GetDataBoxEdgeRoles();

// invoke the operation
string name = "IoTRole1";
DataBoxEdgeRoleData data = new EdgeIotRole()
{
    HostPlatform = DataBoxEdgeOSPlatformType.Linux,
    IotDeviceDetails = new EdgeIotDeviceInfo("iotdevice", "iothub.azure-devices.net")
    {
        SymmetricKeyConnectionString = new AsymmetricEncryptedSecret("Encrypted<<HostName=iothub.azure-devices.net;DeviceId=iotDevice;SharedAccessKey=2C750FscEas3JmQ8Bnui5yQWZPyml0/UiRt1bQwd8=>>", DataBoxEdgeEncryptionAlgorithm.Aes256)
        {
            EncryptionCertThumbprint = "348586569999244",
        },
    },
    IotEdgeDeviceDetails = new EdgeIotDeviceInfo("iotEdge", "iothub.azure-devices.net")
    {
        SymmetricKeyConnectionString = new AsymmetricEncryptedSecret("Encrypted<<HostName=iothub.azure-devices.net;DeviceId=iotEdge;SharedAccessKey=2C750FscEas3JmQ8Bnui5yQWZPyml0/UiRt1bQwd8=>>", DataBoxEdgeEncryptionAlgorithm.Aes256)
        {
            EncryptionCertThumbprint = "1245475856069999244",
        },
    },
    ShareMappings =
    {
    },
    RoleStatus = DataBoxEdgeRoleStatus.Enabled,
};
ArmOperation<DataBoxEdgeRoleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, name, data);
DataBoxEdgeRoleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DataBoxEdgeRoleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

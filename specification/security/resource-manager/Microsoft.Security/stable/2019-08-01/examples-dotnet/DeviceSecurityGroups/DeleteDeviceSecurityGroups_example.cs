using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/DeviceSecurityGroups/DeleteDeviceSecurityGroups_example.json
// this example is just showing the usage of "DeviceSecurityGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceSecurityGroupResource created on azure
// for more information of creating DeviceSecurityGroupResource, please refer to the document of DeviceSecurityGroupResource
string resourceId = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Devices/iotHubs/sampleiothub";
string deviceSecurityGroupName = "samplesecuritygroup";
ResourceIdentifier deviceSecurityGroupResourceId = DeviceSecurityGroupResource.CreateResourceIdentifier(resourceId, deviceSecurityGroupName);
DeviceSecurityGroupResource deviceSecurityGroup = client.GetDeviceSecurityGroupResource(deviceSecurityGroupResourceId);

// invoke the operation
await deviceSecurityGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");

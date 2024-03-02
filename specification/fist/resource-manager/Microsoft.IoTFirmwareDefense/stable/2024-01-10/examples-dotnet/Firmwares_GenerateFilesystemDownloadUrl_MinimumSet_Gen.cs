using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.IotFirmwareDefense;
using Azure.ResourceManager.IotFirmwareDefense.Models;

// Generated from example definition: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/Firmwares_GenerateFilesystemDownloadUrl_MinimumSet_Gen.json
// this example is just showing the usage of "Firmwares_GenerateFilesystemDownloadUrl" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotFirmwareResource created on azure
// for more information of creating IotFirmwareResource, please refer to the document of IotFirmwareResource
string subscriptionId = "685C0C6F-9867-4B1C-A534-AA3A05B54BCE";
string resourceGroupName = "rgworkspaces-firmwares";
string workspaceName = "A7";
string firmwareId = "umrkdttp";
ResourceIdentifier iotFirmwareResourceId = IotFirmwareResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, firmwareId);
IotFirmwareResource iotFirmware = client.GetIotFirmwareResource(iotFirmwareResourceId);

// invoke the operation
FirmwareUriToken result = await iotFirmware.GenerateFilesystemDownloadUriAsync();

Console.WriteLine($"Succeeded: {result}");

using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.IotFirmwareDefense;
using Azure.ResourceManager.IotFirmwareDefense.Models;

// Generated from example definition: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateCveSummary_MinimumSet_Gen.json
// this example is just showing the usage of "Firmware_GenerateCveSummary" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirmwareResource created on azure
// for more information of creating FirmwareResource, please refer to the document of FirmwareResource
string subscriptionId = "685C0C6F-9867-4B1C-A534-AA3A05B54BCE";
string resourceGroupName = "rgworkspaces-firmwares";
string workspaceName = "A7";
string firmwareName = "umrkdttp";
ResourceIdentifier firmwareResourceId = FirmwareResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, firmwareName);
FirmwareResource firmware = client.GetFirmwareResource(firmwareResourceId);

// invoke the operation
CveSummary result = await firmware.GetCveSummaryAsync();

Console.WriteLine($"Succeeded: {result}");

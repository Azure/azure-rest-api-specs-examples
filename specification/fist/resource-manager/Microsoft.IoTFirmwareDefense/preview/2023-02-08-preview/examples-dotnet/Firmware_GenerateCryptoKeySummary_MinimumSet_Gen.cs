using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.IotFirmwareDefense;
using Azure.ResourceManager.IotFirmwareDefense.Models;

// Generated from example definition: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Firmware_GenerateCryptoKeySummary_MinimumSet_Gen.json
// this example is just showing the usage of "Firmware_GenerateCryptoKeySummary" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirmwareResource created on azure
// for more information of creating FirmwareResource, please refer to the document of FirmwareResource
string subscriptionId = "C589E84A-5C11-4A25-9CF9-4E9C2F1EBFCA";
string resourceGroupName = "rgworkspaces-firmwares";
string workspaceName = "j5QE_";
string firmwareName = "wujtpcgypfpqseyrsebolarkspy";
ResourceIdentifier firmwareResourceId = FirmwareResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, firmwareName);
FirmwareResource firmware = client.GetFirmwareResource(firmwareResourceId);

// invoke the operation
FirmwareCryptoKeySummary result = await firmware.GetCryptoKeySummaryAsync();

Console.WriteLine($"Succeeded: {result}");

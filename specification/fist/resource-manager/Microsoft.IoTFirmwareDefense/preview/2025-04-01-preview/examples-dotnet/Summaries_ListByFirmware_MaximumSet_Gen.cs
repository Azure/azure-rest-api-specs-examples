using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotFirmwareDefense.Models;
using Azure.ResourceManager.IotFirmwareDefense;

// Generated from example definition: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2025-04-01-preview/examples/Summaries_ListByFirmware_MaximumSet_Gen.json
// this example is just showing the usage of "Summaries_ListByFirmware" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotFirmwareResource created on azure
// for more information of creating IotFirmwareResource, please refer to the document of IotFirmwareResource
string subscriptionId = "5C707B5F-6130-4F71-819E-953A28942E88";
string resourceGroupName = "rgiotfirmwaredefense";
string workspaceName = "exampleWorkspaceName";
string firmwareId = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier iotFirmwareResourceId = IotFirmwareResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, firmwareId);
IotFirmwareResource iotFirmware = client.GetIotFirmwareResource(iotFirmwareResourceId);

// get the collection of this FirmwareAnalysisSummaryResource
FirmwareAnalysisSummaryCollection collection = iotFirmware.GetFirmwareAnalysisSummaries();

// invoke the operation and iterate over the result
await foreach (FirmwareAnalysisSummaryResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    FirmwareAnalysisSummaryData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");

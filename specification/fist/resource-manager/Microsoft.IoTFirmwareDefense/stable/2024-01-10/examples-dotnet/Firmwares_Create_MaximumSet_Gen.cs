using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.IotFirmwareDefense;
using Azure.ResourceManager.IotFirmwareDefense.Models;

// Generated from example definition: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/Firmwares_Create_MaximumSet_Gen.json
// this example is just showing the usage of "Firmwares_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirmwareAnalysisWorkspaceResource created on azure
// for more information of creating FirmwareAnalysisWorkspaceResource, please refer to the document of FirmwareAnalysisWorkspaceResource
string subscriptionId = "685C0C6F-9867-4B1C-A534-AA3A05B54BCE";
string resourceGroupName = "rgworkspaces-firmwares";
string workspaceName = "A7";
ResourceIdentifier firmwareAnalysisWorkspaceResourceId = FirmwareAnalysisWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
FirmwareAnalysisWorkspaceResource firmwareAnalysisWorkspace = client.GetFirmwareAnalysisWorkspaceResource(firmwareAnalysisWorkspaceResourceId);

// get the collection of this IotFirmwareResource
IotFirmwareCollection collection = firmwareAnalysisWorkspace.GetIotFirmwares();

// invoke the operation
string firmwareId = "umrkdttp";
IotFirmwareData data = new IotFirmwareData()
{
    FileName = "wresexxulcdsdd",
    Vendor = "vycmdhgtmepcptyoubztiuudpkcpd",
    Model = "f",
    Version = "s",
    Description = "uz",
    FileSize = 17,
    Status = FirmwareAnalysisStatus.Pending,
    StatusMessages =
    {
    new FirmwareAnalysisStatusMessage()
    {
    Message = "ulvhmhokezathzzauiitu",
    }
    },
};
ArmOperation<IotFirmwareResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, firmwareId, data);
IotFirmwareResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotFirmwareData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

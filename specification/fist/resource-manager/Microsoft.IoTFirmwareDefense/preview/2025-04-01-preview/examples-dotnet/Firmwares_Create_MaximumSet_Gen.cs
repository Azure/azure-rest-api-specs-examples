using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotFirmwareDefense.Models;
using Azure.ResourceManager.IotFirmwareDefense;

// Generated from example definition: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2025-04-01-preview/examples/Firmwares_Create_MaximumSet_Gen.json
// this example is just showing the usage of "Firmwares_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirmwareAnalysisWorkspaceResource created on azure
// for more information of creating FirmwareAnalysisWorkspaceResource, please refer to the document of FirmwareAnalysisWorkspaceResource
string subscriptionId = "5C707B5F-6130-4F71-819E-953A28942E88";
string resourceGroupName = "rgiotfirmwaredefense";
string workspaceName = "exampleWorkspaceName";
ResourceIdentifier firmwareAnalysisWorkspaceResourceId = FirmwareAnalysisWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
FirmwareAnalysisWorkspaceResource firmwareAnalysisWorkspace = client.GetFirmwareAnalysisWorkspaceResource(firmwareAnalysisWorkspaceResourceId);

// get the collection of this IotFirmwareResource
IotFirmwareCollection collection = firmwareAnalysisWorkspace.GetIotFirmwares();

// invoke the operation
string firmwareId = "00000000-0000-0000-0000-000000000000";
IotFirmwareData data = new IotFirmwareData
{
    FileName = "dmnqhyxssutvnewntlb",
    Vendor = "hymojocxpxqhtblioaavylnzyg",
    Model = "wmyfbyjsggbvxcuin",
    Version = "nhtxzslgcbtptu",
    Description = "sqt",
    FileSize = 30L,
    Status = FirmwareAnalysisStatus.Pending,
    StatusMessages = {new FirmwareAnalysisStatusMessage
    {
    ErrorCode = 20L,
    Message = "edtylkjvj",
    }},
};
ArmOperation<IotFirmwareResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, firmwareId, data);
IotFirmwareResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotFirmwareData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

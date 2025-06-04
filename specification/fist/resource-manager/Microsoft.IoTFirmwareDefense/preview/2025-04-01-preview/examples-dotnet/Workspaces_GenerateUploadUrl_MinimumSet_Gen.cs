using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotFirmwareDefense.Models;
using Azure.ResourceManager.IotFirmwareDefense;

// Generated from example definition: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2025-04-01-preview/examples/Workspaces_GenerateUploadUrl_MinimumSet_Gen.json
// this example is just showing the usage of "Workspaces_GenerateUploadUri" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirmwareAnalysisWorkspaceResource created on azure
// for more information of creating FirmwareAnalysisWorkspaceResource, please refer to the document of FirmwareAnalysisWorkspaceResource
string subscriptionId = "5443A01A-5242-4950-AC1A-2DD362180254";
string resourceGroupName = "rgworkspaces";
string workspaceName = "E___-3";
ResourceIdentifier firmwareAnalysisWorkspaceResourceId = FirmwareAnalysisWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
FirmwareAnalysisWorkspaceResource firmwareAnalysisWorkspace = client.GetFirmwareAnalysisWorkspaceResource(firmwareAnalysisWorkspaceResourceId);

// invoke the operation
FirmwareUploadUriContent content = new FirmwareUploadUriContent
{
    FirmwareId = "ktnnf",
};
FirmwareUriToken result = await firmwareAnalysisWorkspace.GenerateUploadUriAsync(content);

Console.WriteLine($"Succeeded: {result}");

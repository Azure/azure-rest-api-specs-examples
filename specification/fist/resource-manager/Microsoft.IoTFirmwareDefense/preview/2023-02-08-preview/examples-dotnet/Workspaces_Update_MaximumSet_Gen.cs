using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.IotFirmwareDefense;
using Azure.ResourceManager.IotFirmwareDefense.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Workspaces_Update_MaximumSet_Gen.json
// this example is just showing the usage of "Workspaces_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirmwareWorkspaceResource created on azure
// for more information of creating FirmwareWorkspaceResource, please refer to the document of FirmwareWorkspaceResource
string subscriptionId = "5443A01A-5242-4950-AC1A-2DD362180254";
string resourceGroupName = "rgworkspaces";
string workspaceName = "E___-3";
ResourceIdentifier firmwareWorkspaceResourceId = FirmwareWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
FirmwareWorkspaceResource firmwareWorkspace = client.GetFirmwareWorkspaceResource(firmwareWorkspaceResourceId);

// invoke the operation
FirmwareWorkspacePatch patch = new FirmwareWorkspacePatch();
FirmwareWorkspaceResource result = await firmwareWorkspace.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FirmwareWorkspaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

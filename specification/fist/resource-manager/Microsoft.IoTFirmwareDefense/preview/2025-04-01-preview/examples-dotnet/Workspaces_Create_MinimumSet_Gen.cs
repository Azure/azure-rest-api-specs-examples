using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotFirmwareDefense.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.IotFirmwareDefense;

// Generated from example definition: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2025-04-01-preview/examples/Workspaces_Create_MinimumSet_Gen.json
// this example is just showing the usage of "Workspaces_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "5443A01A-5242-4950-AC1A-2DD362180254";
string resourceGroupName = "rgworkspaces";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this FirmwareAnalysisWorkspaceResource
FirmwareAnalysisWorkspaceCollection collection = resourceGroupResource.GetFirmwareAnalysisWorkspaces();

// invoke the operation
string workspaceName = "E___-3";
FirmwareAnalysisWorkspaceData data = new FirmwareAnalysisWorkspaceData(new AzureLocation("emiscxuo"));
ArmOperation<FirmwareAnalysisWorkspaceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, workspaceName, data);
FirmwareAnalysisWorkspaceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FirmwareAnalysisWorkspaceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

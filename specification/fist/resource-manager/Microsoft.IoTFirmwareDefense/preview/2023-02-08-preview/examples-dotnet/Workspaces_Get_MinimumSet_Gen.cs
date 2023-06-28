using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.IotFirmwareDefense;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Workspaces_Get_MinimumSet_Gen.json
// this example is just showing the usage of "Workspaces_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "9781B4B5-0922-472A-80F0-B743D0596694";
string resourceGroupName = "rgworkspaces";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this FirmwareWorkspaceResource
FirmwareWorkspaceCollection collection = resourceGroupResource.GetFirmwareWorkspaces();

// invoke the operation
string workspaceName = "E_US";
bool result = await collection.ExistsAsync(workspaceName);

Console.WriteLine($"Succeeded: {result}");

using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ConnectedVMwarevSphere;
using Azure.ResourceManager.ConnectedVMwarevSphere.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2020-10-01-preview/examples/DeleteVCenter.json
// this example is just showing the usage of "VCenters_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VCenterResource created on azure
// for more information of creating VCenterResource, please refer to the document of VCenterResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
string vcenterName = "ContosoVCenter";
ResourceIdentifier vCenterResourceId = VCenterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, vcenterName);
VCenterResource vCenter = client.GetVCenterResource(vCenterResourceId);

// invoke the operation
await vCenter.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

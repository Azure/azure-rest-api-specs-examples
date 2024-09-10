using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HardwareSecurityModules.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.HardwareSecurityModules;

// Generated from example definition: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/PaymentHsm_Update.json
// this example is just showing the usage of "DedicatedHsm_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DedicatedHsmResource created on azure
// for more information of creating DedicatedHsmResource, please refer to the document of DedicatedHsmResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "hsm-group";
string name = "hsm1";
ResourceIdentifier dedicatedHsmResourceId = DedicatedHsmResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
DedicatedHsmResource dedicatedHsm = client.GetDedicatedHsmResource(dedicatedHsmResourceId);

// invoke the operation
DedicatedHsmPatch patch = new DedicatedHsmPatch()
{
    Tags =
    {
    ["Dept"] = "hsm",
    ["Environment"] = "dogfood",
    ["Slice"] = "A",
    },
};
ArmOperation<DedicatedHsmResource> lro = await dedicatedHsm.UpdateAsync(WaitUntil.Completed, patch);
DedicatedHsmResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DedicatedHsmData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHost_Update_MaximumSet_Gen.json
// this example is just showing the usage of "DedicatedHosts_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DedicatedHostResource created on azure
// for more information of creating DedicatedHostResource, please refer to the document of DedicatedHostResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string hostGroupName = "aaaaaaaaa";
string hostName = "aaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier dedicatedHostResourceId = DedicatedHostResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostGroupName, hostName);
DedicatedHostResource dedicatedHost = client.GetDedicatedHostResource(dedicatedHostResourceId);

// invoke the operation
DedicatedHostPatch patch = new DedicatedHostPatch()
{
    PlatformFaultDomain = 1,
    AutoReplaceOnFailure = true,
    LicenseType = DedicatedHostLicenseType.WindowsServerHybrid,
    Tags =
    {
    ["key8813"] = "aaaaaaaaaaaaaaaaaaaaaaaaaaa",
    },
};
ArmOperation<DedicatedHostResource> lro = await dedicatedHost.UpdateAsync(WaitUntil.Completed, patch);
DedicatedHostResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DedicatedHostData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/dedicatedHostExamples/DedicatedHostGroup_Update_MaximumSet_Gen.json
// this example is just showing the usage of "DedicatedHostGroups_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DedicatedHostGroupResource created on azure
// for more information of creating DedicatedHostGroupResource, please refer to the document of DedicatedHostGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string hostGroupName = "aaaa";
ResourceIdentifier dedicatedHostGroupResourceId = DedicatedHostGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, hostGroupName);
DedicatedHostGroupResource dedicatedHostGroup = client.GetDedicatedHostGroupResource(dedicatedHostGroupResourceId);

// invoke the operation
DedicatedHostGroupPatch patch = new DedicatedHostGroupPatch()
{
    Zones =
    {
    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
    },
    PlatformFaultDomainCount = 3,
    SupportAutomaticPlacement = true,
    Tags =
    {
    ["key9921"] = "aaaaaaaaaa",
    },
};
DedicatedHostGroupResource result = await dedicatedHostGroup.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DedicatedHostGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

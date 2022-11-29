using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/availabilitySetExamples/AvailabilitySets_Update_MaximumSet_Gen.json
// this example is just showing the usage of "AvailabilitySets_Update" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this AvailabilitySetResource created on azure
// for more information of creating AvailabilitySetResource, please refer to the document of AvailabilitySetResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string availabilitySetName = "aaaaaaaaaaaaaaaaaaa";
ResourceIdentifier availabilitySetResourceId = AvailabilitySetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, availabilitySetName);
AvailabilitySetResource availabilitySet = client.GetAvailabilitySetResource(availabilitySetResourceId);

// invoke the operation
AvailabilitySetPatch patch = new AvailabilitySetPatch()
{
    Sku = new ComputeSku()
    {
        Name = "DSv3-Type1",
        Tier = "aaa",
        Capacity = 7,
    },
    PlatformUpdateDomainCount = 20,
    PlatformFaultDomainCount = 2,
    VirtualMachines =
    {
    new WritableSubResource()
    {
    Id = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
    }
    },
    ProximityPlacementGroupId = new ResourceIdentifier("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
    Tags =
    {
    ["key2574"] = "aaaaaaaa",
    },
};
AvailabilitySetResource result = await availabilitySet.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AvailabilitySetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

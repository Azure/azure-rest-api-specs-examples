using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/capacityReservationExamples/CapacityReservationGroup_Update_MaximumSet_Gen.json
// this example is just showing the usage of "CapacityReservationGroups_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CapacityReservationGroupResource created on azure
// for more information of creating CapacityReservationGroupResource, please refer to the document of CapacityReservationGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string capacityReservationGroupName = "aaaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier capacityReservationGroupResourceId = CapacityReservationGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, capacityReservationGroupName);
CapacityReservationGroupResource capacityReservationGroup = client.GetCapacityReservationGroupResource(capacityReservationGroupResourceId);

// invoke the operation
CapacityReservationGroupPatch patch = new CapacityReservationGroupPatch()
{
    Tags =
    {
    ["key5355"] = "aaa",
    },
};
CapacityReservationGroupResource result = await capacityReservationGroup.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CapacityReservationGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/capacityReservationExamples/CapacityReservation_CreateOrUpdate.json
// this example is just showing the usage of "CapacityReservations_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CapacityReservationGroupResource created on azure
// for more information of creating CapacityReservationGroupResource, please refer to the document of CapacityReservationGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string capacityReservationGroupName = "myCapacityReservationGroup";
ResourceIdentifier capacityReservationGroupResourceId = CapacityReservationGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, capacityReservationGroupName);
CapacityReservationGroupResource capacityReservationGroup = client.GetCapacityReservationGroupResource(capacityReservationGroupResourceId);

// get the collection of this CapacityReservationResource
CapacityReservationCollection collection = capacityReservationGroup.GetCapacityReservations();

// invoke the operation
string capacityReservationName = "myCapacityReservation";
CapacityReservationData data = new CapacityReservationData(new AzureLocation("westus"), new ComputeSku()
{
    Name = "Standard_DS1_v2",
    Capacity = 4,
})
{
    Zones =
    {
    "1"
    },
    Tags =
    {
    ["department"] = "HR",
    },
};
ArmOperation<CapacityReservationResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, capacityReservationName, data);
CapacityReservationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CapacityReservationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

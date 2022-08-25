using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/capacityReservationExamples/CapacityReservation_ListByReservationGroup.json
// this example is just showing the usage of "CapacityReservations_ListByCapacityReservationGroup" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this CapacityReservationGroupResource created on azure
// for more information of creating CapacityReservationGroupResource, please refer to the document of CapacityReservationGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string capacityReservationGroupName = "myCapacityReservationGroup";
ResourceIdentifier capacityReservationGroupResourceId = CapacityReservationGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, capacityReservationGroupName);
CapacityReservationGroupResource capacityReservationGroup = client.GetCapacityReservationGroupResource(capacityReservationGroupResourceId);
            
// get the collection of this CapacityReservationResource
CapacityReservationCollection collection = capacityReservationGroup.GetCapacityReservations();
            
// invoke the operation and iterate over the result
await foreach (CapacityReservationResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    CapacityReservationData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
            
Console.WriteLine($"Succeeded");

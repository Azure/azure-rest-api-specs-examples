using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/capacityReservationExamples/CapacityReservation_Update_MinimumSet_Gen.json
// this example is just showing the usage of "CapacityReservations_Update" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this CapacityReservationResource created on azure
// for more information of creating CapacityReservationResource, please refer to the document of CapacityReservationResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string capacityReservationGroupName = "aaaaaaaaaaaaaaaaaaaaaaaaaa";
string capacityReservationName = "aaa";
ResourceIdentifier capacityReservationResourceId = CapacityReservationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, capacityReservationGroupName, capacityReservationName);
CapacityReservationResource capacityReservation = client.GetCapacityReservationResource(capacityReservationResourceId);
            
// invoke the operation
CapacityReservationPatch patch = new CapacityReservationPatch();
ArmOperation<CapacityReservationResource> lro = await capacityReservation.UpdateAsync(WaitUntil.Completed, patch);
CapacityReservationResource result = lro.Value;
            
// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CapacityReservationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

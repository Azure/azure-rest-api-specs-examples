using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/capacityReservationExamples/CapacityReservationGroup_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "CapacityReservationGroups_Delete" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this CapacityReservationGroupResource created on azure
// for more information of creating CapacityReservationGroupResource, please refer to the document of CapacityReservationGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string capacityReservationGroupName = "aaaaaaaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier capacityReservationGroupResourceId = CapacityReservationGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, capacityReservationGroupName);
CapacityReservationGroupResource capacityReservationGroup = client.GetCapacityReservationGroupResource(capacityReservationGroupResourceId);
            
// invoke the operation
await capacityReservationGroup.DeleteAsync(WaitUntil.Completed);
            
Console.WriteLine($"Succeeded");

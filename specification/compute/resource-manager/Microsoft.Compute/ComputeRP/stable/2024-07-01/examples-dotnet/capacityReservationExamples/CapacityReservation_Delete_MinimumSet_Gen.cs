using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/capacityReservationExamples/CapacityReservation_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "CapacityReservations_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CapacityReservationResource created on azure
// for more information of creating CapacityReservationResource, please refer to the document of CapacityReservationResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string capacityReservationGroupName = "aaa";
string capacityReservationName = "aaaaaa";
ResourceIdentifier capacityReservationResourceId = CapacityReservationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, capacityReservationGroupName, capacityReservationName);
CapacityReservationResource capacityReservation = client.GetCapacityReservationResource(capacityReservationResourceId);

// invoke the operation
await capacityReservation.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

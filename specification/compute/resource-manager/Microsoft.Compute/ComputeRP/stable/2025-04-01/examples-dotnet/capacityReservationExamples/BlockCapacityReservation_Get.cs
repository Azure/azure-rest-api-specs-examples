using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/capacityReservationExamples/BlockCapacityReservation_Get.json
// this example is just showing the usage of "CapacityReservations_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CapacityReservationResource created on azure
// for more information of creating CapacityReservationResource, please refer to the document of CapacityReservationResource
string subscriptionId = "{subscriptionId}";
string resourceGroupName = "myResourceGroup";
string capacityReservationGroupName = "blockCapacityReservationGroup";
string capacityReservationName = "blockCapacityReservation";
ResourceIdentifier capacityReservationResourceId = CapacityReservationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, capacityReservationGroupName, capacityReservationName);
CapacityReservationResource capacityReservation = client.GetCapacityReservationResource(capacityReservationResourceId);

// invoke the operation
CapacityReservationInstanceViewType? expand = CapacityReservationInstanceViewType.InstanceView;
CapacityReservationResource result = await capacityReservation.GetAsync(expand: expand);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CapacityReservationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

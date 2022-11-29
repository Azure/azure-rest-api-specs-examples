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

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/availabilitySetExamples/AvailabilitySets_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "AvailabilitySets_Delete" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this AvailabilitySetResource created on azure
// for more information of creating AvailabilitySetResource, please refer to the document of AvailabilitySetResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string availabilitySetName = "aaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier availabilitySetResourceId = AvailabilitySetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, availabilitySetName);
AvailabilitySetResource availabilitySet = client.GetAvailabilitySetResource(availabilitySetResourceId);

// invoke the operation
await availabilitySet.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");

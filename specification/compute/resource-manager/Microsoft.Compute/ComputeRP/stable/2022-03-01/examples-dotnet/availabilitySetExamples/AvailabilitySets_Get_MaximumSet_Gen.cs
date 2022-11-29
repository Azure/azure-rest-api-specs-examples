using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/availabilitySetExamples/AvailabilitySets_Get_MaximumSet_Gen.json
// this example is just showing the usage of "AvailabilitySets_Get" operation, for the dependent resources, they will have to be created separately.

// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AvailabilitySetResource
AvailabilitySetCollection collection = resourceGroupResource.GetAvailabilitySets();

// invoke the operation
string availabilitySetName = "aaaaaaaaaaaa";
bool result = await collection.ExistsAsync(availabilitySetName);

Console.WriteLine($"Succeeded: {result}");

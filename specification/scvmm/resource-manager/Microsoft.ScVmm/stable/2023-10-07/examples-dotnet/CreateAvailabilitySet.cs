using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.ScVmm;

// Generated from example definition: specification/scvmm/resource-manager/Microsoft.ScVmm/stable/2023-10-07/examples/CreateAvailabilitySet.json
// this example is just showing the usage of "AvailabilitySets_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "testrg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ScVmmAvailabilitySetResource
ScVmmAvailabilitySetCollection collection = resourceGroupResource.GetScVmmAvailabilitySets();

// invoke the operation
string availabilitySetResourceName = "HRAvailabilitySet";
ScVmmAvailabilitySetData data = new ScVmmAvailabilitySetData(new AzureLocation("East US"), new ExtendedLocation
{
    Name = "/subscriptions/a5015e1c-867f-4533-8541-85cd470d0cfb/resourceGroups/demoRG/providers/Microsoft.Arc/customLocations/contoso",
})
{
    AvailabilitySetName = "hr-avset",
    VmmServerId = new ResourceIdentifier("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.ScVmm/VMMServers/ContosoVMMServer"),
};
ArmOperation<ScVmmAvailabilitySetResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, availabilitySetResourceName, data);
ScVmmAvailabilitySetResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScVmmAvailabilitySetData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

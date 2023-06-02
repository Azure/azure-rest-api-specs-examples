using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/proximityPlacementGroupExamples/ProximityPlacementGroup_CreateOrUpdate.json
// this example is just showing the usage of "ProximityPlacementGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ProximityPlacementGroupResource
ProximityPlacementGroupCollection collection = resourceGroupResource.GetProximityPlacementGroups();

// invoke the operation
string proximityPlacementGroupName = "myProximityPlacementGroup";
ProximityPlacementGroupData data = new ProximityPlacementGroupData(new AzureLocation("westus"))
{
    Zones =
    {
    "1"
    },
    ProximityPlacementGroupType = ProximityPlacementGroupType.Standard,
    IntentVmSizes =
    {
    "Basic_A0","Basic_A2"
    },
};
ArmOperation<ProximityPlacementGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, proximityPlacementGroupName, data);
ProximityPlacementGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProximityPlacementGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

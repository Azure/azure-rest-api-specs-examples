using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/proximityPlacementGroupExamples/ProximityPlacementGroup_Patch.json
// this example is just showing the usage of "ProximityPlacementGroups_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProximityPlacementGroupResource created on azure
// for more information of creating ProximityPlacementGroupResource, please refer to the document of ProximityPlacementGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string proximityPlacementGroupName = "myProximityPlacementGroup";
ResourceIdentifier proximityPlacementGroupResourceId = ProximityPlacementGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, proximityPlacementGroupName);
ProximityPlacementGroupResource proximityPlacementGroup = client.GetProximityPlacementGroupResource(proximityPlacementGroupResourceId);

// invoke the operation
ProximityPlacementGroupPatch patch = new ProximityPlacementGroupPatch
{
    Tags =
    {
    ["additionalProp1"] = "string"
    },
};
ProximityPlacementGroupResource result = await proximityPlacementGroup.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProximityPlacementGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

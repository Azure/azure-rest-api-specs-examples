using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/TransitHub_CreateOrUpdate.json
// this example is just showing the usage of "TransitHubResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEnclaveCommunityResource created on azure
// for more information of creating VirtualEnclaveCommunityResource, please refer to the document of VirtualEnclaveCommunityResource
string subscriptionId = "CA1CB369-DD26-4DB2-9D43-9AFEF0F22093";
string resourceGroupName = "rgopenapi";
string communityName = "TestMyCommunity";
ResourceIdentifier virtualEnclaveCommunityResourceId = VirtualEnclaveCommunityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, communityName);
VirtualEnclaveCommunityResource virtualEnclaveCommunity = client.GetVirtualEnclaveCommunityResource(virtualEnclaveCommunityResourceId);

// get the collection of this VirtualEnclaveTransitHubResource
VirtualEnclaveTransitHubCollection collection = virtualEnclaveCommunity.GetVirtualEnclaveTransitHubs();

// invoke the operation
string transitHubName = "TestThName";
VirtualEnclaveTransitHubData data = new VirtualEnclaveTransitHubData(new AzureLocation("westcentralus"))
{
    Properties = new VirtualEnclaveTransitHubProperties
    {
        State = TransitHubState.PendingApproval,
        TransitOption = new VirtualEnclaveTransitOptionProperties
        {
            Type = TransitOptionType.ExpressRoute,
            Params = new TransitOptionParams
            {
                ScaleUnits = 1L,
            },
        },
    },
    Tags =
    {
    ["Tag1"] = "Value1"
    },
};
ArmOperation<VirtualEnclaveTransitHubResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, transitHubName, data);
VirtualEnclaveTransitHubResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualEnclaveTransitHubData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

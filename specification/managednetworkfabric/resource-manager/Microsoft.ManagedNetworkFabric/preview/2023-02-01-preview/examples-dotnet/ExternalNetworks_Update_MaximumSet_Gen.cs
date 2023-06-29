using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/ExternalNetworks_Update_MaximumSet_Gen.json
// this example is just showing the usage of "ExternalNetworks_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExternalNetworkResource created on azure
// for more information of creating ExternalNetworkResource, please refer to the document of ExternalNetworkResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string l3IsolationDomainName = "example-l3domain";
string externalNetworkName = "example-externalnetwork";
ResourceIdentifier externalNetworkResourceId = ExternalNetworkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, l3IsolationDomainName, externalNetworkName);
ExternalNetworkResource externalNetwork = client.GetExternalNetworkResource(externalNetworkResourceId);

// invoke the operation
ExternalNetworkPatch patch = new ExternalNetworkPatch()
{
    Annotation = "Lab 1",
    PeeringOption = PeeringOption.OptionA,
    OptionBProperties = new OptionBProperties()
    {
        ImportRouteTargets =
        {
        "65046:10039"
        },
        ExportRouteTargets =
        {
        "65046:10039"
        },
    },
    OptionAProperties = new Layer3OptionAProperties()
    {
        Mtu = 1500,
        VlanId = 1001,
        PeerASN = 65047,
        BfdConfiguration = new BfdConfiguration(),
        PrimaryIPv4Prefix = "10.1.1.0/30",
        PrimaryIPv6Prefix = "3FFE:FFFF:0:CD30::a0/126",
        SecondaryIPv4Prefix = "10.1.1.4/30",
        SecondaryIPv6Prefix = "3FFE:FFFF:0:CD30::a4/126",
    },
    ImportRoutePolicyId = "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
    ExportRoutePolicyId = "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/routePolicies/routePolicyName",
};
ArmOperation<ExternalNetworkResource> lro = await externalNetwork.UpdateAsync(WaitUntil.Completed, patch);
ExternalNetworkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExternalNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

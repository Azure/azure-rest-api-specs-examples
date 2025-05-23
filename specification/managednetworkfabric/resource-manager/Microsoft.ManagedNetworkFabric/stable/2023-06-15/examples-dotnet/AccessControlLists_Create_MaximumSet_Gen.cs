using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/AccessControlLists_Create_MaximumSet_Gen.json
// this example is just showing the usage of "AccessControlLists_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this NetworkFabricAccessControlListResource
NetworkFabricAccessControlListCollection collection = resourceGroupResource.GetNetworkFabricAccessControlLists();

// invoke the operation
string accessControlListName = "example-acl";
NetworkFabricAccessControlListData data = new NetworkFabricAccessControlListData(new AzureLocation("eastUs"))
{
    Annotation = "annotation",
    ConfigurationType = NetworkFabricConfigurationType.File,
    AclsUri = new Uri("https://ACL-Storage-URL"),
    DefaultAction = CommunityActionType.Permit,
    MatchConfigurations = {new AccessControlListMatchConfiguration
    {
    MatchConfigurationName = "example-match",
    SequenceNumber = 123L,
    IPAddressType = NetworkFabricIPAddressType.IPv4,
    MatchConditions = {new AccessControlListMatchCondition
    {
    EtherTypes = {"0x1"},
    Fragments = {"0xff00-0xffff"},
    IPLengths = {"4094-9214"},
    TtlValues = {"23"},
    DscpMarkings = {"32"},
    PortCondition = new AccessControlListPortCondition(Layer4Protocol.Tcp)
    {
    Flags = {"established"},
    PortType = NetworkFabricPortType.SourcePort,
    Ports = {"1-20"},
    PortGroupNames = {"example-portGroup"},
    },
    ProtocolTypes = {"TCP"},
    VlanMatchCondition = new VlanMatchCondition
    {
    Vlans = {"20-30"},
    InnerVlans = {"30"},
    VlanGroupNames = {"example-vlanGroup"},
    },
    IPCondition = new IPMatchCondition
    {
    SourceDestinationType = SourceDestinationType.SourceIP,
    PrefixType = IPMatchConditionPrefixType.Prefix,
    IPPrefixValues = {"10.20.20.20/12"},
    IPGroupNames = {"example-ipGroup"},
    },
    }},
    Actions = {new AccessControlListAction
    {
    AclActionType = AclActionType.Count,
    CounterName = "example-counter",
    }},
    }},
    DynamicMatchConfigurations = {new CommonDynamicMatchConfiguration
    {
    IPGroups = {new MatchConfigurationIPGroupProperties
    {
    Name = "example-ipGroup",
    IPAddressType = NetworkFabricIPAddressType.IPv4,
    IPPrefixes = {"10.20.3.1/20"},
    }},
    VlanGroups = {new VlanGroupProperties
    {
    Name = "example-vlanGroup",
    Vlans = {"20-30"},
    }},
    PortGroups = {new PortGroupProperties
    {
    Name = "example-portGroup",
    Ports = {"100-200"},
    }},
    }},
    Tags =
    {
    ["keyID"] = "KeyValue"
    },
};
ArmOperation<NetworkFabricAccessControlListResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, accessControlListName, data);
NetworkFabricAccessControlListResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricAccessControlListData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

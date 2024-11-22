using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters;

// Generated from example definition: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-09-01-preview/examples/ManagedClusterPutOperation_example_max.json
// this example is just showing the usage of "ManagedClusters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ServiceFabricManagedClusterResource
ServiceFabricManagedClusterCollection collection = resourceGroupResource.GetServiceFabricManagedClusters();

// invoke the operation
string clusterName = "mycluster";
ServiceFabricManagedClusterData data = new ServiceFabricManagedClusterData(new AzureLocation("eastus"), new ServiceFabricManagedClustersSku(ServiceFabricManagedClustersSkuName.Basic))
{
    DnsName = "mycluster",
    ClientConnectionPort = 19000,
    HttpGatewayConnectionPort = 19080,
    AdminUserName = "vmadmin",
    AdminPassword = "{vm-password}",
    LoadBalancingRules =
    {
    new ManagedClusterLoadBalancingRule(80,80,new ManagedClusterLoadBalancingRuleTransportProtocol("http"),ManagedClusterLoadBalanceProbeProtocol.Http)
    {
    ProbePort = 80,
    },new ManagedClusterLoadBalancingRule(443,443,new ManagedClusterLoadBalancingRuleTransportProtocol("http"),ManagedClusterLoadBalanceProbeProtocol.Http)
    {
    ProbePort = 443,
    },new ManagedClusterLoadBalancingRule(10000,10000,ManagedClusterLoadBalancingRuleTransportProtocol.Tcp,ManagedClusterLoadBalanceProbeProtocol.Http)
    {
    ProbePort = 10000,
    LoadDistribution = "Default",
    }
    },
    IsRdpAccessAllowed = true,
    NetworkSecurityRules =
    {
    new ServiceFabricManagedNetworkSecurityRule("TestName",ServiceFabricManagedNsgProtocol.Tcp,ServiceFabricManagedNetworkTrafficAccess.Allow,1010,ServiceFabricManagedNetworkSecurityRuleDirection.Inbound)
    {
    Description = "Test description",
    SourceAddressPrefixes =
    {
    "*"
    },
    DestinationAddressPrefixes =
    {
    "*"
    },
    SourcePortRanges =
    {
    "*"
    },
    DestinationPortRanges =
    {
    "*"
    },
    },new ServiceFabricManagedNetworkSecurityRule("AllowARM",new ServiceFabricManagedNsgProtocol("*"),ServiceFabricManagedNetworkTrafficAccess.Allow,2002,ServiceFabricManagedNetworkSecurityRuleDirection.Inbound)
    {
    SourceAddressPrefix = "AzureResourceManager",
    DestinationAddressPrefix = "*",
    SourcePortRange = "*",
    DestinationPortRange = "33500-33699",
    }
    },
    FabricSettings =
    {
    new ClusterFabricSettingsSection("ManagedIdentityTokenService",new ClusterFabricSettingsParameterDescription[]
    {
    new ClusterFabricSettingsParameterDescription("IsEnabled","true")
    })
    },
    ClusterCodeVersion = "7.1.168.9494",
    ClusterUpgradeMode = ManagedClusterUpgradeMode.Manual,
    AddOnFeatures =
    {
    ManagedClusterAddOnFeature.DnsService,ManagedClusterAddOnFeature.BackupRestoreService,ManagedClusterAddOnFeature.ResourceMonitorService
    },
    IsAutoOSUpgradeEnabled = true,
    HasZoneResiliency = true,
    MaxUnusedVersionsToKeep = 3,
    IsIPv6Enabled = true,
    IPTags =
    {
    new ManagedClusterIPTag("FirstPartyUsage","SQL")
    },
    AuxiliarySubnets =
    {
    new ManagedClusterSubnet("testSubnet1")
    {
    IsIPv6Enabled = true,
    PrivateEndpointNetworkPolicies = ManagedClusterSubnetPrivateEndpointNetworkPoliciesState.Enabled,
    PrivateLinkServiceNetworkPolicies = ManagedClusterSubnetPrivateLinkServiceNetworkPoliciesState.Enabled,
    NetworkSecurityGroupId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/networkSecurityGroups/sn1"),
    }
    },
    ServiceEndpoints =
    {
    new ManagedClusterServiceEndpoint("Microsoft.Storage")
    {
    Locations =
    {
    new AzureLocation("eastus2"),new AzureLocation("usnorth")
    },
    }
    },
    ZonalUpdateMode = ZonalUpdateMode.Fast,
    UseCustomVnet = true,
    PublicIPPrefixId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.Network/publicIPPrefixes/myPublicIPPrefix"),
    PublicIPv6PrefixId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.Network/publicIPPrefixes/myPublicIPv6Prefix"),
    DdosProtectionPlanId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/ddosProtectionPlans/myDDoSProtectionPlan"),
    UpgradeDescription = new ManagedClusterUpgradePolicy()
    {
        ForceRestart = false,
        HealthPolicy = new ManagedClusterHealthPolicy(10, 30),
        DeltaHealthPolicy = new ManagedClusterUpgradeDeltaHealthPolicy(20)
        {
            MaxPercentUpgradeDomainDeltaUnhealthyNodes = 40,
            MaxPercentDeltaUnhealthyApplications = 40,
        },
        MonitoringPolicy = new ManagedClusterMonitoringPolicy(TimeSpan.Parse("00:05:00"), TimeSpan.Parse("00:45:00"), "00:55:00", "12:00:00", "03:00:00"),
    },
    HttpGatewayTokenAuthConnectionPort = 19081,
    IsHttpGatewayExclusiveAuthModeEnabled = true,
    AutoGeneratedDomainNameLabelScope = AutoGeneratedDomainNameLabelScope.SubscriptionReuse,
    AllocatedOutboundPorts = 0,
    Tags =
    {
    },
};
ArmOperation<ServiceFabricManagedClusterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clusterName, data);
ServiceFabricManagedClusterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceFabricManagedClusterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

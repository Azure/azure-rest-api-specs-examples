using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters;

// Generated from example definition: specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/stable/2024-04-01/examples/NodeTypePutOperation_example_max.json
// this example is just showing the usage of "NodeTypes_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricManagedClusterResource created on azure
// for more information of creating ServiceFabricManagedClusterResource, please refer to the document of ServiceFabricManagedClusterResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string clusterName = "myCluster";
ResourceIdentifier serviceFabricManagedClusterResourceId = ServiceFabricManagedClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
ServiceFabricManagedClusterResource serviceFabricManagedCluster = client.GetServiceFabricManagedClusterResource(serviceFabricManagedClusterResourceId);

// get the collection of this ServiceFabricManagedNodeTypeResource
ServiceFabricManagedNodeTypeCollection collection = serviceFabricManagedCluster.GetServiceFabricManagedNodeTypes();

// invoke the operation
string nodeTypeName = "BE-testResourceGroup-testRegion-test";
ServiceFabricManagedNodeTypeData data = new ServiceFabricManagedNodeTypeData()
{
    IsPrimary = false,
    VmInstanceCount = 10,
    DataDiskSizeInGB = 200,
    DataDiskType = ServiceFabricManagedDataDiskType.PremiumLrs,
    DataDiskLetter = "S",
    PlacementProperties =
    {
    ["HasSSD"] = "true",
    ["NodeColor"] = "green",
    ["SomeProperty"] = "5",
    },
    Capacities =
    {
    ["ClientConnections"] = "65536",
    },
    VmSize = "Standard_DS3",
    VmImagePublisher = "MicrosoftWindowsServer",
    VmImageOffer = "WindowsServer",
    VmImageSku = "2016-Datacenter-Server-Core",
    VmImageVersion = "latest",
    VmSecrets =
    {
    new NodeTypeVaultSecretGroup(new WritableSubResource()
    {
    Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.KeyVault/vaults/myVault"),
    },new NodeTypeVaultCertificate[]
    {
    new NodeTypeVaultCertificate(new Uri("https://myVault.vault.azure.net:443/secrets/myCert/ef1a31d39e1f46bca33def54b6cda54c"),"My")
    })
    },
    VmExtensions =
    {
    new NodeTypeVmssExtension("Microsoft.Azure.Geneva.GenevaMonitoring","Microsoft.Azure.Geneva","GenevaMonitoring","2.0")
    {
    AutoUpgradeMinorVersion = true,
    Settings = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
    }),
    ForceUpdateTag = "v.1.0",
    IsAutomaticUpgradeEnabled = true,
    SetupOrder =
    {
    VmssExtensionSetupOrder.BeforeSFRuntime
    },
    }
    },
    UserAssignedIdentities =
    {
    new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity"),new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity2")
    },
    IsStateless = true,
    HasMultiplePlacementGroups = true,
    FrontendConfigurations =
    {
    new NodeTypeFrontendConfiguration()
    {
    LoadBalancerBackendAddressPoolId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/backendAddressPools/LoadBalancerBEAddressPool"),
    LoadBalancerInboundNatPoolId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/inboundNatPools/LoadBalancerNATPool"),
    ApplicationGatewayBackendAddressPoolId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/applicationGateways/appgw-test/backendAddressPools/appgwBepoolTest"),
    }
    },
    AdditionalDataDisks =
    {
    new NodeTypeVmssDataDisk(1,256,ServiceFabricManagedDataDiskType.StandardSsdLrs,"F"),new NodeTypeVmssDataDisk(2,150,ServiceFabricManagedDataDiskType.PremiumLrs,"G")
    },
    IsEncryptionAtHostEnabled = true,
    IsAcceleratedNetworkingEnabled = true,
    UseDefaultPublicLoadBalancer = true,
    IsOverProvisioningEnabled = false,
    IsSpotVm = true,
    UseEphemeralOSDisk = true,
    SpotRestoreTimeout = "PT30M",
    EvictionPolicy = SpotNodeVmEvictionPolicyType.Deallocate,
    SubnetId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
    VmSetupActions =
    {
    VmSetupAction.EnableContainers,VmSetupAction.EnableHyperV
    },
    SecurityType = ServiceFabricManagedClusterSecurityType.TrustedLaunch,
    IsSecureBootEnabled = true,
    IsNodePublicIPEnabled = true,
    IsNodePublicIPv6Enabled = true,
    NatGatewayId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/natGateways/myNatGateway"),
    ServiceArtifactReferenceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Compute/galleries/myGallery/serviceArtifacts/myServiceArtifact/vmArtifactsProfiles/myVmArtifactProfile"),
    DscpConfigurationId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/dscpConfigurations/myDscpConfig"),
    AdditionalNetworkInterfaceConfigurations =
    {
    new AdditionalNetworkInterfaceConfiguration("nic-1",new ServiceFabricManagedClusterIPConfiguration[]
    {
    new ServiceFabricManagedClusterIPConfiguration("ipconfig-1")
    {
    ApplicationGatewayBackendAddressPools =
    {
    new WritableSubResource()
    {
    Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/applicationGateways/appgw-test/backendAddressPools/appgwBepoolTest"),
    }
    },
    LoadBalancerBackendAddressPools =
    {
    new WritableSubResource()
    {
    Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/backendAddressPools/LoadBalancerBEAddressPool"),
    }
    },
    LoadBalancerInboundNatPools =
    {
    new WritableSubResource()
    {
    Id = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/loadBalancers/test-LB/inboundNatPools/LoadBalancerNATPool"),
    }
    },
    SubnetId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
    PrivateIPAddressVersion = ServiceFabricManagedClusterPrivateIPAddressVersion.IPv4,
    PublicIPAddressConfiguration = new ServiceFabricManagedClusterPublicIPAddressConfiguration("publicip-1")
    {
    IPTags =
    {
    new ManagedClusterIPTag("RoutingPreference","Internet")
    },
    PublicIPAddressVersion = ServiceFabricManagedClusterPublicIPAddressVersion.IPv4,
    },
    }
    })
    {
    EnableAcceleratedNetworking = true,
    DscpConfigurationId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resRg/providers/Microsoft.Network/dscpConfigurations/myDscpConfig"),
    }
    },
    ComputerNamePrefix = "BE",
};
ArmOperation<ServiceFabricManagedNodeTypeResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, nodeTypeName, data);
ServiceFabricManagedNodeTypeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceFabricManagedNodeTypeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

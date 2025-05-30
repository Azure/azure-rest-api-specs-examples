using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Hci.Models;
using Azure.ResourceManager.Hci;

// Generated from example definition: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PutDeploymentSettings.json
// this example is just showing the usage of "DeploymentSettings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HciClusterResource created on azure
// for more information of creating HciClusterResource, please refer to the document of HciClusterResource
string subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
string resourceGroupName = "test-rg";
string clusterName = "myCluster";
ResourceIdentifier hciClusterResourceId = HciClusterResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName);
HciClusterResource hciCluster = client.GetHciClusterResource(hciClusterResourceId);

// get the collection of this HciClusterDeploymentSettingResource
HciClusterDeploymentSettingCollection collection = hciCluster.GetHciClusterDeploymentSettings();

// invoke the operation
string deploymentSettingsName = "default";
HciClusterDeploymentSettingData data = new HciClusterDeploymentSettingData
{
    ArcNodeResourceIds = { new ResourceIdentifier("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1"), new ResourceIdentifier("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-2") },
    DeploymentMode = EceDeploymentMode.Deploy,
    OperationType = HciClusterOperationType.ClusterProvisioning,
    DeploymentConfiguration = new HciClusterDeploymentConfiguration(new DeploymentSettingScaleUnits[]
{
new DeploymentSettingScaleUnits(new HciClusterDeploymentInfo
{
SecuritySettings = new HciClusterDeploymentSecuritySettings
{
IsHvciProtectionEnabled = true,
IsDrtmProtectionEnabled = true,
IsDriftControlEnforced = true,
IsCredentialGuardEnforced = false,
IsSmbSigningEnforced = true,
IsSmbClusterEncryptionEnabled = false,
IsSideChannelMitigationEnforced = true,
IsBitlockerBootVolumeEnabled = true,
AreBitlockerDataVolumesEnabled = true,
IsWdacEnforced = true,
},
Observability = new DeploymentSettingObservability
{
IsStreamingDataClientEnabled = true,
IsEULocation = false,
IsEpisodicDataUploadEnabled = true,
},
Cluster = new HciDeploymentCluster
{
Name = "testHCICluster",
WitnessType = "Cloud",
WitnessPath = "Cloud",
CloudAccountName = "myasestoragacct",
AzureServiceEndpoint = "core.windows.net",
},
StorageConfigurationMode = "Express",
NamingPrefix = "ms169",
DomainFqdn = "ASZ1PLab8.nttest.microsoft.com",
InfrastructureNetwork = {new DeploymentSettingInfrastructureNetwork
{
SubnetMask = "255.255.248.0",
Gateway = "255.255.248.0",
IPPools = {new DeploymentSettingIPPools
{
StartingAddress = "10.57.48.60",
EndingAddress = "10.57.48.66",
}},
DnsServers = {"10.57.50.90"},
}},
PhysicalNodes = {new DeploymentSettingPhysicalNodes
{
Name = "ms169host",
IPv4Address = "10.57.51.224",
}, new DeploymentSettingPhysicalNodes
{
Name = "ms154host",
IPv4Address = "10.57.53.236",
}},
HostNetwork = new DeploymentSettingHostNetwork
{
Intents = {new DeploymentSettingIntents
{
Name = "Compute_Management",
TrafficType = {"Compute", "Management"},
Adapter = {"Port2"},
OverrideVirtualSwitchConfiguration = false,
VirtualSwitchConfigurationOverrides = new DeploymentSettingVirtualSwitchConfigurationOverrides
{
EnableIov = "True",
LoadBalancingAlgorithm = "HyperVPort",
},
OverrideQosPolicy = false,
QosPolicyOverrides = new DeploymentSettingQosPolicyOverrides
{
PriorityValue8021ActionCluster = "7",
PriorityValue8021ActionSmb = "3",
BandwidthPercentageSmb = "50",
},
OverrideAdapterProperty = false,
AdapterPropertyOverrides = new DeploymentSettingAdapterPropertyOverrides
{
JumboPacket = "1514",
NetworkDirect = "Enabled",
NetworkDirectTechnology = "iWARP",
},
}},
StorageNetworks = {new DeploymentSettingStorageNetworks
{
Name = "Storage1Network",
NetworkAdapterName = "Port3",
VlanId = "5",
StorageAdapterIPInfo = {new DeploymentSettingStorageAdapterIPInfo
{
PhysicalNode = "string",
IPv4Address = "10.57.48.60",
SubnetMask = "255.255.248.0",
}},
}},
StorageConnectivitySwitchless = true,
EnableStorageAutoIP = false,
},
SdnIntegrationNetworkController = new DeploymentSettingNetworkController
{
MacAddressPoolStart = "00-0D-3A-1B-C7-21",
MacAddressPoolStop = "00-0D-3A-1B-C7-29",
NetworkVirtualizationEnabled = true,
},
AdouPath = "OU=ms169,DC=ASZ1PLab8,DC=nttest,DC=microsoft,DC=com",
SecretsLocation = "/subscriptions/db4e2fdb-6d80-4e6e-b7cd-xxxxxxx/resourceGroups/test-rg/providers/Microsoft.KeyVault/vaults/abcd123",
Secrets = {new EceDeploymentSecrets
{
SecretName = "cluster1-BmcAdminUser-f5bcc1d9-23af-4ae9-aca1-041d0f593a63",
EceSecretName = new EceSecret("BMCAdminUserCred"),
SecretLocation = new Uri("https://sclusterkvnirhci35.vault.azure.net/secrets/cluster-34232342-BmcAdminUser-f5bcc1d9-23af-4ae9-aca1-041d0f593a63/9276354aabfc492fa9b2cdbefb54ae4b"),
}, new EceDeploymentSecrets
{
SecretName = "cluster2-AzureStackLCMUserCredential-f5bcc1d9-23af-4ae9-aca1-041d0f593a63",
EceSecretName = EceSecret.AzureStackLcmUserCredential,
SecretLocation = new Uri("https://sclusterkvnirhci35.vault.azure.net/secrets/cluster-34232342-AzureStackLCMUserCredential-f5bcc1d9-23af-4ae9-aca1-041d0f593a63/9276354aabfc492fa9b2cdbefb54ae4c"),
}},
OptionalServicesCustomLocation = "customLocationName",
})
{
SbePartnerInfo = new SbePartnerInfo
{
SbeDeploymentInfo = new SbeDeploymentInfo
{
Version = "4.0.2309.13",
Family = "Gen5",
Publisher = "Contoso",
SbeManifestSource = "default",
SbeManifestCreationOn = DateTimeOffset.Parse("2023-07-25T02:40:33Z"),
},
PartnerProperties = {new SbePartnerProperties
{
Name = "EnableBMCIpV6",
Value = "false",
}, new SbePartnerProperties
{
Name = "PhoneHomePort",
Value = "1653",
}, new SbePartnerProperties
{
Name = "BMCSecurityState",
Value = "HighSecurity",
}},
CredentialList = {new SbeCredentials
{
SecretName = "cluster1-DownloadConnectorCred-f5bcc1d9-23af-4ae9-aca1-041d0f593a63",
EceSecretName = "DownloadConnectorCred",
SecretLocation = new Uri("https://sclusterkvnirhci35.vault.azure.net/secrets/cluster-34232342-DownloadConnectorCred-f5bcc1d9-23af-4ae9-aca1-041d0f593a63/9276354aabfc492fa9b2cdbefb54ae4b"),
}},
},
}
})
    {
        Version = "string",
    },
};
ArmOperation<HciClusterDeploymentSettingResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, deploymentSettingsName, data);
HciClusterDeploymentSettingResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HciClusterDeploymentSettingData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");

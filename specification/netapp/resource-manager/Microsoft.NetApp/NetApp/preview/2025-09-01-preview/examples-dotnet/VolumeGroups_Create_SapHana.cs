using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-09-01-preview/examples/VolumeGroups_Create_SapHana.json
// this example is just showing the usage of "VolumeGroups_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppVolumeGroupResource created on azure
// for more information of creating NetAppVolumeGroupResource, please refer to the document of NetAppVolumeGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string volumeGroupName = "group1";
ResourceIdentifier netAppVolumeGroupResourceId = NetAppVolumeGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, volumeGroupName);
NetAppVolumeGroupResource netAppVolumeGroup = client.GetNetAppVolumeGroupResource(netAppVolumeGroupResourceId);

// invoke the operation
NetAppVolumeGroupData data = new NetAppVolumeGroupData
{
    Location = new AzureLocation("westus"),
    GroupMetaData = new NetAppVolumeGroupMetadata
    {
        GroupDescription = "Volume group",
        ApplicationType = NetAppApplicationType.SapHana,
        ApplicationIdentifier = "SH9",
    },
    Volumes = {new NetAppVolumeGroupVolume("test-data-mnt00001", 107374182400L, new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"))
    {
    Name = "test-data-mnt00001",
    ServiceLevel = NetAppFileServiceLevel.Premium,
    ExportRules = {new NetAppVolumeExportPolicyRule
    {
    RuleIndex = 1,
    IsUnixReadOnly = true,
    IsUnixReadWrite = true,
    IsKerberos5ReadOnly = false,
    IsKerberos5ReadWrite = false,
    IsKerberos5iReadOnly = false,
    IsKerberos5iReadWrite = false,
    IsKerberos5pReadOnly = false,
    IsKerberos5pReadWrite = false,
    AllowCifsProtocol = false,
    AllowNfsV3Protocol = false,
    AllowNfsV41Protocol = true,
    AllowedClients = "0.0.0.0/0",
    HasRootAccess = true,
    }},
    ProtocolTypes = {"NFSv4.1"},
    ThroughputMibps = 10,
    CapacityPoolResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
    ProximityPlacementGroupId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
    VolumeSpecName = "data",
    }, new NetAppVolumeGroupVolume("test-log-mnt00001", 107374182400L, new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"))
    {
    Name = "test-log-mnt00001",
    ServiceLevel = NetAppFileServiceLevel.Premium,
    ExportRules = {new NetAppVolumeExportPolicyRule
    {
    RuleIndex = 1,
    IsUnixReadOnly = true,
    IsUnixReadWrite = true,
    IsKerberos5ReadOnly = false,
    IsKerberos5ReadWrite = false,
    IsKerberos5iReadOnly = false,
    IsKerberos5iReadWrite = false,
    IsKerberos5pReadOnly = false,
    IsKerberos5pReadWrite = false,
    AllowCifsProtocol = false,
    AllowNfsV3Protocol = false,
    AllowNfsV41Protocol = true,
    AllowedClients = "0.0.0.0/0",
    HasRootAccess = true,
    }},
    ProtocolTypes = {"NFSv4.1"},
    ThroughputMibps = 10,
    CapacityPoolResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
    ProximityPlacementGroupId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
    VolumeSpecName = "log",
    }, new NetAppVolumeGroupVolume("test-shared", 107374182400L, new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"))
    {
    Name = "test-shared",
    ServiceLevel = NetAppFileServiceLevel.Premium,
    ExportRules = {new NetAppVolumeExportPolicyRule
    {
    RuleIndex = 1,
    IsUnixReadOnly = true,
    IsUnixReadWrite = true,
    IsKerberos5ReadOnly = false,
    IsKerberos5ReadWrite = false,
    IsKerberos5iReadOnly = false,
    IsKerberos5iReadWrite = false,
    IsKerberos5pReadOnly = false,
    IsKerberos5pReadWrite = false,
    AllowCifsProtocol = false,
    AllowNfsV3Protocol = false,
    AllowNfsV41Protocol = true,
    AllowedClients = "0.0.0.0/0",
    HasRootAccess = true,
    }},
    ProtocolTypes = {"NFSv4.1"},
    ThroughputMibps = 10,
    CapacityPoolResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
    ProximityPlacementGroupId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
    VolumeSpecName = "shared",
    }, new NetAppVolumeGroupVolume("test-data-backup", 107374182400L, new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"))
    {
    Name = "test-data-backup",
    ServiceLevel = NetAppFileServiceLevel.Premium,
    ExportRules = {new NetAppVolumeExportPolicyRule
    {
    RuleIndex = 1,
    IsUnixReadOnly = true,
    IsUnixReadWrite = true,
    IsKerberos5ReadOnly = false,
    IsKerberos5ReadWrite = false,
    IsKerberos5iReadOnly = false,
    IsKerberos5iReadWrite = false,
    IsKerberos5pReadOnly = false,
    IsKerberos5pReadWrite = false,
    AllowCifsProtocol = false,
    AllowNfsV3Protocol = false,
    AllowNfsV41Protocol = true,
    AllowedClients = "0.0.0.0/0",
    HasRootAccess = true,
    }},
    ProtocolTypes = {"NFSv4.1"},
    ThroughputMibps = 10,
    CapacityPoolResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
    ProximityPlacementGroupId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
    VolumeSpecName = "data-backup",
    }, new NetAppVolumeGroupVolume("test-log-backup", 107374182400L, new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"))
    {
    Name = "test-log-backup",
    ServiceLevel = NetAppFileServiceLevel.Premium,
    ExportRules = {new NetAppVolumeExportPolicyRule
    {
    RuleIndex = 1,
    IsUnixReadOnly = true,
    IsUnixReadWrite = true,
    IsKerberos5ReadOnly = false,
    IsKerberos5ReadWrite = false,
    IsKerberos5iReadOnly = false,
    IsKerberos5iReadWrite = false,
    IsKerberos5pReadOnly = false,
    IsKerberos5pReadWrite = false,
    AllowCifsProtocol = false,
    AllowNfsV3Protocol = false,
    AllowNfsV41Protocol = true,
    AllowedClients = "0.0.0.0/0",
    HasRootAccess = true,
    }},
    ProtocolTypes = {"NFSv4.1"},
    ThroughputMibps = 10,
    CapacityPoolResourceId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1"),
    ProximityPlacementGroupId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg"),
    VolumeSpecName = "log-backup",
    }},
};
ArmOperation<NetAppVolumeGroupResource> lro = await netAppVolumeGroup.UpdateAsync(WaitUntil.Completed, data);
NetAppVolumeGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppVolumeGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");


import com.azure.resourcemanager.netapp.models.ApplicationType;
import com.azure.resourcemanager.netapp.models.ExportPolicyRule;
import com.azure.resourcemanager.netapp.models.ServiceLevel;
import com.azure.resourcemanager.netapp.models.VolumeGroupMetadata;
import com.azure.resourcemanager.netapp.models.VolumeGroupVolumeProperties;
import com.azure.resourcemanager.netapp.models.VolumePropertiesExportPolicy;
import java.util.Arrays;

/**
 * Samples for VolumeGroups Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/VolumeGroups_Create_SapHana.json
     */
    /**
     * Sample code: VolumeGroups_Create_SapHana.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsCreateSapHana(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .volumeGroups().define("group1").withExistingNetAppAccount("myRG", "account1").withRegion(
                "westus")
            .withGroupMetadata(
                new VolumeGroupMetadata()
                    .withGroupDescription("Volume group").withApplicationType(
                        ApplicationType.SAP_HANA)
                    .withApplicationIdentifier("SH9"))
            .withVolumes(Arrays.asList(new VolumeGroupVolumeProperties().withName("test-data-mnt00001")
                .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                .withUsageThreshold(107374182400L)
                .withExportPolicy(new VolumePropertiesExportPolicy().withRules(Arrays.asList(new ExportPolicyRule()
                    .withRuleIndex(1).withUnixReadOnly(true).withUnixReadWrite(true).withKerberos5ReadOnly(false)
                    .withKerberos5ReadWrite(false).withKerberos5IReadOnly(false).withKerberos5IReadWrite(false)
                    .withKerberos5PReadOnly(false).withKerberos5PReadWrite(false).withCifs(false).withNfsv3(false)
                    .withNfsv41(true).withAllowedClients("0.0.0.0/0").withHasRootAccess(true))))
                .withProtocolTypes(Arrays.asList("NFSv4.1"))
                .withSubnetId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                .withThroughputMibps(10.0F)
                .withCapacityPoolResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                .withProximityPlacementGroup(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg")
                .withVolumeSpecName("data"),
                new VolumeGroupVolumeProperties().withName("test-log-mnt00001")
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L)
                    .withExportPolicy(new VolumePropertiesExportPolicy().withRules(Arrays.asList(new ExportPolicyRule()
                        .withRuleIndex(1).withUnixReadOnly(true).withUnixReadWrite(true).withKerberos5ReadOnly(false)
                        .withKerberos5ReadWrite(false).withKerberos5IReadOnly(false).withKerberos5IReadWrite(false)
                        .withKerberos5PReadOnly(false).withKerberos5PReadWrite(false).withCifs(false).withNfsv3(false)
                        .withNfsv41(true).withAllowedClients("0.0.0.0/0").withHasRootAccess(true))))
                    .withProtocolTypes(Arrays.asList("NFSv4.1"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withThroughputMibps(10.0F)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withProximityPlacementGroup(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg")
                    .withVolumeSpecName("log"),
                new VolumeGroupVolumeProperties().withName("test-shared").withCreationToken("fakeTokenPlaceholder")
                    .withServiceLevel(ServiceLevel.PREMIUM).withUsageThreshold(107374182400L)
                    .withExportPolicy(new VolumePropertiesExportPolicy().withRules(Arrays.asList(new ExportPolicyRule()
                        .withRuleIndex(1).withUnixReadOnly(true).withUnixReadWrite(true).withKerberos5ReadOnly(false)
                        .withKerberos5ReadWrite(false).withKerberos5IReadOnly(false).withKerberos5IReadWrite(false)
                        .withKerberos5PReadOnly(false).withKerberos5PReadWrite(false).withCifs(false).withNfsv3(false)
                        .withNfsv41(true).withAllowedClients("0.0.0.0/0").withHasRootAccess(true))))
                    .withProtocolTypes(Arrays.asList("NFSv4.1"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withThroughputMibps(10.0F)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withProximityPlacementGroup(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg")
                    .withVolumeSpecName("shared"),
                new VolumeGroupVolumeProperties().withName("test-data-backup").withCreationToken("fakeTokenPlaceholder")
                    .withServiceLevel(ServiceLevel.PREMIUM).withUsageThreshold(107374182400L)
                    .withExportPolicy(new VolumePropertiesExportPolicy().withRules(Arrays.asList(new ExportPolicyRule()
                        .withRuleIndex(1).withUnixReadOnly(true).withUnixReadWrite(true).withKerberos5ReadOnly(false)
                        .withKerberos5ReadWrite(false).withKerberos5IReadOnly(false).withKerberos5IReadWrite(false)
                        .withKerberos5PReadOnly(false).withKerberos5PReadWrite(false).withCifs(false).withNfsv3(false)
                        .withNfsv41(true).withAllowedClients("0.0.0.0/0").withHasRootAccess(true))))
                    .withProtocolTypes(Arrays.asList("NFSv4.1"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withThroughputMibps(10.0F)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withProximityPlacementGroup(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg")
                    .withVolumeSpecName("data-backup"),
                new VolumeGroupVolumeProperties().withName("test-log-backup").withCreationToken("fakeTokenPlaceholder")
                    .withServiceLevel(ServiceLevel.PREMIUM).withUsageThreshold(107374182400L)
                    .withExportPolicy(new VolumePropertiesExportPolicy().withRules(Arrays.asList(new ExportPolicyRule()
                        .withRuleIndex(1).withUnixReadOnly(true).withUnixReadWrite(true).withKerberos5ReadOnly(false)
                        .withKerberos5ReadWrite(false).withKerberos5IReadOnly(false).withKerberos5IReadWrite(false)
                        .withKerberos5PReadOnly(false).withKerberos5PReadWrite(false).withCifs(false).withNfsv3(false)
                        .withNfsv41(true).withAllowedClients("0.0.0.0/0").withHasRootAccess(true))))
                    .withProtocolTypes(Arrays.asList("NFSv4.1"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withThroughputMibps(10.0F)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withProximityPlacementGroup(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/cys_sjain_fcp_rg/providers/Microsoft.Compute/proximityPlacementGroups/svlqa_sjain_multivolume_ppg")
                    .withVolumeSpecName("log-backup")))
            .create();
    }
}

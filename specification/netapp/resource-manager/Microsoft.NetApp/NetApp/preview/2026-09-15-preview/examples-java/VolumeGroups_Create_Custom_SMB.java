
import com.azure.resourcemanager.netapp.models.ApplicationType;
import com.azure.resourcemanager.netapp.models.AvsDataStore;
import com.azure.resourcemanager.netapp.models.ServiceLevel;
import com.azure.resourcemanager.netapp.models.SmbAccessBasedEnumeration;
import com.azure.resourcemanager.netapp.models.SmbNonBrowsable;
import com.azure.resourcemanager.netapp.models.VolumeGroupMetadata;
import com.azure.resourcemanager.netapp.models.VolumeGroupVolumeProperties;
import com.azure.resourcemanager.netapp.models.WorkloadApplicationType;
import java.util.Arrays;

/**
 * Samples for VolumeGroups Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/VolumeGroups_Create_Custom_SMB.json
     */
    /**
     * Sample code: VolumeGroups_Create_Custom_SMB.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeGroupsCreateCustomSMB(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeGroups().define("group1").withExistingNetAppAccount("myRG", "account1").withRegion("westus")
            .withGroupMetadata(new VolumeGroupMetadata().withGroupDescription("Volume group")
                .withApplicationType(ApplicationType.CUSTOM).withApplicationIdentifier("CU2"))
            .withVolumes(Arrays.asList(new VolumeGroupVolumeProperties().withName("test-cus-data1")
                .withZones(Arrays.asList("1")).withCreationToken("fakeTokenPlaceholder")
                .withServiceLevel(ServiceLevel.PREMIUM).withUsageThreshold(107374182400L)
                .withProtocolTypes(Arrays.asList("CIFS"))
                .withSubnetId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                .withCapacityPoolResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                .withVolumeSpecName("data1").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX),
                new VolumeGroupVolumeProperties().withName("test-cus-data2").withZones(Arrays.asList("1"))
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L).withProtocolTypes(Arrays.asList("CIFS"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                    .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withVolumeSpecName("data2").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX),
                new VolumeGroupVolumeProperties().withName("test-cus-data3").withZones(Arrays.asList("1"))
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L).withProtocolTypes(Arrays.asList("CIFS"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                    .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withVolumeSpecName("data3").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX),
                new VolumeGroupVolumeProperties().withName("test-cus-data4").withZones(Arrays.asList("1"))
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L).withProtocolTypes(Arrays.asList("CIFS"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                    .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withVolumeSpecName("data4").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX),
                new VolumeGroupVolumeProperties().withName("test-cus-data5").withZones(Arrays.asList("1"))
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L).withProtocolTypes(Arrays.asList("CIFS"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                    .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withVolumeSpecName("data5").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX),
                new VolumeGroupVolumeProperties().withName("test-cus-data6").withZones(Arrays.asList("1"))
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L).withProtocolTypes(Arrays.asList("CIFS"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                    .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withVolumeSpecName("data6").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX),
                new VolumeGroupVolumeProperties().withName("test-cus-data7").withZones(Arrays.asList("1"))
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L).withProtocolTypes(Arrays.asList("CIFS"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                    .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withVolumeSpecName("data7").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX),
                new VolumeGroupVolumeProperties().withName("test-cus-data8").withZones(Arrays.asList("1"))
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L).withProtocolTypes(Arrays.asList("CIFS"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                    .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withVolumeSpecName("data8").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX),
                new VolumeGroupVolumeProperties().withName("test-cus-data9").withZones(Arrays.asList("1"))
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L).withProtocolTypes(Arrays.asList("CIFS"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                    .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withVolumeSpecName("data9").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX),
                new VolumeGroupVolumeProperties().withName("test-cus-data10").withZones(Arrays.asList("1"))
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L).withProtocolTypes(Arrays.asList("CIFS"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                    .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withVolumeSpecName("data10").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX),
                new VolumeGroupVolumeProperties().withName("test-cus-data11").withZones(Arrays.asList("1"))
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L).withProtocolTypes(Arrays.asList("CIFS"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                    .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withVolumeSpecName("data11").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX),
                new VolumeGroupVolumeProperties().withName("test-cus-data12").withZones(Arrays.asList("1"))
                    .withCreationToken("fakeTokenPlaceholder").withServiceLevel(ServiceLevel.PREMIUM)
                    .withUsageThreshold(107374182400L).withProtocolTypes(Arrays.asList("CIFS"))
                    .withSubnetId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
                    .withSmbEncryption(false).withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.DISABLED)
                    .withSmbNonBrowsable(SmbNonBrowsable.DISABLED).withSmbContinuouslyAvailable(false)
                    .withThroughputMibps(10.0F).withAvsDataStore(AvsDataStore.DISABLED)
                    .withCapacityPoolResourceId(
                        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRG/providers/Microsoft.NetApp/netAppAccounts/account1/capacityPools/pool1")
                    .withVolumeSpecName("data12").withWorkloadApplicationType(WorkloadApplicationType.AVDFSLOGIX)))
            .create();
    }
}


import com.azure.resourcemanager.netapp.models.ServiceLevel;
import com.azure.resourcemanager.netapp.models.SmbAccessBasedEnumeration;
import com.azure.resourcemanager.netapp.models.SmbNonBrowsable;
import com.azure.resourcemanager.netapp.models.SmbOplocks;
import java.util.Arrays;

/**
 * Samples for Volumes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/Volumes_CreateOrUpdate_SMB.json
     */
    /**
     * Sample code: Volumes_CreateOrUpdate_SMB.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesCreateOrUpdateSMB(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().define("volume1").withRegion("eastus").withExistingCapacityPool("myRG", "account1", "pool1")
            .withCreationToken("my-smb-volume").withUsageThreshold(107374182400L)
            .withSubnetId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
            .withServiceLevel(ServiceLevel.PREMIUM).withProtocolTypes(Arrays.asList("CIFS")).withSmbEncryption(true)
            .withSmbAccessBasedEnumeration(SmbAccessBasedEnumeration.ENABLED)
            .withSmbNonBrowsable(SmbNonBrowsable.ENABLED).withSmbContinuouslyAvailable(true)
            .withSmbOplocks(SmbOplocks.ENABLED).create();
    }
}

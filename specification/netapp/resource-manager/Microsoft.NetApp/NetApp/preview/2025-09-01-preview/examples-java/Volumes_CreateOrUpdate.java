
import com.azure.resourcemanager.netapp.models.ServiceLevel;

/**
 * Samples for Volumes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Volumes_CreateOrUpdate.json
     */
    /**
     * Sample code: Volumes_CreateOrUpdate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().define("volume1").withRegion("eastus").withExistingCapacityPool("myRG", "account1", "pool1")
            .withCreationToken("my-unique-file-path").withUsageThreshold(107374182400L)
            .withSubnetId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3")
            .withServiceLevel(ServiceLevel.PREMIUM).create();
    }
}

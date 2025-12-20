
import com.azure.resourcemanager.avs.models.ManagementCluster;
import com.azure.resourcemanager.avs.models.PrivateCloud;

/**
 * Samples for PrivateClouds Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PrivateClouds_Update_Stretched.json
     */
    /**
     * Sample code: PrivateClouds_Update_Stretched.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsUpdateStretched(com.azure.resourcemanager.avs.AvsManager manager) {
        PrivateCloud resource = manager.privateClouds()
            .getByResourceGroupWithResponse("group1", "cloud1", com.azure.core.util.Context.NONE).getValue();
        resource.update().withManagementCluster(new ManagementCluster().withClusterSize(4)).apply();
    }
}

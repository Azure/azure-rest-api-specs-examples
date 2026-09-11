
/**
 * Samples for PrivateClouds GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PrivateClouds_Get.json
     */
    /**
     * Sample code: PrivateClouds_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().getByResourceGroupWithResponse("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}

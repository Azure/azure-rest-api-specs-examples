
/**
 * Samples for PrivateClouds Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PrivateClouds_Delete.json
     */
    /**
     * Sample code: PrivateClouds_Delete.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsDelete(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().delete("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}

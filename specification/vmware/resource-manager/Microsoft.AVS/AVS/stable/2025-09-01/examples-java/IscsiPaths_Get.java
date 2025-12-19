
/**
 * Samples for IscsiPaths Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/IscsiPaths_Get.json
     */
    /**
     * Sample code: IscsiPaths_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void iscsiPathsGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.iscsiPaths().getWithResponse("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}

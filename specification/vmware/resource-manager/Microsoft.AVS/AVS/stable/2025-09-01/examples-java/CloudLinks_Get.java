
/**
 * Samples for CloudLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/CloudLinks_Get.json
     */
    /**
     * Sample code: CloudLinks_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void cloudLinksGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.cloudLinks().getWithResponse("group1", "cloud1", "cloudLink1", com.azure.core.util.Context.NONE);
    }
}

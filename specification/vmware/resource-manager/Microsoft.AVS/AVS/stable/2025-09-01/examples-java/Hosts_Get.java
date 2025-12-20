
/**
 * Samples for Hosts Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Hosts_Get.json
     */
    /**
     * Sample code: Hosts_Get.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void hostsGet(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.hosts().getWithResponse("group1", "cloud1", "cluster1",
            "esx03-r52.1111111111111111111.westcentralus.prod.azure.com", com.azure.core.util.Context.NONE);
    }
}

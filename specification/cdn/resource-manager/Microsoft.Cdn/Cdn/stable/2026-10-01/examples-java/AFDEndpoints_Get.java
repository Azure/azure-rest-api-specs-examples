
/**
 * Samples for AfdEndpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDEndpoints_Get.json
     */
    /**
     * Sample code: AFDEndpoints_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDEndpointsGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdEndpoints().getWithResponse("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}

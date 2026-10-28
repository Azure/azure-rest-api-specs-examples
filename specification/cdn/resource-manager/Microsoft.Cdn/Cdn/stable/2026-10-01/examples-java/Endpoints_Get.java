
/**
 * Samples for Endpoints Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Endpoints_Get.json
     */
    /**
     * Sample code: Endpoints_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void endpointsGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getEndpoints().getWithResponse("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}

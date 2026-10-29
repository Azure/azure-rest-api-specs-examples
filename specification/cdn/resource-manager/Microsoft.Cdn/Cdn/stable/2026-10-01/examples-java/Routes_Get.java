
/**
 * Samples for Routes Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Routes_Get.json
     */
    /**
     * Sample code: Routes_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void routesGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRoutes().getWithResponse("RG", "profile1", "endpoint1", "route1",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for Origins Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Origins_Get.json
     */
    /**
     * Sample code: Origins_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void originsGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getOrigins().getWithResponse("RG", "profile1", "endpoint1", "www-someDomain-net",
            com.azure.core.util.Context.NONE);
    }
}

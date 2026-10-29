
/**
 * Samples for AfdOrigins Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDOrigins_Get.json
     */
    /**
     * Sample code: AFDOrigins_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDOriginsGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdOrigins().getWithResponse("RG", "profile1", "origingroup1", "origin1",
            com.azure.core.util.Context.NONE);
    }
}

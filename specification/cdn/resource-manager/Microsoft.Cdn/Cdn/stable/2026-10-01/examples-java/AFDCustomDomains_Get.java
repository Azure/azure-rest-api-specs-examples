
/**
 * Samples for AfdCustomDomains Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDCustomDomains_Get.json
     */
    /**
     * Sample code: AFDCustomDomains_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDCustomDomainsGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdCustomDomains().getWithResponse("RG", "profile1", "domain1",
            com.azure.core.util.Context.NONE);
    }
}

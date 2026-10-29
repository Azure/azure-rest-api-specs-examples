
/**
 * Samples for CustomDomains Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/CustomDomains_Get.json
     */
    /**
     * Sample code: CustomDomains_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void customDomainsGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getCustomDomains().getWithResponse("RG", "profile1", "endpoint1", "www-someDomain-net",
            com.azure.core.util.Context.NONE);
    }
}

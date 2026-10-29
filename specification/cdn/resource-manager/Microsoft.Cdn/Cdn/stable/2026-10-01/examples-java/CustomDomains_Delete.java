
/**
 * Samples for CustomDomains Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/CustomDomains_Delete.json
     */
    /**
     * Sample code: CustomDomains_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void customDomainsDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getCustomDomains().delete("RG", "profile1", "endpoint1", "www-someDomain-net",
            com.azure.core.util.Context.NONE);
    }
}

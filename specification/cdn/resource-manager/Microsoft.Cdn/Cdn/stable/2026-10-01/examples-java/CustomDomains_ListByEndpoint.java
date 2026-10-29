
/**
 * Samples for CustomDomains ListByEndpoint.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/CustomDomains_ListByEndpoint.json
     */
    /**
     * Sample code: CustomDomains_ListByEndpoint.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void customDomainsListByEndpoint(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getCustomDomains().listByEndpoint("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for AfdCustomDomains Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDCustomDomains_Delete.json
     */
    /**
     * Sample code: AFDCustomDomains_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDCustomDomainsDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdCustomDomains().delete("RG", "profile1", "domain1",
            com.azure.core.util.Context.NONE);
    }
}

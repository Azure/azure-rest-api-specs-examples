
/**
 * Samples for AfdCustomDomains RefreshValidationToken.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDCustomDomains_RefreshValidationToken.json
     */
    /**
     * Sample code: AFDCustomDomains_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDCustomDomainsDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdCustomDomains().refreshValidationToken("RG", "profile1", "domain1",
            com.azure.core.util.Context.NONE);
    }
}

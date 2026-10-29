
/**
 * Samples for AfdCustomDomains ListByProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDCustomDomains_ListByProfile.json
     */
    /**
     * Sample code: AFDCustomDomains_ListByProfile.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDCustomDomainsListByProfile(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdCustomDomains().listByProfile("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}

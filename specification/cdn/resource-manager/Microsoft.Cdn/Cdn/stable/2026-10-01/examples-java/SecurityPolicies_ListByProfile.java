
/**
 * Samples for SecurityPolicies ListByProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/SecurityPolicies_ListByProfile.json
     */
    /**
     * Sample code: SecurityPolicies_ListByProfile.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void securityPoliciesListByProfile(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getSecurityPolicies().listByProfile("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}

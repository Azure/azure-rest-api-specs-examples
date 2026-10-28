
/**
 * Samples for SecurityPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/SecurityPolicies_Delete.json
     */
    /**
     * Sample code: SecurityPolicies_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void securityPoliciesDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getSecurityPolicies().delete("RG", "profile1", "securityPolicy1",
            com.azure.core.util.Context.NONE);
    }
}

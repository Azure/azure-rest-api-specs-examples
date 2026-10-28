
/**
 * Samples for SecurityPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/SecurityPolicies_Get.json
     */
    /**
     * Sample code: SecurityPolicies_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void securityPoliciesGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getSecurityPolicies().getWithResponse("RG", "profile1", "securityPolicy1",
            com.azure.core.util.Context.NONE);
    }
}

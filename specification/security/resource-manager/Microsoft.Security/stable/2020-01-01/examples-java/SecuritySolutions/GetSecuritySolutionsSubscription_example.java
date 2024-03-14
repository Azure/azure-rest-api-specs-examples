
/**
 * Samples for SecuritySolutions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/SecuritySolutions/
     * GetSecuritySolutionsSubscription_example.json
     */
    /**
     * Sample code: Get security solutions.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecuritySolutions(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securitySolutions().list(com.azure.core.util.Context.NONE);
    }
}

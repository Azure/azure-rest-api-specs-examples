
/**
 * Samples for DiscoveredSecuritySolutions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/DiscoveredSecuritySolutions
     * /GetDiscoveredSecuritySolutionsSubscription_example.json
     */
    /**
     * Sample code: Get discovered security solutions.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getDiscoveredSecuritySolutions(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.discoveredSecuritySolutions().list(com.azure.core.util.Context.NONE);
    }
}

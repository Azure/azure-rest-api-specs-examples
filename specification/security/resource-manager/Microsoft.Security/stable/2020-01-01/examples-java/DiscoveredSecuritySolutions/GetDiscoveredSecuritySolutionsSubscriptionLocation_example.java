
/**
 * Samples for DiscoveredSecuritySolutions ListByHomeRegion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/DiscoveredSecuritySolutions
     * /GetDiscoveredSecuritySolutionsSubscriptionLocation_example.json
     */
    /**
     * Sample code: Get discovered security solutions from a security data location.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getDiscoveredSecuritySolutionsFromASecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.discoveredSecuritySolutions().listByHomeRegion("centralus", com.azure.core.util.Context.NONE);
    }
}

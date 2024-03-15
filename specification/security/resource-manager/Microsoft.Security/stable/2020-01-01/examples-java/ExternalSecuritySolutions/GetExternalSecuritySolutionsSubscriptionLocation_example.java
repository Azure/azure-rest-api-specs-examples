
/**
 * Samples for ExternalSecuritySolutions ListByHomeRegion.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ExternalSecuritySolutions/
     * GetExternalSecuritySolutionsSubscriptionLocation_example.json
     */
    /**
     * Sample code: Get external security solutions on a subscription from security data location.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getExternalSecuritySolutionsOnASubscriptionFromSecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.externalSecuritySolutions().listByHomeRegion("centralus", com.azure.core.util.Context.NONE);
    }
}

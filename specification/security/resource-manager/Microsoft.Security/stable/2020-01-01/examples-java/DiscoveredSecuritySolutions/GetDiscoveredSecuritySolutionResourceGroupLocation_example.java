
/**
 * Samples for DiscoveredSecuritySolutions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/DiscoveredSecuritySolutions
     * /GetDiscoveredSecuritySolutionResourceGroupLocation_example.json
     */
    /**
     * Sample code: Get discovered security solution from a security data location.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getDiscoveredSecuritySolutionFromASecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.discoveredSecuritySolutions().getWithResponse("myRg2", "centralus", "paloalto7",
            com.azure.core.util.Context.NONE);
    }
}

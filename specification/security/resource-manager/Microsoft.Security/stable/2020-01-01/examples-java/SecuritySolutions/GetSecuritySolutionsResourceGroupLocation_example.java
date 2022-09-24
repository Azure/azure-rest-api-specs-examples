import com.azure.core.util.Context;

/** Samples for SecuritySolutions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/SecuritySolutions/GetSecuritySolutionsResourceGroupLocation_example.json
     */
    /**
     * Sample code: Get a security solution from a security data location.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getASecuritySolutionFromASecurityDataLocation(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securitySolutions().getWithResponse("myRg2", "centralus", "paloalto7", Context.NONE);
    }
}

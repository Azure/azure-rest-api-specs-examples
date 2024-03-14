
/**
 * Samples for ExternalSecuritySolutions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ExternalSecuritySolutions/
     * GetExternalSecuritySolution_example.json
     */
    /**
     * Sample code: Get external security solution.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getExternalSecuritySolution(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.externalSecuritySolutions().getWithResponse("defaultresourcegroup-eus", "centralus",
            "aad_defaultworkspace-20ff7fc3-e762-44dd-bd96-b71116dcdc23-eus", com.azure.core.util.Context.NONE);
    }
}

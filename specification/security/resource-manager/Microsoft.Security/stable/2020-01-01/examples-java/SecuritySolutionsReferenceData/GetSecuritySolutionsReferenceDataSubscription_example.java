
/**
 * Samples for SecuritySolutionsReferenceData List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/
     * SecuritySolutionsReferenceData/GetSecuritySolutionsReferenceDataSubscription_example.json
     */
    /**
     * Sample code: Get security solutions.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecuritySolutions(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securitySolutionsReferenceDatas().listWithResponse(com.azure.core.util.Context.NONE);
    }
}

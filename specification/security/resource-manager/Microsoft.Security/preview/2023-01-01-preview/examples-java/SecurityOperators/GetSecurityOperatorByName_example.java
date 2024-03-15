
/**
 * Samples for SecurityOperators Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2023-01-01-preview/examples/SecurityOperators/
     * GetSecurityOperatorByName_example.json
     */
    /**
     * Sample code: Get a specific security operator by scope and securityOperatorName.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getASpecificSecurityOperatorByScopeAndSecurityOperatorName(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityOperators().getWithResponse("CloudPosture", "DefenderCSPMSecurityOperator",
            com.azure.core.util.Context.NONE);
    }
}

/** Samples for SecurityOperators List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-01-01-preview/examples/SecurityOperators/ListSecurityOperators_example.json
     */
    /**
     * Sample code: List SecurityOperators.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listSecurityOperators(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityOperators().listWithResponse("CloudPosture", com.azure.core.util.Context.NONE);
    }
}

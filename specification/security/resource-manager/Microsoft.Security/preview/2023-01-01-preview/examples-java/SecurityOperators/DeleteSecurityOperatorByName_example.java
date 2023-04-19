/** Samples for SecurityOperators Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2023-01-01-preview/examples/SecurityOperators/DeleteSecurityOperatorByName_example.json
     */
    /**
     * Sample code: Delete SecurityOperator on subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteSecurityOperatorOnSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .securityOperators()
            .deleteByResourceGroupWithResponse(
                "CloudPosture", "DefenderCSPMSecurityOperator", com.azure.core.util.Context.NONE);
    }
}

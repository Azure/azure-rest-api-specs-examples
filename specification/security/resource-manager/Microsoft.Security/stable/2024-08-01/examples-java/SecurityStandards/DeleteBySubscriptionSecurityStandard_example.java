
/**
 * Samples for SecurityStandards Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/SecurityStandards/
     * DeleteBySubscriptionSecurityStandard_example.json
     */
    /**
     * Sample code: Delete a security standard over subscription scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        deleteASecurityStandardOverSubscriptionScope(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityStandards().deleteByResourceGroupWithResponse(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
            com.azure.core.util.Context.NONE);
    }
}

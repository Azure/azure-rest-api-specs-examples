
/**
 * Samples for SecurityStandards List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/SecurityStandards/
     * ListBySubscriptionSecurityStandards_example.json
     */
    /**
     * Sample code: List security standards by subscription scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        listSecurityStandardsBySubscriptionScope(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityStandards().list("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for SecurityStandards Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2024-08-01/examples/SecurityStandards/
     * GetBySubscriptionSecurityStandard_example.json
     */
    /**
     * Sample code: Get a security standard over subscription scope.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        getASecurityStandardOverSubscriptionScope(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.securityStandards().getWithResponse("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23",
            "1f3afdf9-d0c9-4c3d-847f-89da613e70a8", com.azure.core.util.Context.NONE);
    }
}

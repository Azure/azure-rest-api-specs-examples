
/**
 * Samples for Alerts List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2022-01-01/examples/Alerts/
     * GetAlertsSubscription_example.json
     */
    /**
     * Sample code: Get security alerts on a subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityAlertsOnASubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.alerts().list(com.azure.core.util.Context.NONE);
    }
}

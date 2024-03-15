
/**
 * Samples for AlertsSuppressionRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/
     * AlertsSuppressionRules/GetAlertsSuppressionRules_example.json
     */
    /**
     * Sample code: Get suppression rules for subscription.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSuppressionRulesForSubscription(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.alertsSuppressionRules().list(null, com.azure.core.util.Context.NONE);
    }
}

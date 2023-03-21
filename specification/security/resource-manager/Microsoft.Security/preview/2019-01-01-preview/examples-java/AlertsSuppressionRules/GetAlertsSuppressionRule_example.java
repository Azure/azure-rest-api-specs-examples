/** Samples for AlertsSuppressionRules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/GetAlertsSuppressionRule_example.json
     */
    /**
     * Sample code: Get suppression alert rule for subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getSuppressionAlertRuleForSubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.alertsSuppressionRules().getWithResponse("dismissIpAnomalyAlerts", com.azure.core.util.Context.NONE);
    }
}

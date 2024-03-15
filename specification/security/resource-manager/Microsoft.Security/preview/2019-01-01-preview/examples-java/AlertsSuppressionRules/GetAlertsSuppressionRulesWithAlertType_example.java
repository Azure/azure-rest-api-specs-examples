
/**
 * Samples for AlertsSuppressionRules List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/
     * AlertsSuppressionRules/GetAlertsSuppressionRulesWithAlertType_example.json
     */
    /**
     * Sample code: Get suppression alert rule for subscription, filtered by AlertType.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getSuppressionAlertRuleForSubscriptionFilteredByAlertType(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.alertsSuppressionRules().list("IpAnomaly", com.azure.core.util.Context.NONE);
    }
}

/** Samples for AlertsSuppressionRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/DeleteAlertsSuppressionRule_example.json
     */
    /**
     * Sample code: Delete suppression rule data for a subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void deleteSuppressionRuleDataForASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.alertsSuppressionRules().deleteWithResponse("dismissIpAnomalyAlerts", com.azure.core.util.Context.NONE);
    }
}

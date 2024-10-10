
/**
 * Samples for AlertRuleRecommendations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2023-01-01-preview/examples/
     * AlertRuleRecommendations_GetBySubscription_MAC.json
     */
    /**
     * Sample code: List alert rule recommendations for Monitoring accounts at subscription level.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void listAlertRuleRecommendationsForMonitoringAccountsAtSubscriptionLevel(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertRuleRecommendations().list("microsoft.monitor/accounts", com.azure.core.util.Context.NONE);
    }
}

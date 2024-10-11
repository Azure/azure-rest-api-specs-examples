
/**
 * Samples for AlertRuleRecommendations ListByResource.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2023-01-01-preview/examples/
     * AlertRuleRecommendations_GetByResource_MAC.json
     */
    /**
     * Sample code: List alert rule recommendations for Monitoring accounts at resource level.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void listAlertRuleRecommendationsForMonitoringAccountsAtResourceLevel(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.alertRuleRecommendations().listByResource(
            "subscriptions/2f00cc51-6809-498f-9ffc-48c42aff570d/resourceGroups/GenevaAlertRP-RunnerResources-eastus/providers/microsoft.monitor/accounts/alertsrp-eastus-pgms",
            com.azure.core.util.Context.NONE);
    }
}

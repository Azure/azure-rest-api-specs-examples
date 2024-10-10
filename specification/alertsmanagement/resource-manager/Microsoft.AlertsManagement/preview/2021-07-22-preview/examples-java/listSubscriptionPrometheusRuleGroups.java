
/**
 * Samples for PrometheusRuleGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2021-07-22-preview/examples/
     * listSubscriptionPrometheusRuleGroups.json
     */
    /**
     * Sample code: ListSubscriptionResourcePrometheusRuleGroups.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void listSubscriptionResourcePrometheusRuleGroups(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.prometheusRuleGroups().list(com.azure.core.util.Context.NONE);
    }
}

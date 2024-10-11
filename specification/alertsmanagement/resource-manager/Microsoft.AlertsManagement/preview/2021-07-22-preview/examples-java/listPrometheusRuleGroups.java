
/**
 * Samples for PrometheusRuleGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2021-07-22-preview/examples/
     * listPrometheusRuleGroups.json
     */
    /**
     * Sample code: ListResourcePrometheusRuleGroups.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void
        listResourcePrometheusRuleGroups(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.prometheusRuleGroups().listByResourceGroup("promResourceGroup", com.azure.core.util.Context.NONE);
    }
}

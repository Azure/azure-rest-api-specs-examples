
/**
 * Samples for PrometheusRuleGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2021-07-22-preview/examples/
     * getPrometheusRuleGroup.json
     */
    /**
     * Sample code: GetPrometheusRuleGroup.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void
        getPrometheusRuleGroup(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.prometheusRuleGroups().getByResourceGroupWithResponse("promResourceGroup", "myPrometheusRuleGroup",
            com.azure.core.util.Context.NONE);
    }
}

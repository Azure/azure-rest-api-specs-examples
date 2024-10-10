
/**
 * Samples for PrometheusRuleGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2021-07-22-preview/examples/
     * deletePrometheusRuleGroup.json
     */
    /**
     * Sample code: DeletePrometheusRuleGroup.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void
        deletePrometheusRuleGroup(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.prometheusRuleGroups().deleteByResourceGroupWithResponse("promResourceGroup", "myPrometheusRuleGroup",
            com.azure.core.util.Context.NONE);
    }
}

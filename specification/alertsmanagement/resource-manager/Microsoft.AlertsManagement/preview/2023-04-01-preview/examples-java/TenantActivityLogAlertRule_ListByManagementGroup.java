
/**
 * Samples for TenantActivityLogAlerts ListByManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2023-04-01-preview/examples/
     * TenantActivityLogAlertRule_ListByManagementGroup.json
     */
    /**
     * Sample code: List Activity Log Alerts by management group.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void listActivityLogAlertsByManagementGroup(
        com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.tenantActivityLogAlerts().listByManagementGroup("72f988bf-86f1-41af-91ab-2d7cd011db47",
            com.azure.core.util.Context.NONE);
    }
}

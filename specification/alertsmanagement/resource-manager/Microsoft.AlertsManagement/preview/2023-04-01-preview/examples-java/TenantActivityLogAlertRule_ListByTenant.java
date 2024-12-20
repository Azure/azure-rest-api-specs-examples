
/**
 * Samples for TenantActivityLogAlerts ListByTenant.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2023-04-01-preview/examples/
     * TenantActivityLogAlertRule_ListByTenant.json
     */
    /**
     * Sample code: List Activity Log Alerts by tenant.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void
        listActivityLogAlertsByTenant(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.tenantActivityLogAlerts().listByTenant(com.azure.core.util.Context.NONE);
    }
}

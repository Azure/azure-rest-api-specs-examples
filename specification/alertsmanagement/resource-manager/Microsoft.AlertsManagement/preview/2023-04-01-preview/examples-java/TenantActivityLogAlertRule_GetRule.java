
/**
 * Samples for TenantActivityLogAlerts Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2023-04-01-preview/examples/
     * TenantActivityLogAlertRule_GetRule.json
     */
    /**
     * Sample code: Get a Tenant Activity Log Alert rule.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void
        getATenantActivityLogAlertRule(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.tenantActivityLogAlerts().getWithResponse("72f988bf-86f1-41af-91ab-2d7cd011db47",
            "SampleActivityLogAlertSHRuleOnTenantLevel", com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for TenantActivityLogAlerts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2023-04-01-preview/examples/
     * TenantActivityLogAlertRule_DeleteRule.json
     */
    /**
     * Sample code: Delete a Tenant Activity Log Alert rule.
     * 
     * @param manager Entry point to AlertsManagementManager.
     */
    public static void
        deleteATenantActivityLogAlertRule(com.azure.resourcemanager.alertsmanagement.AlertsManagementManager manager) {
        manager.tenantActivityLogAlerts().deleteByResourceGroupWithResponse("72f988bf-86f1-41af-91ab-2d7cd011db47",
            "SampleActivityLogAlertSHRuleOnTenantLevel", com.azure.core.util.Context.NONE);
    }
}

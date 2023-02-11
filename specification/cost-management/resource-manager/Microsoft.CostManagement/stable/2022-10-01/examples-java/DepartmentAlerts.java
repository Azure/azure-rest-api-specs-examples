/** Samples for Alerts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/DepartmentAlerts.json
     */
    /**
     * Sample code: DepartmentAlerts.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void departmentAlerts(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .alerts()
            .listWithResponse(
                "providers/Microsoft.Billing/billingAccounts/12345:6789/departments/123",
                com.azure.core.util.Context.NONE);
    }
}

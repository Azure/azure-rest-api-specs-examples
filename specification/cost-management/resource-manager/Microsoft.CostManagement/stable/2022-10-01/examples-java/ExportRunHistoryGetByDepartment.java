/** Samples for Exports GetExecutionHistory. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportRunHistoryGetByDepartment.json
     */
    /**
     * Sample code: ExportRunHistoryGetByDepartment.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunHistoryGetByDepartment(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getExecutionHistoryWithResponse(
                "providers/Microsoft.Billing/billingAccounts/12/departments/1234",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}

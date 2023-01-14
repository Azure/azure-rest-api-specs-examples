/** Samples for Exports GetExecutionHistory. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportRunHistoryGetByBillingAccount.json
     */
    /**
     * Sample code: ExportRunHistoryGetByBillingAccount.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunHistoryGetByBillingAccount(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getExecutionHistoryWithResponse(
                "providers/Microsoft.Billing/billingAccounts/123456", "TestExport", com.azure.core.util.Context.NONE);
    }
}

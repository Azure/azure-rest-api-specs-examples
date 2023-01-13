/** Samples for Exports Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportRunByBillingAccount.json
     */
    /**
     * Sample code: ExportRunByBillingAccount.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunByBillingAccount(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .executeWithResponse(
                "providers/Microsoft.Billing/billingAccounts/123456", "TestExport", com.azure.core.util.Context.NONE);
    }
}

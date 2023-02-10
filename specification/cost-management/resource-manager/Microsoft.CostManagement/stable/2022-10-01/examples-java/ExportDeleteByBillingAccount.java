/** Samples for Exports Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportDeleteByBillingAccount.json
     */
    /**
     * Sample code: ExportDeleteByBillingAccount.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportDeleteByBillingAccount(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .deleteByResourceGroupWithResponse(
                "providers/Microsoft.Billing/billingAccounts/123456", "TestExport", com.azure.core.util.Context.NONE);
    }
}

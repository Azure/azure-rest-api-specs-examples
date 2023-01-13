/** Samples for Exports Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportGetByBillingAccount.json
     */
    /**
     * Sample code: ExportGetByBillingAccount.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportGetByBillingAccount(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getWithResponse(
                "providers/Microsoft.Billing/billingAccounts/123456",
                "TestExport",
                null,
                com.azure.core.util.Context.NONE);
    }
}

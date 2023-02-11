/** Samples for Exports List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportsGetByBillingAccount.json
     */
    /**
     * Sample code: ExportsGetByBillingAccount.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportsGetByBillingAccount(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .listWithResponse(
                "providers/Microsoft.Billing/billingAccounts/123456", null, com.azure.core.util.Context.NONE);
    }
}

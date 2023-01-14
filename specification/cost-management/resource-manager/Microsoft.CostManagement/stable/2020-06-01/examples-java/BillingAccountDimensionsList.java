/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/BillingAccountDimensionsList.json
     */
    /**
     * Sample code: BillingAccountDimensionsList-Legacy.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void billingAccountDimensionsListLegacy(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Billing/billingAccounts/100",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}

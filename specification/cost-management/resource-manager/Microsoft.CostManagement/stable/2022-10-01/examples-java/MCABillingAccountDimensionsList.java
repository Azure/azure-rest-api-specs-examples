/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/MCABillingAccountDimensionsList.json
     */
    /**
     * Sample code: BillingAccountDimensionsList-MCA.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void billingAccountDimensionsListMCA(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Billing/billingAccounts/12345:6789",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}

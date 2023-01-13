/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/BillingAccountDimensionsListWithFilter.json
     */
    /**
     * Sample code: BillingAccountDimensionsListWithFilter-Legacy.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void billingAccountDimensionsListWithFilterLegacy(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Billing/billingAccounts/100",
                "properties/category eq 'resourceId'",
                "properties/data",
                null,
                5,
                com.azure.core.util.Context.NONE);
    }
}

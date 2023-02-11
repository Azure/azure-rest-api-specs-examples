/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/SubscriptionDimensionsList.json
     */
    /**
     * Sample code: SubscriptionDimensionsList-Legacy.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void subscriptionDimensionsListLegacy(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                null,
                "properties/data",
                null,
                5,
                com.azure.core.util.Context.NONE);
    }
}

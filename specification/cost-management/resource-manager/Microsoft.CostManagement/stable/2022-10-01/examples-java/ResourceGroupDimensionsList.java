/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ResourceGroupDimensionsList.json
     */
    /**
     * Sample code: ResourceGroupDimensionsList-Legacy.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void resourceGroupDimensionsListLegacy(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/system.orlando",
                null,
                "properties/data",
                null,
                5,
                com.azure.core.util.Context.NONE);
    }
}

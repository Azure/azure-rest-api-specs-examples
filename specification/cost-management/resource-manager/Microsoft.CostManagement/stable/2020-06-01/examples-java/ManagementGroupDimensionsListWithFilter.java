/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ManagementGroupDimensionsListWithFilter.json
     */
    /**
     * Sample code: ManagementGroupDimensionsListWithFilter-Legacy.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void managementGroupDimensionsListWithFilterLegacy(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Management/managementGroups/MyMgId",
                "properties/category eq 'resourceId'",
                "properties/data",
                null,
                5,
                com.azure.core.util.Context.NONE);
    }
}

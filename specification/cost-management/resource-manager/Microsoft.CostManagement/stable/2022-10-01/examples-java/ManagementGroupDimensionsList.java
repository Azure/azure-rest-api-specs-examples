/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ManagementGroupDimensionsList.json
     */
    /**
     * Sample code: ManagementGroupDimensionsList-Legacy.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void managementGroupDimensionsListLegacy(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Management/managementGroups/MyMgId",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}

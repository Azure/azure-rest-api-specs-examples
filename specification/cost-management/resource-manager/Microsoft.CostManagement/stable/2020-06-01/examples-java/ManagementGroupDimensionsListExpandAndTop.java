/** Samples for Dimensions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ManagementGroupDimensionsListExpandAndTop.json
     */
    /**
     * Sample code: ManagementGroupDimensionsListExpandAndTop-Legacy.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void managementGroupDimensionsListExpandAndTopLegacy(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .list(
                "providers/Microsoft.Management/managementGroups/MyMgId",
                null,
                "properties/data",
                null,
                5,
                com.azure.core.util.Context.NONE);
    }
}

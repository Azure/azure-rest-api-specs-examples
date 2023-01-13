/** Samples for Exports Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportDeleteByManagementGroup.json
     */
    /**
     * Sample code: ExportDeleteByManagementGroup.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportDeleteByManagementGroup(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .deleteByResourceGroupWithResponse(
                "providers/Microsoft.Management/managementGroups/TestMG",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}

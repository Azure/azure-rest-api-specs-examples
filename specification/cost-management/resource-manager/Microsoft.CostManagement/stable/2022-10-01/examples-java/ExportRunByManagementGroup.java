/** Samples for Exports Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportRunByManagementGroup.json
     */
    /**
     * Sample code: ExportRunByManagementGroup.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunByManagementGroup(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .executeWithResponse(
                "providers/Microsoft.Management/managementGroups/TestMG",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}

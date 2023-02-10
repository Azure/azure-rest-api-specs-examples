/** Samples for Exports GetExecutionHistory. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportRunHistoryGetByManagementGroup.json
     */
    /**
     * Sample code: ExportRunHistoryGetByManagementGroup.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunHistoryGetByManagementGroup(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getExecutionHistoryWithResponse(
                "providers/Microsoft.Management/managementGroups/TestMG",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}

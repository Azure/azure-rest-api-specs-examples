/** Samples for Exports Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportGetByManagementGroup.json
     */
    /**
     * Sample code: ExportGetByManagementGroup.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportGetByManagementGroup(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getWithResponse(
                "providers/Microsoft.Management/managementGroups/TestMG",
                "TestExport",
                null,
                com.azure.core.util.Context.NONE);
    }
}

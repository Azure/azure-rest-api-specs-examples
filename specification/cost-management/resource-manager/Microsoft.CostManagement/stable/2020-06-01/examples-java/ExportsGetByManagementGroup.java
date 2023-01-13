/** Samples for Exports List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportsGetByManagementGroup.json
     */
    /**
     * Sample code: ExportsGetByManagementGroup.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportsGetByManagementGroup(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .listWithResponse(
                "providers/Microsoft.Management/managementGroups/TestMG", null, com.azure.core.util.Context.NONE);
    }
}

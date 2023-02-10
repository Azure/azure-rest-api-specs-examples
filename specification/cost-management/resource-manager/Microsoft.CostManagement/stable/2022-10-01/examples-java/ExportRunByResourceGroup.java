/** Samples for Exports Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportRunByResourceGroup.json
     */
    /**
     * Sample code: ExportRunByResourceGroup.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunByResourceGroup(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .executeWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}

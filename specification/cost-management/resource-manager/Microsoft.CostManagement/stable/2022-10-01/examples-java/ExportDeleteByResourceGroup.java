/** Samples for Exports Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportDeleteByResourceGroup.json
     */
    /**
     * Sample code: ExportDeleteByResourceGroup.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportDeleteByResourceGroup(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .deleteByResourceGroupWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}

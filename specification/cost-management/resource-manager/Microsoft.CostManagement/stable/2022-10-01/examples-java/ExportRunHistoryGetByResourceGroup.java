/** Samples for Exports GetExecutionHistory. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExportRunHistoryGetByResourceGroup.json
     */
    /**
     * Sample code: ExportRunHistoryGetByResourceGroup.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunHistoryGetByResourceGroup(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getExecutionHistoryWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG",
                "TestExport",
                com.azure.core.util.Context.NONE);
    }
}

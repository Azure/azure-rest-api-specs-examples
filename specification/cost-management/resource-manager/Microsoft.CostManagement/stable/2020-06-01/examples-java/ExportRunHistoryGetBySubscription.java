/** Samples for Exports GetExecutionHistory. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportRunHistoryGetBySubscription.json
     */
    /**
     * Sample code: ExportRunHistoryGetBySubscription.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunHistoryGetBySubscription(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getExecutionHistoryWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000", "TestExport", com.azure.core.util.Context.NONE);
    }
}

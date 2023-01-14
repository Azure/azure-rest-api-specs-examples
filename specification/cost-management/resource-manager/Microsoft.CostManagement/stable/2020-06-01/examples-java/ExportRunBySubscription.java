/** Samples for Exports Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportRunBySubscription.json
     */
    /**
     * Sample code: ExportRunBySubscription.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportRunBySubscription(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .executeWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000", "TestExport", com.azure.core.util.Context.NONE);
    }
}

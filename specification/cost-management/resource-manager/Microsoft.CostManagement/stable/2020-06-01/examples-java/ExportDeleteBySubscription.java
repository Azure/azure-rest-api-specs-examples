/** Samples for Exports Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportDeleteBySubscription.json
     */
    /**
     * Sample code: ExportDeleteBySubscription.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportDeleteBySubscription(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .deleteByResourceGroupWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000", "TestExport", com.azure.core.util.Context.NONE);
    }
}

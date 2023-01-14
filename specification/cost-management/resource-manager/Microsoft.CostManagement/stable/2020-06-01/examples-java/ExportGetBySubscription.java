/** Samples for Exports Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportGetBySubscription.json
     */
    /**
     * Sample code: ExportGetBySubscription.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportGetBySubscription(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .getWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                "TestExport",
                null,
                com.azure.core.util.Context.NONE);
    }
}

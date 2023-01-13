/** Samples for Exports List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ExportsGetBySubscription.json
     */
    /**
     * Sample code: ExportsGetBySubscription.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void exportsGetBySubscription(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .exports()
            .listWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000", null, com.azure.core.util.Context.NONE);
    }
}

/** Samples for Alerts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/ResourceGroupAlerts.json
     */
    /**
     * Sample code: ResourceGroupAlerts.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void resourceGroupAlerts(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .alerts()
            .listWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ScreenSharingTest-peer",
                com.azure.core.util.Context.NONE);
    }
}

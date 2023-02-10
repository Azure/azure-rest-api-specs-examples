/** Samples for Alerts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/SingleResourceGroupAlert.json
     */
    /**
     * Sample code: SingleResourceGroupAlerts.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void singleResourceGroupAlerts(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .alerts()
            .getWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/ScreenSharingTest-peer",
                "22222222-2222-2222-2222-222222222222",
                com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for ManagedDashboards GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/Dashboard_Get.json
     */
    /**
     * Sample code: Dashboard_Get.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void dashboardGet(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.managedDashboards().getByResourceGroupWithResponse("myResourceGroup", "myDashboard",
            com.azure.core.util.Context.NONE);
    }
}

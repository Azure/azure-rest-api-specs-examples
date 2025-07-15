
/**
 * Samples for ManagedDashboards Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Dashboard_Delete.json
     */
    /**
     * Sample code: Dashboard_Delete.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void dashboardDelete(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.managedDashboards().deleteByResourceGroupWithResponse("myResourceGroup", "myDashboard",
            com.azure.core.util.Context.NONE);
    }
}

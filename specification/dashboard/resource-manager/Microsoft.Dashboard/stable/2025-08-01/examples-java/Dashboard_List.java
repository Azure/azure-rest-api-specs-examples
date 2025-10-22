
/**
 * Samples for ManagedDashboards List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/Dashboard_List.json
     */
    /**
     * Sample code: Dashboard_ListByResourceGroup.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void dashboardListByResourceGroup(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.managedDashboards().list(com.azure.core.util.Context.NONE);
    }
}

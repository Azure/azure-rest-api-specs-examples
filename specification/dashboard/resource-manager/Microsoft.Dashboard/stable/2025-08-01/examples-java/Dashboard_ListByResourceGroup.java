
/**
 * Samples for ManagedDashboards ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/Dashboard_ListByResourceGroup.json
     */
    /**
     * Sample code: Dashboard_ListByResourceGroup.
     * 
     * @param manager Entry point to DashboardManager.
     */
    public static void dashboardListByResourceGroup(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.managedDashboards().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}

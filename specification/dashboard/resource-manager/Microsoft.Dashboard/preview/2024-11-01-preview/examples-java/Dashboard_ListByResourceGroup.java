
/**
 * Samples for ManagedDashboards ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/Dashboard_ListByResourceGroup.json
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

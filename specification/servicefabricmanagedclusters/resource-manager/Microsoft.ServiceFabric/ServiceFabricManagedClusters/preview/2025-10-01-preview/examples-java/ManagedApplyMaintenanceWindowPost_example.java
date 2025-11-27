
/**
 * Samples for ManagedApplyMaintenanceWindow Post.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ManagedApplyMaintenanceWindowPost_example.json
     */
    /**
     * Sample code: Apply Maintenance Window Status.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void applyMaintenanceWindowStatus(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedApplyMaintenanceWindows().postWithResponse("resourceGroup1", "mycluster1",
            com.azure.core.util.Context.NONE);
    }
}

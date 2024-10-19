
/**
 * Samples for ManagedApplyMaintenanceWindow Post.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicefabricmanagedclusters/resource-manager/Microsoft.ServiceFabric/preview/2024-06-01-preview/
     * examples/ManagedApplyMaintenanceWindowPost_example.json
     */
    /**
     * Sample code: Maintenance Window Status.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void maintenanceWindowStatus(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedApplyMaintenanceWindows().postWithResponse("resourceGroup1", "mycluster1",
            com.azure.core.util.Context.NONE);
    }
}

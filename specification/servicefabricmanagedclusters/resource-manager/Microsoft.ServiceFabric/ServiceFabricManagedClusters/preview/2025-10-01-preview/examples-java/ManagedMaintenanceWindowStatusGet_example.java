
/**
 * Samples for ManagedMaintenanceWindowStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/ManagedMaintenanceWindowStatusGet_example.json
     */
    /**
     * Sample code: Get Maintenance Window Status.
     * 
     * @param manager Entry point to ServiceFabricManagedClustersManager.
     */
    public static void getMaintenanceWindowStatus(
        com.azure.resourcemanager.servicefabricmanagedclusters.ServiceFabricManagedClustersManager manager) {
        manager.managedMaintenanceWindowStatuses().getWithResponse("resourceGroup1", "mycluster1",
            com.azure.core.util.Context.NONE);
    }
}

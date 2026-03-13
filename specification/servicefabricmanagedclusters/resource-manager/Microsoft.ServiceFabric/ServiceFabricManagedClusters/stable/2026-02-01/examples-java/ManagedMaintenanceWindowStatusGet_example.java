
/**
 * Samples for ManagedMaintenanceWindowStatuses Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ManagedMaintenanceWindowStatusGet_example.json
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

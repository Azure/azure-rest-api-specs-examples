
/**
 * Samples for MaintenanceConfigurations ListByManagedCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/MaintenanceConfigurationsList_MaintenanceWindow.json
     */
    /**
     * Sample code: List maintenance configurations configured with maintenance window by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listMaintenanceConfigurationsConfiguredWithMaintenanceWindowByManagedCluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getMaintenanceConfigurations().listByManagedCluster("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}

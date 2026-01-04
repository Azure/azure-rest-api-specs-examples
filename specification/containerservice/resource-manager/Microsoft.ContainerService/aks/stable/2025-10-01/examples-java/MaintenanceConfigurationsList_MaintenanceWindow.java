
/**
 * Samples for MaintenanceConfigurations ListByManagedCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * MaintenanceConfigurationsList_MaintenanceWindow.json
     */
    /**
     * Sample code: List maintenance configurations configured with maintenance window by Managed Cluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listMaintenanceConfigurationsConfiguredWithMaintenanceWindowByManagedCluster(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getMaintenanceConfigurations().listByManagedCluster("rg1",
            "clustername1", com.azure.core.util.Context.NONE);
    }
}

import com.azure.core.util.Context;

/** Samples for MaintenanceConfigurations ListByManagedCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-10-01/examples/MaintenanceConfigurationsList.json
     */
    /**
     * Sample code: List maintenance configurations by Managed Cluster.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listMaintenanceConfigurationsByManagedCluster(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getMaintenanceConfigurations()
            .listByManagedCluster("rg1", "clustername1", Context.NONE);
    }
}

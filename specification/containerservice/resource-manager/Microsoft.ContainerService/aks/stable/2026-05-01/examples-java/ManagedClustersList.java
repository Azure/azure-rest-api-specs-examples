
/**
 * Samples for ManagedClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/ManagedClustersList.json
     */
    /**
     * Sample code: List Managed Clusters.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listManagedClusters(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().list(com.azure.core.util.Context.NONE);
    }
}

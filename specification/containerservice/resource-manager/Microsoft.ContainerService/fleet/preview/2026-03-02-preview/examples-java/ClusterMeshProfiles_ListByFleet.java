
/**
 * Samples for ClusterMeshProfiles ListByFleet.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/ClusterMeshProfiles_ListByFleet.json
     */
    /**
     * Sample code: List the ClusterMeshProfile resources by fleet.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void listTheClusterMeshProfileResourcesByFleet(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.clusterMeshProfiles().listByFleet("rgfleets", "fleet1", com.azure.core.util.Context.NONE);
    }
}

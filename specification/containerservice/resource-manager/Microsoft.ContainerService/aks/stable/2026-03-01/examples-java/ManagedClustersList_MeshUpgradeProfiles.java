
/**
 * Samples for ManagedClusters ListMeshUpgradeProfiles.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ManagedClustersList_MeshUpgradeProfiles.json
     */
    /**
     * Sample code: Lists version compatibility and upgrade profile for all service meshes in a cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listsVersionCompatibilityAndUpgradeProfileForAllServiceMeshesInACluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().listMeshUpgradeProfiles("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}

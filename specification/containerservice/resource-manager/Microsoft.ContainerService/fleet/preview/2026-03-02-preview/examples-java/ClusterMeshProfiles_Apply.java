
/**
 * Samples for ClusterMeshProfiles Apply.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/ClusterMeshProfiles_Apply.json
     */
    /**
     * Sample code: Apply a ClusterMeshProfile.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void
        applyAClusterMeshProfile(com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.clusterMeshProfiles().apply("rgfleets", "fleet1", "clustermeshprofile1", null,
            com.azure.core.util.Context.NONE);
    }
}

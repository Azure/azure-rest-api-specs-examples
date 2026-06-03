
/**
 * Samples for ClusterMeshProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/ClusterMeshProfiles_Delete.json
     */
    /**
     * Sample code: Delete an ClusterMeshProfile resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void deleteAnClusterMeshProfileResource(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.clusterMeshProfiles().delete("rg1", "fleet1", "clustermeshprofile1", null,
            com.azure.core.util.Context.NONE);
    }
}

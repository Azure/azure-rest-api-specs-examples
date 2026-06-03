
/**
 * Samples for ClusterMeshProfiles Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02-preview/ClusterMeshProfiles_Get.json
     */
    /**
     * Sample code: Get a ClusterMeshProfile resource.
     * 
     * @param manager Entry point to ContainerServiceFleetManager.
     */
    public static void getAClusterMeshProfileResource(
        com.azure.resourcemanager.containerservicefleet.ContainerServiceFleetManager manager) {
        manager.clusterMeshProfiles().getWithResponse("rgfleets", "fleet1", "clustermeshprofile1",
            com.azure.core.util.Context.NONE);
    }
}


/**
 * Samples for ManagedClusters GetMeshRevisionProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ManagedClustersGet_MeshRevisionProfile.json
     */
    /**
     * Sample code: Get a mesh revision profile for a mesh mode.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void getAMeshRevisionProfileForAMeshMode(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().getMeshRevisionProfileWithResponse("location1", "istio",
            com.azure.core.util.Context.NONE);
    }
}

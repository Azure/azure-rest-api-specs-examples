
/**
 * Samples for ManagedClusters ListMeshRevisionProfiles.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/ManagedClustersList_MeshRevisionProfiles.json
     */
    /**
     * Sample code: List mesh revision profiles in a location.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listMeshRevisionProfilesInALocation(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().listMeshRevisionProfiles("location1",
            com.azure.core.util.Context.NONE);
    }
}

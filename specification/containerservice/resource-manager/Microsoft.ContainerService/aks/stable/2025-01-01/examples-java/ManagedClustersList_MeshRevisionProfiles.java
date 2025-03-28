
/**
 * Samples for ManagedClusters ListMeshRevisionProfiles.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-01-01/examples/
     * ManagedClustersList_MeshRevisionProfiles.json
     */
    /**
     * Sample code: List mesh revision profiles in a location.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listMeshRevisionProfilesInALocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters().listMeshRevisionProfiles("location1",
            com.azure.core.util.Context.NONE);
    }
}

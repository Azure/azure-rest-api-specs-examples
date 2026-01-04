
/**
 * Samples for ManagedClusters GetMeshRevisionProfile.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedClustersGet_MeshRevisionProfile.json
     */
    /**
     * Sample code: Get a mesh revision profile for a mesh mode.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAMeshRevisionProfileForAMeshMode(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters()
            .getMeshRevisionProfileWithResponse("location1", "istio", com.azure.core.util.Context.NONE);
    }
}

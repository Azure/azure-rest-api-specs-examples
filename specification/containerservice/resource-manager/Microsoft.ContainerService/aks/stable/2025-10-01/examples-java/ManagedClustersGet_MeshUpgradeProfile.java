
/**
 * Samples for ManagedClusters GetMeshUpgradeProfile.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedClustersGet_MeshUpgradeProfile.json
     */
    /**
     * Sample code: Gets version compatibility and upgrade profile for a service mesh in a cluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsVersionCompatibilityAndUpgradeProfileForAServiceMeshInACluster(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters()
            .getMeshUpgradeProfileWithResponse("rg1", "clustername1", "istio", com.azure.core.util.Context.NONE);
    }
}

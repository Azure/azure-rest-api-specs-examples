
/**
 * Samples for ManagedClusters ListMeshUpgradeProfiles.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedClustersList_MeshUpgradeProfiles.json
     */
    /**
     * Sample code: Lists version compatibility and upgrade profile for all service meshes in a cluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsVersionCompatibilityAndUpgradeProfileForAllServiceMeshesInACluster(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters().listMeshUpgradeProfiles("rg1",
            "clustername1", com.azure.core.util.Context.NONE);
    }
}
